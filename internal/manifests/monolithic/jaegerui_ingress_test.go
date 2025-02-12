package monolithic

import (
	"errors"
	"testing"

	routev1 "github.com/openshift/api/route/v1"
	"github.com/stretchr/testify/require"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/controller-runtime/pkg/client"

	configv1alpha1 "github.com/grafana/tempo-operator/api/config/v1alpha1"
	"github.com/grafana/tempo-operator/api/tempo/v1alpha1"
	"github.com/grafana/tempo-operator/internal/manifests/manifestutils"
)

func TestBuildJaegerUIIngress(t *testing.T) {
	opts := Options{
		Tempo: v1alpha1.TempoMonolithic{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "sample",
				Namespace: "default",
			},
		},
	}
	labels := ComponentLabels(manifestutils.JaegerUIComponentName, "sample")

	tests := []struct {
		name     string
		input    v1alpha1.TempoMonolithicSpec
		expected client.Object
	}{
		{
			name: "ingress",
			input: v1alpha1.TempoMonolithicSpec{
				JaegerUI: &v1alpha1.MonolithicJaegerUISpec{
					Enabled: true,
					Ingress: &v1alpha1.MonolithicJaegerUIIngressSpec{
						Enabled: true,
					},
				},
			},
			expected: &networkingv1.Ingress{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "networking.k8s.io/v1",
					Kind:       "Ingress",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "tempo-sample-jaegerui",
					Namespace: "default",
					Labels:    labels,
				},
				Spec: networkingv1.IngressSpec{
					DefaultBackend: &networkingv1.IngressBackend{
						Service: &networkingv1.IngressServiceBackend{
							Name: "tempo-sample-jaegerui",
							Port: networkingv1.ServiceBackendPort{
								Name: "jaeger-ui",
							},
						},
					},
				},
			},
		},
		{
			name: "ingress with host",
			input: v1alpha1.TempoMonolithicSpec{
				JaegerUI: &v1alpha1.MonolithicJaegerUISpec{
					Enabled: true,
					Ingress: &v1alpha1.MonolithicJaegerUIIngressSpec{
						Enabled: true,
						Host:    "abc",
					},
				},
			},
			expected: &networkingv1.Ingress{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "networking.k8s.io/v1",
					Kind:       "Ingress",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "tempo-sample-jaegerui",
					Namespace: "default",
					Labels:    labels,
				},
				Spec: networkingv1.IngressSpec{
					Rules: []networkingv1.IngressRule{
						{
							Host: "abc",
							IngressRuleValue: networkingv1.IngressRuleValue{
								HTTP: &networkingv1.HTTPIngressRuleValue{
									Paths: []networkingv1.HTTPIngressPath{
										{
											Path:     "/",
											PathType: ptr.To(networkingv1.PathTypePrefix),
											Backend: networkingv1.IngressBackend{
												Service: &networkingv1.IngressServiceBackend{
													Name: "tempo-sample-jaegerui",
													Port: networkingv1.ServiceBackendPort{
														Name: "jaeger-ui",
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			opts.Tempo.Spec = test.input
			opts.Tempo.Default(configv1alpha1.ProjectConfig{})

			obj := BuildJaegerUIIngress(opts)
			require.Equal(t, test.expected, obj)
		})
	}
}

func TestBuildJaegerUIRoute(t *testing.T) {
	opts := Options{
		Tempo: v1alpha1.TempoMonolithic{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "sample",
				Namespace: "default",
			},
		},
	}
	labels := ComponentLabels(manifestutils.JaegerUIComponentName, "sample")

	tests := []struct {
		name        string
		input       v1alpha1.TempoMonolithicSpec
		expected    *routev1.Route
		expectedErr error
	}{
		{
			name: "route",
			input: v1alpha1.TempoMonolithicSpec{
				JaegerUI: &v1alpha1.MonolithicJaegerUISpec{
					Enabled: true,
					Route: &v1alpha1.MonolithicJaegerUIRouteSpec{
						Enabled:     true,
						Termination: "edge",
					},
				},
			},
			expected: &routev1.Route{
				TypeMeta: metav1.TypeMeta{
					APIVersion: networkingv1.SchemeGroupVersion.String(),
					Kind:       "Ingress",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "tempo-sample-jaegerui",
					Namespace: "default",
					Labels:    labels,
					Annotations: map[string]string{
						"haproxy.router.openshift.io/timeout": "30s",
					},
				},
				Spec: routev1.RouteSpec{
					Host: "",
					To: routev1.RouteTargetReference{
						Kind: "Service",
						Name: "tempo-sample-jaegerui",
					},
					Port: &routev1.RoutePort{
						TargetPort: intstr.FromString("jaeger-ui"),
					},
					TLS: &routev1.TLSConfig{Termination: routev1.TLSTerminationEdge},
				},
			},
		},
		{
			name: "route with invalid tls",
			input: v1alpha1.TempoMonolithicSpec{
				JaegerUI: &v1alpha1.MonolithicJaegerUISpec{
					Enabled: true,
					Route: &v1alpha1.MonolithicJaegerUIRouteSpec{
						Enabled:     true,
						Termination: "invalid",
					},
				},
			},
			expected:    nil,
			expectedErr: errors.New("unsupported tls termination 'invalid' specified for route"),
		},
		{
			name: "route with gateway",
			input: v1alpha1.TempoMonolithicSpec{
				Multitenancy: &v1alpha1.MonolithicMultitenancySpec{
					Enabled: true,
					TenantsSpec: v1alpha1.TenantsSpec{
						Authentication: []v1alpha1.AuthenticationSpec{
							{
								TenantName: "dev",
								TenantID:   "dev",
							},
						},
					},
				},
				JaegerUI: &v1alpha1.MonolithicJaegerUISpec{
					Enabled: true,
					Route: &v1alpha1.MonolithicJaegerUIRouteSpec{
						Enabled: true,
					},
				},
			},
			expected: &routev1.Route{
				TypeMeta: metav1.TypeMeta{
					APIVersion: networkingv1.SchemeGroupVersion.String(),
					Kind:       "Ingress",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "tempo-sample-jaegerui",
					Namespace: "default",
					Labels:    labels,
					Annotations: map[string]string{
						"haproxy.router.openshift.io/timeout": "30s",
					},
				},
				Spec: routev1.RouteSpec{
					Host: "",
					To: routev1.RouteTargetReference{
						Kind: "Service",
						Name: "tempo-sample-gateway",
					},
					Port: &routev1.RoutePort{
						TargetPort: intstr.FromString("public"),
					},
					TLS: &routev1.TLSConfig{Termination: routev1.TLSTerminationEdge},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			opts.Tempo.Spec = test.input
			opts.Tempo.Default(configv1alpha1.ProjectConfig{})

			obj, err := BuildJaegerUIRoute(opts)
			require.Equal(t, test.expectedErr, err)
			require.Equal(t, test.expected, obj)
		})
	}
}
