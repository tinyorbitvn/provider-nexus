package clients

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	namespacedv1beta1 "github.com/tinyorbitvn/provider-nexus/apis/namespaced/v1beta1"
)

func TestBuildConfiguration(t *testing.T) {
	delay := int32(1500)
	cases := map[string]struct {
		spec     namespacedv1beta1.ProviderConfigSpec
		password string
		want     map[string]any
	}{
		"minimal": {
			spec:     namespacedv1beta1.ProviderConfigSpec{URL: "http://nexus.nexus.svc:8081", Username: "admin"},
			password: "s3cret",
			want:     map[string]any{"url": "http://nexus.nexus.svc:8081", "username": "admin", "password": "s3cret"},
		},
		"with delay": {
			spec:     namespacedv1beta1.ProviderConfigSpec{URL: "http://n:8081", Username: "u", ClusterStabilisationDelayMs: &delay},
			password: "p",
			want:     map[string]any{"url": "http://n:8081", "username": "u", "password": "p", "cluster_stabilisation_delay_ms": int32(1500)},
		},
		"password with trailing newline is trimmed": {
			spec:     namespacedv1beta1.ProviderConfigSpec{URL: "http://n:8081", Username: "u"},
			password: "p\n",
			want:     map[string]any{"url": "http://n:8081", "username": "u", "password": "p"},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := buildConfiguration(&tc.spec, tc.password)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("buildConfiguration() -want +got:\n%s", diff)
			}
		})
	}
}
