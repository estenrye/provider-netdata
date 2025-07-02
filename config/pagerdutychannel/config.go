package pagerdutychannel

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netdata_notification_pagerduty_channel", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "pagerdutychannel"

		r.References["space_id"] = config.Reference{
			Type: "github.com/estenrye/provider-netdata/apis/space/v1alpha1.Space",
		}

		r.References["room_ids"] = config.Reference{
			Type: "github.com/estenrye/provider-netdata/apis/room/v1alpha1.Room",
		}
	})
}
