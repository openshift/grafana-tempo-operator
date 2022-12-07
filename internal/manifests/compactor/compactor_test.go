package compactor

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8slabels "k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/intstr"

	"github.com/os-observability/tempo-operator/apis/tempo/v1alpha1"
	"github.com/os-observability/tempo-operator/internal/manifests/manifestutils"
	"github.com/os-observability/tempo-operator/internal/manifests/memberlist"
)

func TestBuildCompactor(t *testing.T) {
	objects, err := BuildCompactor(v1alpha1.Microservices{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test",
			Namespace: "project1",
		},
		Spec: v1alpha1.MicroservicesSpec{
			Images: v1alpha1.ImagesSpec{
				Tempo: "docker.io/grafana/tempo:1.5.0",
			},
			ServiceAccount: "tempo-test-serviceaccount",
			Components: v1alpha1.TempoComponentsSpec{
				Compactor: &v1alpha1.TempoComponentSpec{
					NodeSelector: map[string]string{"a": "b"},
					Tolerations: []corev1.Toleration{
						{
							Key: "c",
						},
					},
				},
			},
			Resources: v1alpha1.Resources{
				Total: &corev1.ResourceRequirements{
					Limits: corev1.ResourceList{
						corev1.ResourceCPU:    resource.MustParse("1000m"),
						corev1.ResourceMemory: resource.MustParse("2Gi"),
					},
				},
			},
		},
	})
	require.NoError(t, err)

	labels := manifestutils.ComponentLabels("compactor", "test")
	assert.Equal(t, 2, len(objects))

	assert.Equal(t, &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "tempo-test-compactor",
			Namespace: "project1",
			Labels:    labels,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name:       "http-memberlist",
					Protocol:   corev1.ProtocolTCP,
					Port:       memberlist.PortMemberlist,
					TargetPort: intstr.FromInt(memberlist.PortMemberlist),
				},
				{
					Name:       httpPortName,
					Protocol:   corev1.ProtocolTCP,
					Port:       portHTTPServer,
					TargetPort: intstr.FromString(httpPortName),
				},
			},
			Selector: labels,
		},
	}, objects[1])

	assert.Equal(t, &v1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: v1.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "tempo-test-compactor",
			Namespace: "project1",
			Labels:    labels,
		},
		Spec: v1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: k8slabels.Merge(labels, map[string]string{"tempo-gossip-member": "true"}),
				},
				Spec: corev1.PodSpec{
					ServiceAccountName: "tempo-test-serviceaccount",
					NodeSelector:       map[string]string{"a": "b"},
					Tolerations: []corev1.Toleration{
						{
							Key: "c",
						},
					},
					Containers: []corev1.Container{
						{
							Name:  "tempo",
							Image: "docker.io/grafana/tempo:1.5.0",
							Args:  []string{"-target=compactor", "-config.file=/conf/tempo.yaml"},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      configVolumeName,
									MountPath: "/conf",
									ReadOnly:  true,
								},
							},
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									ContainerPort: portHTTPServer,
									Protocol:      corev1.ProtocolTCP,
								},
								{
									Name:          "http-memberlist",
									ContainerPort: 7946,
									Protocol:      corev1.ProtocolTCP,
								},
							},
							Resources: corev1.ResourceRequirements{
								Limits: corev1.ResourceList{
									corev1.ResourceCPU:    *resource.NewMilliQuantity(160, resource.BinarySI),
									corev1.ResourceMemory: *resource.NewQuantity(386547072, resource.BinarySI),
								},
								Requests: corev1.ResourceList{
									corev1.ResourceCPU:    *resource.NewMilliQuantity(48, resource.BinarySI),
									corev1.ResourceMemory: *resource.NewQuantity(115964128, resource.BinarySI),
								},
							},
						},
					},
					Volumes: []corev1.Volume{
						{
							Name: configVolumeName,
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "tempo-test",
									},
								},
							},
						},
					},
				},
			},
		},
	}, objects[0])
}