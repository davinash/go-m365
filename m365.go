package m365

import (
	"github.com/davinash/go-m365/client"
	"github.com/davinash/go-m365/outlook"
)

type MClient struct {
	auth client.Auth
}

func (c *MClient) Outlook(op outlook.Options) (*outlook.Outlook, error) {
	return outlook.NewOutlook(c.auth, op)
}
