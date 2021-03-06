package claim

import (
	"fmt"

	"github.com/deis/steward/mode"
	"k8s.io/client-go/1.4/pkg/api/v1"
)

// ServicePlanClaimWrapper is a wrapper for a ServicePlanClaim that also contains kubernetes-specific information
type ServicePlanClaimWrapper struct {
	ObjectMeta v1.ObjectMeta
	Claim      *mode.ServicePlanClaim
}

func servicePlanClaimWrapperFromConfigMap(cm *v1.ConfigMap) (*ServicePlanClaimWrapper, error) {
	claim, err := mode.ServicePlanClaimFromMap(cm.Data)
	if err != nil {
		return nil, err
	}
	return &ServicePlanClaimWrapper{
		Claim: claim,
		ObjectMeta: v1.ObjectMeta{
			ResourceVersion: cm.ResourceVersion,
			Name:            cm.Name,
			Namespace:       cm.Namespace,
			Labels:          cm.Labels,
		},
	}, nil
}

// String is the fmt.Stringer implementation
func (spc ServicePlanClaimWrapper) String() string {
	return fmt.Sprintf("%s (resource %s)", *spc.Claim, spc.ObjectMeta.ResourceVersion)
}

func (spc ServicePlanClaimWrapper) toConfigMap() *v1.ConfigMap {
	return &v1.ConfigMap{
		ObjectMeta: spc.ObjectMeta,
		Data:       spc.Claim.ToMap(),
	}
}

// ServicePlanClaimsListWrapper is a wrapper for a list of ServicePlanClaims that also contains kubernetes-specific information.
type ServicePlanClaimsListWrapper struct {
	Claims          []*ServicePlanClaimWrapper
	ResourceVersion string
}
