package config

import (
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
			RPC  string `json:"rpc"`
			GRPC string `json:"grpc"`
		} `json:"canon"`
		Mainnet struct {
			RPC  string `json:"rpc"`
			GRPC string `json:"grpc"`
		} `json:"mainnet"`
	} `json:"networks"`
}

type ValidatorInfo struct {
	Validator string `json:"validator"`
	Moniker   string `json:"moniker"`
}
