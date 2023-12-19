package utils

// GetMissingDenoms returns the missing denominations by comparing the accepted denominations with the submited denoms.
func GetMissingDenoms(accepedDenoms, extRateDenoms []string) []string {
	var missingDenoms []string
	for _, ad := range accepedDenoms {
		if !contains(extRateDenoms, ad) {
			missingDenoms = append(missingDenoms, ad)
		}
	}

	return missingDenoms
}

func contains(denoms []string, denom string) bool {
	for _, d := range denoms {
		if d == denom {
			return true
		}
	}
	return false
}
