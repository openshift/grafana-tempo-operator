package monolithic

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

// CommonLabels returns common labels for each TempoMonolithic object created by the operator.
func CommonLabels(instanceName string) map[string]string {
	return map[string]string{
		"app.kubernetes.io/name":       "tempo-monolithic",
		"app.kubernetes.io/instance":   instanceName,
		"app.kubernetes.io/managed-by": "tempo-operator",
	}
}

// ComponentLabels is a list of all commonLabels including the app.kubernetes.io/component:<component> label.
func ComponentLabels(component, instanceName string) labels.Set {
	return labels.Merge(CommonLabels(instanceName), map[string]string{
		"app.kubernetes.io/component": component,
	})
}

// ClusterScopedCommonLabels returns common labels for cluster-scoped resouces, for example ClusterRole.
func ClusterScopedCommonLabels(instance metav1.ObjectMeta) map[string]string {
	return labels.Merge(CommonLabels(instance.Name), map[string]string{
		"app.kubernetes.io/namespace": instance.Namespace,
	})
}

// ClusterScopedComponentLabels returns common labels for cluster-scoped resouces (e.g. ClusterRole)
// including the app.kubernetes.io/component:<component> label.
func ClusterScopedComponentLabels(instance metav1.ObjectMeta, component string) map[string]string {
	return labels.Merge(ClusterScopedCommonLabels(instance), map[string]string{
		"app.kubernetes.io/component": component,
	})
}
