package account

import (
	"github.com/qcserestipy/instagram-api-go-client/pkg/client"
	"github.com/qcserestipy/instagram-api-go-client/pkg/sdk/v24.0/page/client/access_token"
)

func GetPageAccessToken(params *access_token.GetPageAccessTokenParams) (*access_token.GetPageAccessTokenOK, error) {
	ctx, instagramClient, err := client.ContextWithClient()
	if err != nil {
		return nil, err
	}

	response, err := instagramClient.Page.AccessToken.GetPageAccessToken(ctx, params)
	if err != nil {
		return nil, err
	}
	return response, nil
}
