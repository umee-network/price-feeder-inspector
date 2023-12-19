package rpc

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	ctypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/cosmos/cosmos-sdk/types/module/testutil"
	otypes "github.com/umee-network/umee/v6/x/oracle/types"
)

// GetAcceptedDenoms returns a list of accepted denominations by oracle.
func GetAcceptedDenoms(codec testutil.TestEncodingConfig, grpcEndpoint string) []string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	oQueryClient := GetOracleGrpcClient(grpcEndpoint)
	resp, err := oQueryClient.Params(ctx, &otypes.QueryParams{})
	if err != nil {
		log.Fatalf("Failed to get the oracle params,Error: %s", err.Error())
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

// FetchTxSearchData fetches transaction search data based on the provided RPC client and transaction height.
// It constructs a query based on the transaction height and message action, then fetches the data using the RPC client.
// The function returns an array of ResultTx and an error, if any.
func FetchTxSearchData(rpcClient RPCClient, txHeight int64) ([]*ctypes.ResultTx, error) {
	query := fmt.Sprintf("tx.height=%d AND message.action='/umee.oracle.v1.MsgAggregateExchangeRateVote'", txHeight)
	fetchData := func(page int) ([]*ctypes.ResultTx, int) {
		searchResults := rpcClient.SearchTx(query, page)
		return searchResults.Txs, searchResults.TotalCount
	}

	var resultTxs []*ctypes.ResultTx
	log.Printf("Getting txns on height %d and page %d", txHeight, 1)
	resultTxs, txsCount := fetchData(int(1))

	log.Printf("Total txns %d at Height %d\n", txsCount, txHeight)
	if txsCount > 100 {
		incre := txsCount / 100
		if txsCount%100 != 0 {
			incre = incre + 1
		}
		for page := 2; page <= int(incre); page++ {
			log.Printf("Getting txns on height %d and page %d", txHeight, page)
			r, _ := fetchData(page)
			resultTxs = append(resultTxs, r...)
		}
	}

	return resultTxs, nil
}
