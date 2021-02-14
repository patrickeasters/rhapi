package rhapi

import (
	"context"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

const defaultOauthClient = "rhsm-api"
const defaultTokenURL = "https://sso.redhat.com/auth/realms/redhat-external/protocol/openid-connect/token"
const defaultTimeout = 60 * time.Second

// Client is used to interact with RHSM Subscription Management APIs
type Client struct {
	// Red Hat SSO offline refresh token, which can be obtained from https://access.redhat.com/management/api
	OfflineToken string
	// Max duration for HTTP requests using client
	Timeout time.Duration

	// OAuth token URL used for exchanging the refresh token
	TokenURL string
	// Public OAuth client ID used for refreshing the offline token
	ClientID string
}

// HTTPClient returns an HTTP client that can be used to query RHSM API and automatically includes the authorization context
func (c *Client) HTTPClient() *http.Client {
	if len(c.ClientID) == 0 {
		c.ClientID = defaultOauthClient
	}
	if len(c.TokenURL) == 0 {
		c.TokenURL = defaultTokenURL
	}

	conf := &oauth2.Config{
		ClientID: c.ClientID,
		Endpoint: oauth2.Endpoint{
			TokenURL: c.TokenURL,
		},
	}
	tok := &oauth2.Token{
		RefreshToken: c.OfflineToken,
	}

	return conf.Client(context.Background(), tok)
}
