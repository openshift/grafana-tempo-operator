package gateway

import (
	"context"
	"fmt"

	configv1 "github.com/openshift/api/config/v1"
	operatorv1 "github.com/openshift/api/operator/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/grafana/tempo-operator/api/tempo/v1alpha1"
	"github.com/grafana/tempo-operator/internal/status"
)

// GetOpenShiftBaseDomain returns base domain of OCP cluster.
func GetOpenShiftBaseDomain(ctx context.Context, k8sClient client.Client) (string, error) {
	ictrl := &operatorv1.IngressController{}
	err := k8sClient.Get(ctx, types.NamespacedName{Name: "default", Namespace: "openshift-ingress-operator"}, ictrl)
	if err != nil {
		// The preferred way to get the base domain is via OCP ingress controller
		// this approach works well with CRC and normal OCP cluster.
		// Fallback on cluster DNS might not work with CRC because CRC uses .apps-crc. and not .apps.
		var clusterDNS configv1.DNS
		if err := k8sClient.Get(ctx, client.ObjectKey{Name: "cluster"}, &clusterDNS); err != nil {
			if apierrors.IsNotFound(err) {
				return "", &status.ConfigurationError{
					Message: "Missing OpenShift IngressController and cluster DNS configuration to read base domain",
					Reason:  v1alpha1.ReasonCouldNotGetOpenShiftBaseDomain,
				}
			}
			return "", fmt.Errorf("failed to lookup gateway base domain: %w", err)
		}

		return clusterDNS.Spec.BaseDomain, nil
	}
	return ictrl.Status.Domain, nil
}
