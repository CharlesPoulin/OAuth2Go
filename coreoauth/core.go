package coreoauth

import (
    "golang.org/x/oauth2"
)

// OAuthConfig stores the configuration needed for OAuth
type OAuthConfig struct {
    Config *oauth2.Config
}

// NewOAuthConfig creates a new OAuth configuration
func NewOAuthConfig(clientID, clientSecret, redirectURL string, authURL, tokenURL string, scopes []string) *OAuthConfig {
    return &OAuthConfig{
        Config: &oauth2.Config{
            ClientID:     clientID,
            ClientSecret: clientSecret,
            RedirectURL:  redirectURL,
            Scopes:       scopes,
            Endpoint: oauth2.Endpoint{
                AuthURL:  authURL,
                TokenURL: tokenURL,
            },
        },
    }
}

// GetAuthURL generates an auth URL for the user to login
func (o *OAuthConfig) GetAuthURL(state string) string {
    return o.Config.AuthCodeURL(state)
}

// ExchangeCodeForToken exchanges a code for an access token
func (o *OAuthConfig) ExchangeCodeForToken(code string) (*oauth2.Token, error) {
    return o.Config.Exchange(oauth2.NoContext, code)
}
