package cloudcredentials

import (
	"github.com/ViaQ/logerr/v2/kverrors"
	cloudcredentialv1 "github.com/openshift/cloud-credential-operator/pkg/apis/cloudcredential/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/grafana/tempo-operator/internal/manifests/manifestutils"
)

// BuildCredentialsRequest return new CCO object.
func BuildCredentialsRequest(obj metav1.Object, serviceAccount string, env *manifestutils.TokenCCOAuthConfig) ([]client.Object, error) {

	stack := client.ObjectKey{Name: obj.GetName(), Namespace: obj.GetNamespace()}

	providerSpec, err := encodeProviderSpec(env)
	if err != nil {
		return nil, kverrors.Wrap(err, "failed encoding credentialsrequest provider spec")
	}

	credentialRequest := &cloudcredentialv1.CredentialsRequest{
		ObjectMeta: metav1.ObjectMeta{
			Name:      stack.Name,
			Namespace: stack.Namespace,
		},
		Spec: cloudcredentialv1.CredentialsRequestSpec{
			SecretRef: corev1.ObjectReference{
				Name:      manifestutils.ManagedCredentialsSecretName(stack.Name),
				Namespace: stack.Namespace,
			},
			ProviderSpec: providerSpec,
			ServiceAccountNames: []string{
				serviceAccount,
			},
			CloudTokenPath: manifestutils.ServiceAccountTokenFilePath,
		},
	}

	return []client.Object{credentialRequest}, nil
}

func encodeProviderSpec(env *manifestutils.TokenCCOAuthConfig) (*runtime.RawExtension, error) {
	var spec runtime.Object

	if env.AWS != nil {
		spec = &cloudcredentialv1.AWSProviderSpec{
			StatementEntries: []cloudcredentialv1.StatementEntry{
				{
					Action: []string{
						"s3:ListBucket",
						"s3:PutObject",
						"s3:GetObject",
						"s3:DeleteObject",
					},
					Effect:   "Allow",
					Resource: "arn:aws:s3:*:*:*",
				},
			},
			STSIAMRoleARN: env.AWS.RoleARN,
		}
	}

	if spec != nil {
		encodedSpec, err := cloudcredentialv1.Codec.EncodeProviderSpec(spec.DeepCopyObject())
		return encodedSpec, err
	} else {
		return nil, kverrors.New("unsupported CCO environment")
	}
}
