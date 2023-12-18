package rpc

import (
	"github.com/cosmos/cosmos-sdk/types/module/testutil"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/umee-network/price_inspector/config"
	"github.com/umee-network/price_inspector/utils"
	"github.com/umee-network/umee/v6/app"
	otypes "github.com/umee-network/umee/v6/x/oracle/types"
)

// TxDecoder decodes a base64 encoded transaction data and returns a slice of ValidatorExgRates and an error if the decoding fails.
func TxDecoder(txData string, codec testutil.TestEncodingConfig) ([]config.ValidatorExgRates, error) {
	d, err := utils.DecodeBase64(txData)
	if err != nil {
		return nil, err
	}
	codec = app.MakeEncodingConfig()
	var c txtypes.Tx
	err = codec.Codec.Unmarshal(d, &c)
	if err != nil {
		return nil, err
	}

	var validatorExgRates []config.ValidatorExgRates
	for _, ms := range c.Body.Messages {
		if ms.TypeUrl == "/umee.oracle.v1.MsgAggregateExchangeRateVote" {
			var msg otypes.MsgAggregateExchangeRateVote
			err := msg.Unmarshal(ms.Value)
			if err != nil {
				return nil, err
			}
			exratetupes, err := otypes.ParseExchangeRateTuples(msg.ExchangeRates)
			if err != nil {
				return nil, err
			}
			validatorExgRates = append(validatorExgRates, config.ValidatorExgRates{
				Validator:      msg.Validator,
				ExgRatesTuples: exratetupes,
			})
		}
	}
	return validatorExgRates, err
}
