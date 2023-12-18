package rpc

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/umee-network/price_inspector/config"
	otypes "github.com/umee-network/umee/v6/x/oracle/types"
)

func GetAcceptedDenoms(codec testutil.TestEncodingConfig, api string) []string {
	// This function fetches the accepted denominations from the API.

	// Parameters:
	// api (string): The API from which to fetch the accepted denominations.

	// Returns:
	// []string: The accepted denominations fetched from the API.
	r, err := http.Get(api + "/umee/oracle/v1/params")
	if err != nil {
		log.Fatal("Failed to fetch data from " + api + "/umee/oracle/v1/params. Error: " + err.Error())
	}
	defer r.Body.Close()

	if r.StatusCode != 200 {
		log.Fatal("Failed to fetch data from " + api + "/umee/oracle/v1/params. Status code: " + strconv.Itoa(r.StatusCode))
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed to read response body. Error: %s", err)
		return nil
	}

	var resp otypes.QueryParamsResponse
	if err := codec.Codec.UnmarshalJSON(body, &resp); err != nil {
		log.Fatalf("Error while unmarshal oracle params resp. Error: %s", err)
	}
	symbolDenoms := make(map[string]bool)
	for _, a := range resp.Params.AcceptList {
		symbolDenoms[strings.ToUpper(a.SymbolDenom)] = true
	}

	var accepedDenoms []string
	for a := range symbolDenoms {
		accepedDenoms = append(accepedDenoms, a)
	}
	return accepedDenoms
}

func FetchTxSearchData(codec testutil.TestEncodingConfig, rpc string, txHeight int64) ([]config.Tx, error) {
	// url := fmt.Sprintf("%s/tx_search?query=\"tx.height=47931 AND message.action='/umee.oracle.v1.MsgAggregateExchangeRateVote'\"", rpc)
	url := fmt.Sprintf("%s/tx_search?query=%%22tx.height%%3D%d%%20AND%%20message.action%%3D'%%2Fumee.oracle.v1.MsgAggregateExchangeRateVote'%%22", rpc, txHeight)
	queryParams := "&per_page=100&page=1&order_by=\"desc\""
	fetchData := func(queryParams string) ([]config.Tx, string, error) {
		log.Println("Req URL : ", url+queryParams)
		resp, err := http.Get(url + queryParams)
		if err != nil {
			log.Printf("Failed to fetch data from %s. Error: %s", url, err)
			return nil, "", err
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			log.Printf("Failed to fetch data from %s. Status code: %d", url, resp.StatusCode)
			return nil, "", fmt.Errorf("HTTP request failed with status code %d", resp.StatusCode)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Failed to read response body. Error: %s", err)
			return nil, "", err
		}
		var txSearchResponse config.TxSearchResonse
		err = json.Unmarshal(body, &txSearchResponse)
		if err != nil {
			log.Printf("Failed to unmarshal response body. Error: %s", err)
			return nil, "", err
		}
		return txSearchResponse.Result.Txs, txSearchResponse.Result.TotalCount, nil
	}

	var resultTxs []config.Tx
	log.Printf("Getting txns on height %d and page %d", txHeight, 1)
	resultTxs, txsCount, err := fetchData(queryParams)
	if err != nil {
		log.Printf("Failed to get the txns. Error: %s", err)
		return nil, err
	}

	txsCountInt, err := strconv.ParseInt(txsCount, 10, 64)
	if err != nil {
		log.Printf("Failed to convert txsCount to int64. Error: %s", err)
		return nil, err
	}
	log.Printf("Total txns %d at Height %d\n", txsCountInt, txHeight)
	if txsCountInt > 100 {
		incre := txsCountInt / 100
		if txsCountInt%100 != 0 {
			incre = incre + 1
		}
		for page := 2; page <= int(incre); page++ {
			log.Printf("Getting txns on height %d and page %d", txHeight, page)
			r, _, err := fetchData(fmt.Sprintf("&per_page=100&page=%d&order_by=\"desc\"", page))
			if err != nil {
				log.Printf("Failed to get the txns. Error: %s", err)
				return nil, err
			}
			resultTxs = append(resultTxs, r...)
		}
	}

	return resultTxs, nil
}

func GetLatestHeight(rpc string) (int64, error) {
	url := rpc + "/status"
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to fetch data from %s. Error: %s", url, err)
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Printf("Failed to fetch data from %s. Status code: %d", url, resp.StatusCode)
		return 0, fmt.Errorf("HTTP request failed with status code %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response body. Error: %s", err)
		return 0, err
	}
	var statusResp config.StatusResponse
	err = json.Unmarshal(body, &statusResp)
	if err != nil {
		log.Printf("Failed to unmarshal response body. Error: %s", err)
		return 0, err
	}
	return strconv.ParseInt(statusResp.Result.SyncInfo.LatestBlockHeight, 10, 64)
}
