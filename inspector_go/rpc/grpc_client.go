package rpc

import (
	"log"
	"os"
	"time"

	stypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/umee-network/umee/v6/sdkclient/query"
	otypes "github.com/umee-network/umee/v6/x/oracle/types"
)

func queryClient(grpcEndpoint string) *query.Client {
	logger := log.New(os.Stderr, "chain-client", log.LstdFlags)
	queryClient, err := query.NewClient(logger, grpcEndpoint, 15*time.Second)
	if err != nil {
		log.Fatalf("Failed to initalize the query client %s and Error:  %s", grpcEndpoint, err.Error())
	}

	return queryClient
}

// GetOracleGrpcClient returns a gRPC client for querying the Oracle module.
// It takes the gRPC endpoint as input and returns the Oracle query client.
func GetOracleGrpcClient(grpcEndpoint string) otypes.QueryClient {
	client := queryClient(grpcEndpoint)
	return otypes.NewQueryClient(client.GrpcConn)
}

// GetStakingGrpcClient returns a gRPC client for querying the staking module.
// It takes the gRPC endpoint as input and returns the staking query client.
func GetStakingGrpcClient(grpcEndpoint string) stypes.QueryClient {
	client := queryClient(grpcEndpoint)
	return stypes.NewQueryClient(client.GrpcConn)
}
