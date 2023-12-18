package utils

import "encoding/base64"

// DecodeBase64 decodes a base64 encoded string into a byte slice.
// It returns the decoded byte slice and an error if the decoding fails.
func DecodeBase64(data string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}
