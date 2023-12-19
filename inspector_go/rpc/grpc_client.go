package rpc

import (
	"log"
	"os"
	"time"

	"github.com/umee-network/umee/v6/sdkclient/query"
	otypes "github.com/umee-network/umee/v6/x/oracle/types"
)

// GetOracleGrpcClient returns a gRPC client for querying the Oracle module.
// It takes the gRPC endpoint as input and returns the Oracle query client.
func GetOracleGrpcClient(grpcEndpoint string) otypes.QueryClient {
	logger := log.New(os.Stderr, "chain-client", log.LstdFlags)
	queryClient, err := query.NewClient(logger, grpcEndpoint, 15*time.Second)
	if err != nil {
		log.Fatalf("Failed to initalize the query client %s and Error:  %s", grpcEndpoint, err.Error())
	}

	return otypes.NewQueryClient(queryClient.GrpcConn)
}
