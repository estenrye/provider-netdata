package member

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netdata_space_member", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "member"

		r.ExternalName = config.NewExternalNameFrom(config.IdentifierFromProvider)

		r.References["space_id"] = config.Reference{
			Type: "github.com/estenrye/provider-netdata/apis/space/v1alpha1.Space",
		}
	})
}
