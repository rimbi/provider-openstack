/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	instancev2 "github.com/rimbi/provider-openstack/internal/controller/compute/instancev2"
	providerconfig "github.com/rimbi/provider-openstack/internal/controller/providerconfig"
	volumev3 "github.com/rimbi/provider-openstack/internal/controller/volume/volumev3"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		instancev2.Setup,
		providerconfig.Setup,
		volumev3.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
