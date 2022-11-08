//go:build all || iam
// +build all iam

package iam

import "github.com/akamai/terraform-provider-akamai/v3/pkg/providers/registry"

func init() {
	registry.RegisterProvider(Subprovider())
}
