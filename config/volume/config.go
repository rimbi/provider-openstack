package volume

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p* config.Provider) {
	p.AddResourceConfigurator("openstack_blockstorage_volume_v3", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
        // this resource, which would be "openstack"
		r.ShortGroup = "volume"
        r.ExternalName = config.IdentifierFromProvider
	})
}
