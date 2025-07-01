/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/estenrye/provider-netdata/config/member"
	"github.com/estenrye/provider-netdata/config/room"
	"github.com/estenrye/provider-netdata/config/space"
)

const (
	resourcePrefix = "netdata"
	modulePath     = "github.com/estenrye/provider-netdata"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("netdata.provider.crossplane.rye.ninja"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		member.Configure,
		room.Configure,
		space.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
