package auth

import (
	"errors"
	"strings"
)

// Get API key form the headers of an HTTP request
func GetAPIKey(authHeader string) (string, error) {
	if authHeader == "" {
		return "", errors.New("no authentication into found")
	}
	vals := strings.Split(authHeader, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed authentication header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of authentication header")
	}

	return vals[1], nil
}
