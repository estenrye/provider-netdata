/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/estenrye/provider-netdata/config/discordchannel"
	"github.com/estenrye/provider-netdata/config/member"
	noderoommember "github.com/estenrye/provider-netdata/config/node_room_member"
	"github.com/estenrye/provider-netdata/config/pagerdutychannel"
	"github.com/estenrye/provider-netdata/config/room"
	roommember "github.com/estenrye/provider-netdata/config/room_member"
	"github.com/estenrye/provider-netdata/config/slackchannel"
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
		discordchannel.Configure,
		member.Configure,
		noderoommember.Configure,
		pagerdutychannel.Configure,
		room.Configure,
		roommember.Configure,
		slackchannel.Configure,
		space.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
