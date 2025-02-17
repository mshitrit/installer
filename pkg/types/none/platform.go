package none

import "github.com/openshift/installer/pkg/types/baremetal"

// FencingCredential stores the information about a baremetal host's management controller.
type FencingCredential struct {
	HostName string        `json:"hostName,omitempty" validate:"required,uniqueField"`
	BMC      baremetal.BMC `json:"bmc"`
}

// Platform stores any global configuration used for generic
// platforms.
type Platform struct {
	// FencingCredentials stores the information about a baremetal host's management controller.
	// +optional
	FencingCredentials []*FencingCredential `json:"fencingCredentials,omitempty"`
	//TODO mshitrit add a second PR which uses the credentials to create a Secret

}
