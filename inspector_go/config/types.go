package config

import (
	"time"

	otypes "github.com/umee-network/umee/v6/x/oracle/types"
)

type ValidatorExgRates struct {
	Validator      string                    `json:"string"`
	ExgRatesTuples otypes.ExchangeRateTuples `json:"exgRatesTuples"`
}

type ValidatorsVotes struct {
	Validator         string `json:"validator"`
	ExpectedNoOfVotes int64  `json:"expectedNoOfVotes"`
	// submited total votes
	NoOfVotes int64 `json:"noOfVotes"`
	// votes with out missing denoms
	VotesWithOutMissCounter int64 `json:"votesWithOutMissCounter"`
	// total number of miss counter of validator
	NoOfMissCounter       int64                     `json:"noOfMissCounter"`
	MissingDenomsAtHeight []MissingDenomsWithHeight `json:"missingdenoms"`
	MissingDenomsCount    map[string]int64          `json:"missingDenomsCount"`
}

func NewValidatorVotes(validator string, expectedNoOfVotes int64) ValidatorsVotes {
	return ValidatorsVotes{
		Validator:               validator,
		ExpectedNoOfVotes:       expectedNoOfVotes,
		NoOfVotes:               0,
		VotesWithOutMissCounter: 0,
		NoOfMissCounter:         0,
		MissingDenomsAtHeight:   []MissingDenomsWithHeight{},
		MissingDenomsCount:      make(map[string]int64),
	}
}

type MissingDenomsWithHeight struct {
	Height int64    `json:"height"`
	Denoms []string `json:"denoms"`
}

// Price Inspector Configuration
type Configuration struct {
	Validators []ValidatorInfo `json:"validators"`
	Window     int64           `json:"window"`
	Networks   struct {
		Canon struct {
			RPC string `json:"rpc"`
			API string `json:"api"`
		} `json:"canon"`
		Mainnet struct {
			RPC string `json:"rpc"`
			API string `json:"api"`
		} `json:"mainnet"`
	} `json:"networks"`
}

type ValidatorInfo struct {
	Validator string `json:"validator"`
	Moniker   string `json:"moniker"`
}

// TxSearchResonse of rpc
type TxSearchResonse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Txs        []Tx   `json:"txs"`
		TotalCount string `json:"total_count"`
	} `json:"result"`
}

type Tx struct {
	Hash     string `json:"hash"`
	Height   string `json:"height"`
	Index    int    `json:"index"`
	TxResult struct {
		Code      int    `json:"code"`
		Data      string `json:"data"`
		Log       string `json:"log"`
		Info      string `json:"info"`
		GasWanted string `json:"gas_wanted"`
		GasUsed   string `json:"gas_used"`
		Events    []struct {
			Type       string `json:"type"`
			Attributes []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
				Index bool   `json:"index"`
			} `json:"attributes"`
		} `json:"events"`
		Codespace string `json:"codespace"`
	} `json:"tx_result"`
	Tx string `json:"tx"`
}

type StatusResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		NodeInfo struct {
			ProtocolVersion struct {
				P2P   string `json:"p2p"`
				Block string `json:"block"`
				App   string `json:"app"`
			} `json:"protocol_version"`
			ID         string `json:"id"`
			ListenAddr string `json:"listen_addr"`
			Network    string `json:"network"`
			Version    string `json:"version"`
			Channels   string `json:"channels"`
			Moniker    string `json:"moniker"`
			Other      struct {
				TxIndex    string `json:"tx_index"`
				RPCAddress string `json:"rpc_address"`
			} `json:"other"`
		} `json:"node_info"`
		SyncInfo struct {
			LatestBlockHash     string    `json:"latest_block_hash"`
			LatestAppHash       string    `json:"latest_app_hash"`
			LatestBlockHeight   string    `json:"latest_block_height"`
			LatestBlockTime     time.Time `json:"latest_block_time"`
			EarliestBlockHash   string    `json:"earliest_block_hash"`
			EarliestAppHash     string    `json:"earliest_app_hash"`
			EarliestBlockHeight string    `json:"earliest_block_height"`
			EarliestBlockTime   time.Time `json:"earliest_block_time"`
			CatchingUp          bool      `json:"catching_up"`
		} `json:"sync_info"`
		ValidatorInfo struct {
			Address string `json:"address"`
			PubKey  struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"pub_key"`
			VotingPower string `json:"voting_power"`
		} `json:"validator_info"`
	} `json:"result"`
}
