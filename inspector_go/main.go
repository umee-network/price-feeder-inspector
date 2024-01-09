package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/umee-network/price_inspector/config"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "price_inspector",
		Short: "Web Go is a web server for fetching transaction data",
		Long:  `A Fast and Flexible web server built with love by in Go.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			network, err := cmd.Flags().GetString(config.Network)
			if err != nil {
				return err
			}

			configFile, err := cmd.Flags().GetString(config.Config)
			if err != nil {
				return err
			}

			cfg, err := config.ReadConfig(configFile)
			if err != nil {
				return err
			}
			var rpc, grpc string
			switch network {
			case "mainnet":
				rpc = cfg.Networks.Mainnet.RPC
				grpc = cfg.Networks.Mainnet.GRPC
			case "canon":
				rpc = cfg.Networks.Canon.RPC
				grpc = cfg.Networks.Canon.GRPC
			case "local":
				rpc = cfg.Networks.Local.RPC
				grpc = cfg.Networks.Local.GRPC
			default:
				log.Fatalf("this netwokr %s is not supported", network)
			}
			return StartInspector(cfg, rpc, grpc)
		},
	}

	rootCmd.Flags().String(config.Network, "canon", "Network to use (canon/mainnet/local)")
	rootCmd.Flags().String(config.Config, "config.json", "Path to the configuration file")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
