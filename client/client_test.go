package client

import "testing"

func TestClient(t *testing.T) {
	t.Run("EmptyTenantId", func(t *testing.T) {
		_, err := NewClient(Auth{
			TenantId:          "",
			ApplicationId:     "",
			ClientSecurityKey: "",
		})
		if err == nil {
			t.Fatal(err)
		}
	})
	t.Run("ValidTenantId", func(t *testing.T) {
		_, err := NewClient(Auth{
			TenantId:          "NetVaultPlugIn2.onmicrosoft.com",
			ApplicationId:     "0d373c7b-ce66-4f23-b617-27c4d7658d3d",
			ClientSecurityKey: "9z2L5FKpHvMz2Ql5~pi7-T-_5.66c7z_3V",
		})
		if err != nil {
			t.Fatal(err)
		}
	})
}
