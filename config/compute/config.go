package compute

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p* config.Provider) {
	p.AddResourceConfigurator("openstack_compute_instance_v2", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
        // this resource, which would be "openstack"
		r.ShortGroup = "compute"
        r.ExternalName = config.IdentifierFromProvider
	})
}
