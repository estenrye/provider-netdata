package member

import (
	"github.com/crossplane/upjet/pkg/config"
)

// NewExternalName returns a new ExternalName for the space member resource.
func NewExternalName() config.ExternalName {
	return config.NewExternalNameFrom(config.IdentifierFromProvider)
}

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netdata_space_member", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "member"

		r.ExternalName = NewExternalName()

		r.References["space_id"] = config.Reference{
			TerraformName: "netdata_space",
			Extractor:     "github.com/crossplane/upjet/pkg/resource.ExtractResourceID()",
		}
	})
}
