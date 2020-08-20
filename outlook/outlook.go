package outlook

import (
	"github.com/davinash/go-m365/client"
)

type Options struct {
	UseMaxPages bool
}

type Outlook struct {
	client  *client.Client
	options Options
}

func NewOutlook(c client.Auth, op Options) (*Outlook, error) {
	o := &Outlook{
		client:  nil,
		options: Options{UseMaxPages: op.UseMaxPages},
	}
	client, err := client.NewClient(c)
	if err != nil {
		return nil, err
	}
	o.client = client
	return o, nil
}
