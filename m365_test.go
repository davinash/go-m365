package m365

import (
	"github.com/davinash/go-m365/client"
	"github.com/davinash/go-m365/outlook"
	"testing"
)

var (
	tenantId      = "NetVaultPlugIn2.onmicrosoft.com"
	applicationId = "0d373c7b-ce66-4f23-b617-27c4d7658d3d"
	securityKey   = "9z2L5FKpHvMz2Ql5~pi7-T-_5.66c7z_3V"
)

func TestBasic(t *testing.T) {
	t.Run("BasicTest", func(t *testing.T) {
		client := MClient{
			auth: client.Auth{
				TenantId:          tenantId,
				ApplicationId:     applicationId,
				ClientSecurityKey: securityKey,
			},
		}
		client.Outlook(outlook.Options{})
	})
}
