package main

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/umee-network/price_inspector/config"
	"github.com/umee-network/price_inspector/rpc"
	"github.com/umee-network/price_inspector/utils"
	"github.com/umee-network/umee/v6/app"
)

func StartInspector(cfg config.Configuration, rpcUri, apiUri string) error {
	codec := app.MakeEncodingConfig()
	log.Println("RPC URL : ", rpcUri)
	log.Println("API URL : ", apiUri)
	log.Println("WINDOW  : ", cfg.Window)

	aceeptedDenoms := rpc.GetAcceptedDenoms(codec, apiUri)
	log.Println("Total Accepted Denoms ", len(aceeptedDenoms))
	log.Println("Oracle Accepted Denoms ", aceeptedDenoms)

	latestBlockHeight, err := rpc.GetLatestHeight(rpcUri)
	if err != nil {
		return err
	}
	lastBlockHeight := latestBlockHeight - cfg.Window
	log.Printf("Latest Block Height %d", latestBlockHeight)
	log.Printf("Last Block Height for fetching the data %d\n", lastBlockHeight)
	votePeriod := int64(5)
	expectedNoOfVotes := cfg.Window / votePeriod
	log.Printf("Expected No Of Votes %d in the Period", expectedNoOfVotes)

	votes := make(map[string]config.ValidatorsVotes)
	vInfo := make(map[string]string)

	for _, v := range cfg.Validators {
		votes[v.Validator] = config.NewValidatorVotes(v.Validator, expectedNoOfVotes)
		vInfo[v.Validator] = v.Moniker
	}

	for i := lastBlockHeight; i < latestBlockHeight; i++ {
		log.Printf("Getting transactions on at height : %d\n", i)
		txs, err := rpc.FetchTxSearchData(codec, rpcUri, i)
		if err != nil {
			return err
		}
		var height int64
		for _, r := range txs {
			vr, err := rpc.TxDecoder(r.Tx, codec)
			if err != nil {
				log.Printf("Failed to decode tx data. Error: %s", err)
			}
			height, err = strconv.ParseInt(r.Height, 10, 64)
			if err != nil {
				return err
			}
			for _, c := range vr {
				var denoms []string
				for _, d := range c.ExgRatesTuples {
					denoms = append(denoms, strings.ToUpper(d.Denom))
				}
				val, ok := votes[c.Validator]
				if !ok {
					continue
				}
				val.NoOfVotes = val.NoOfVotes + 1
				missingDenoms := utils.GetMissingDenoms(aceeptedDenoms, denoms)
				if len(missingDenoms) > 0 {
					val.NoOfMissCounter = val.NoOfMissCounter + 1
					val.MissingDenomsAtHeight = append(val.MissingDenomsAtHeight, config.MissingDenomsWithHeight{
						Height: height,
						Denoms: missingDenoms,
					})
					for _, md := range missingDenoms {
						val.MissingDenomsCount[md] = val.MissingDenomsCount[md] + 1
					}
				} else {
					val.VotesWithOutMissCounter = val.VotesWithOutMissCounter + 1
				}

				votes[c.Validator] = val
			}
		}
		log.Println("Sleeping for 1 seconds...")
		time.Sleep(time.Second * 1)
	}

	for v := range votes {
		log.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
		log.Printf("ℹ️ Validator %s and Moniker %s\n", v, vInfo[v])
		log.Println("ExpectedNoOfVotes ", votes[v].ExpectedNoOfVotes)
		log.Println("Total Submited votes ", votes[v].NoOfVotes)
		log.Println("Votes WithOut MissCounter ", votes[v].VotesWithOutMissCounter)
		log.Println("Votes with missing denoms", votes[v].NoOfMissCounter)
		log.Println("Missing denoms at Height =>")
		for _, d := range votes[v].MissingDenomsAtHeight {
			log.Printf("At height %d missing denoms %s\n", d.Height, strings.Join(d.Denoms, ","))
		}
		log.Println("Missing Denoms Count")
		for k, c := range votes[v].MissingDenomsCount {
			log.Printf("Denom %s => %d", k, c)
		}
	}

	return nil
}
