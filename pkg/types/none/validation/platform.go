package validation

import (
	bmvalidation "github.com/openshift/installer/pkg/types/baremetal/validation"
	"github.com/openshift/installer/pkg/types/none"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

// ValidateProvisioningNetworkDisabledSupported validates hosts bmc address support provisioning network is disabled
func ValidateProvisioningNetworkDisabledSupported(fencingCredentials []*none.FencingCredential, fldPath *field.Path) (errors field.ErrorList) {
	for idx, fencingCredential := range fencingCredentials {
		if err := bmvalidation.ValidateSingleProvisioningNetworkDisabledSupported(fencingCredential.BMC, fldPath, idx); err != nil {
			errors = append(errors, err)
		}
	}
	return
}
