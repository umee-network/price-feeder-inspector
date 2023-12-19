package rpc

import (
	"context"
	"log"
	"time"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	ctypes "github.com/cometbft/cometbft/rpc/core/types"
)

type RPCClient struct {
	*rpchttp.HTTP
}

func GetRPCClient(rpcUri string) RPCClient {
	rpcClient, err := rpchttp.New(rpcUri, "/websocket")
	if err != nil {
		log.Fatalf("Error while connecting to %s . Error : %s", rpcUri, err.Error())
	}

	return RPCClient{rpcClient}
}

func (c RPCClient) GetLatestHeight() int64 {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	status, err := c.Status(ctx)
	if err != nil {
		log.Fatalf("Error while get the latest height Error : %s", err.Error())
	}
	return status.SyncInfo.LatestBlockHeight
}

func (c RPCClient) SearchTx(query string, page int) ctypes.ResultTxSearch {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	perPageMax := 100
	result, err := c.TxSearch(ctx, query, false, &page, &perPageMax, "desc")
	if err != nil {
		log.Fatalf("Error while searching the txns Error : %s", err.Error())
	}
	return *result
}
