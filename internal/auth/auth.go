package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization");
	if val == "" {
		return "", errors.New("no authentication info found")
	}

	vals := strings.Split(val, " ");
	if len(vals) != 2 {
		return "", errors.New("invalid authentication info format")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("invalid authentication type, expected 'ApiKey'")
	}
	return vals[1], nil
}