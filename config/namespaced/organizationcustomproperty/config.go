package organizationcustomproperty

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure github_organization_custom_properties resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_organization_custom_properties", func(r *config.Resource) {
		r.Kind = "OrganizationCustomProperty"
		r.ShortGroup = "organization"
	})
}
