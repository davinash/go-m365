package client

import "time"

// Token struct holds the Microsoft Graph API authentication token used by GraphClient to authenticate API-requests to the ms graph API
type Token struct {
	TokenType   string    // should always be "Bearer" for msgraph API-calls
	NotBefore   time.Time // time when the access token starts to be valid
	ExpiresOn   time.Time // time when the access token expires
	Resource    string    // will most likely always be https://graph.microsoft.com, hence the BaseURL
	AccessToken string    // the access-token itself
}
