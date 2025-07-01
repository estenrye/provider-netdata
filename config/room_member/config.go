package roommember

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netdata_room_member", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "roommember"

		r.References["space_id"] = config.Reference{
			Type: "github.com/estenrye/provider-netdata/apis/space/v1alpha1.Space",
		}

		r.References["space_member_id"] = config.Reference{
			Type: "github.com/estenrye/provider-netdata/apis/member/v1alpha1.Member",
		}

		r.References["room_id"] = config.Reference{
			Type: "github.com/estenrye/provider-netdata/apis/room/v1alpha1.Room",
		}
	})
}
