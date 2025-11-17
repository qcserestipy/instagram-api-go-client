package config

import (
	"fmt"
	"net/url"
	"os"
	"sync"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
)

var (
	accessToken     string
	accessTokenOnce sync.Once
	accessTokenErr  error

	clientSecret     string
	clientSecretOnce sync.Once
	clientSecretErr  error

	pageID     string
	pageIDOnce sync.Once
	pageIDErr  error
)

// GetAccessToken retrieves the access token from environment variable
func GetAccessToken() (string, error) {
	accessTokenOnce.Do(func() {
		accessToken = os.Getenv("ACCESS_TOKEN")
		if accessToken == "" {
			accessTokenErr = fmt.Errorf("ACCESS_TOKEN environment variable is not set")
		}
	})
	return accessToken, accessTokenErr
}

// GetClientSecret retrieves the client secret from environment variable
func GetClientSecret() (string, error) {
	clientSecretOnce.Do(func() {
		clientSecret = os.Getenv("CLIENT_SECRET")
		if clientSecret == "" {
			clientSecretErr = fmt.Errorf("CLIENT_SECRET environment variable is not set")
		}
	})
	return clientSecret, clientSecretErr
}

// GetPageID retrieves the Facebook Page ID from environment variable
func GetPageID() (string, error) {
	pageIDOnce.Do(func() {
		pageID = os.Getenv("PAGE_ID")
		if pageID == "" {
			pageIDErr = fmt.Errorf("PAGE_ID environment variable is not set")
		}
	})
	return pageID, pageIDErr
}

// CreateClientConfig creates a common configuration for Instagram API clients
func CreateClientConfig() (*url.URL, runtime.ClientAuthInfoWriter, error) {
	token, err := GetAccessToken()
	if err != nil {
		return nil, nil, err
	}

	apiURL := &url.URL{
		Scheme: "https",
		Host:   "graph.facebook.com",
		Path:   "/v24.0",
	}

	authInfo := httptransport.APIKeyAuth("access_token", "query", token)

	return apiURL, authInfo, nil
}
