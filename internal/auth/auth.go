package auth

import (
	"errors"
	"net/http"
	"strings"
)

func ExtractApiKey(h http.Header) (string, error) {
	authHeaderString := h.Get("Authorization")
	if authHeaderString == "" {
		return "", errors.New("Unauthorized invalid token")
	}
	authHeaderSplit := strings.Split(authHeaderString, " ")
	if len(authHeaderSplit) < 2 || authHeaderSplit[0] != "ApiKey" {
		return "", errors.New("Unauthorized invalid token")
	}
	return authHeaderSplit[1], nil
}
