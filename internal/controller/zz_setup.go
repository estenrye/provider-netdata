// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	discordchannel "github.com/estenrye/provider-netdata/internal/controller/discordchannel/discordchannel"
	member "github.com/estenrye/provider-netdata/internal/controller/member/member"
	roommember "github.com/estenrye/provider-netdata/internal/controller/noderoommember/roommember"
	pagerdutychannel "github.com/estenrye/provider-netdata/internal/controller/pagerdutychannel/pagerdutychannel"
	providerconfig "github.com/estenrye/provider-netdata/internal/controller/providerconfig"
	room "github.com/estenrye/provider-netdata/internal/controller/room/room"
	memberroommember "github.com/estenrye/provider-netdata/internal/controller/roommember/member"
	space "github.com/estenrye/provider-netdata/internal/controller/space/space"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		discordchannel.Setup,
		member.Setup,
		roommember.Setup,
		pagerdutychannel.Setup,
		providerconfig.Setup,
		room.Setup,
		memberroommember.Setup,
		space.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
