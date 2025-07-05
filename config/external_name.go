/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/estenrye/provider-netdata/config/discordchannel"
	member "github.com/estenrye/provider-netdata/config/member"
	noderoommember "github.com/estenrye/provider-netdata/config/node_room_member"
	"github.com/estenrye/provider-netdata/config/pagerdutychannel"
	"github.com/estenrye/provider-netdata/config/room"
	roommember "github.com/estenrye/provider-netdata/config/room_member"
	"github.com/estenrye/provider-netdata/config/slackchannel"
	"github.com/estenrye/provider-netdata/config/space"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"netdata_notification_discord_channel":   discordchannel.NewExternalName(),
	"netdata_notification_pagerduty_channel": pagerdutychannel.NewExternalName(),
	"netdata_notification_slack_channel":     slackchannel.NewExternalName(),
	"netdata_node_room_member":               noderoommember.NewExternalName(),
	"netdata_room":                           room.NewExternalName(),
	"netdata_room_member":                    roommember.NewExternalName(),
	"netdata_space":                          space.NewExternalName(),
	"netdata_space_member":                   member.NewExternalName(),
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
