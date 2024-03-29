package main

import (
	"log"
	"strings"
	"time"

	"github.com/umee-network/price_inspector/config"
	"github.com/umee-network/price_inspector/rpc"
	"github.com/umee-network/price_inspector/utils"
	"github.com/umee-network/umee/v6/app"
)

func StartInspector(cfg config.Configuration, rpcUri, grpcUri string) error {
	codec := app.MakeEncodingConfig()
	log.Println("RPC URL : ", rpcUri)
	log.Println("GRPC URL : ", grpcUri)
	log.Println("WINDOW  : ", cfg.Window)

	rpcClient := rpc.GetRPCClient(rpcUri)

	aceeptedDenoms := rpc.GetAcceptedDenoms(codec, grpcUri)
	log.Println("Total Accepted Denoms ", len(aceeptedDenoms))
	log.Println("Oracle Accepted Denoms ", strings.Join(aceeptedDenoms, ","))

	latestBlockHeight := rpcClient.GetLatestHeight()
	lastBlockHeight := latestBlockHeight - cfg.Window
	log.Printf("Latest Block Height %d", latestBlockHeight)
	log.Printf("Last Block Height for fetching the data %d\n", lastBlockHeight)
	votePeriod := int64(5)
	expectedNoOfVotes := cfg.Window / votePeriod
	log.Printf("Expected No Of Votes %d in the Period", expectedNoOfVotes)

	votes := make(map[string]config.ValidatorsVotes)
	vInfo := make(map[string]string)

	if len(cfg.Validators) == 0 {
		log.Println("ℹ️ Getting all staking validators...")
		validators := rpc.GetValidators(codec, grpcUri)
		for _, val := range validators {
			votes[val.OperatorAddress] = config.NewValidatorVotes(val.OperatorAddress, expectedNoOfVotes)
			vInfo[val.OperatorAddress] = val.GetMoniker()
		}
	} else {
		for _, v := range cfg.Validators {
			votes[v.Validator] = config.NewValidatorVotes(v.Validator, expectedNoOfVotes)
			vInfo[v.Validator] = v.Moniker
		}
	}

	for i := lastBlockHeight; i < latestBlockHeight; i++ {
		log.Printf("Getting transactions on at height : %d\n", i)
		txs, err := rpc.FetchTxSearchData(rpcClient, i)
		if err != nil {
			return err
		}
		for _, r := range txs {
			vr, err := rpc.TxDecoder(r.Tx, codec)
			if err != nil {
				log.Printf("Failed to decode tx data. Error: %s", err)
			}
			height := r.Height
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
		printValidatorVotes(votes[v])
	}

	return nil
}

func printValidatorVotes(validatorVotes config.ValidatorsVotes) {
	log.Println("ExpectedNoOfVotes ", validatorVotes.ExpectedNoOfVotes)
	log.Println("Total Submited votes ", validatorVotes.NoOfVotes)
	log.Println("Votes WithOut MissCounter ", validatorVotes.VotesWithOutMissCounter)
	log.Println("Votes with missing denoms", validatorVotes.NoOfMissCounter)
	log.Println("Missing denoms at Height =>")
	for _, d := range validatorVotes.MissingDenomsAtHeight {
		log.Printf("At height %d missing denoms %s\n", d.Height, strings.Join(d.Denoms, ","))
	}
	log.Println("Missing Denoms Count")
	for k, c := range validatorVotes.MissingDenomsCount {
		log.Printf("Denom %s => %d", k, c)
	}
}
