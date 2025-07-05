package space

import "github.com/crossplane/upjet/pkg/config"

// NewExternalName returns a new ExternalName for the Slack channel resource.
func NewExternalName() config.ExternalName {
	return config.NewExternalNameFrom(config.IdentifierFromProvider)
}

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netdata_space", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "space"

		r.ExternalName = NewExternalName()

		r.ExternalName = config.IdentifierFromProvider
	})
}
