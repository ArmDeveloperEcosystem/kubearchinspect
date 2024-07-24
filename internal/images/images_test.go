package images_test

import (
	"testing"

	"github.com/Arm-Debug/armer/internal/images"
	"github.com/stretchr/testify/assert"
)

func TestToFullURL(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"quay":                     {input: "quay.io/prometheus/alertmanager:v0.27.0", want: "quay.io/prometheus/alertmanager:v0.27.0"},
		"public ecr":               {input: "public.ecr.aws/eks/aws-load-balancer-controller:v2.5.2", want: "public.ecr.aws/eks/aws-load-balancer-controller:v2.5.2"},
		"ghcr":                     {input: "ghcr.io/dexidp/dex:v2.38.0", want: "ghcr.io/dexidp/dex:v2.38.0"},
		"dockerhub library":        {input: "nginx:latest", want: "docker.io/library/nginx:latest"},
		"dockerhub namespace":      {input: "envoyproxy/envoy:latest", want: "docker.io/envoyproxy/envoy:latest"},
		"dockerhub full library":   {input: "docker.io/library/nginx:latest", want: "docker.io/library/nginx:latest"},
		"dockerhub full namespace": {input: "docker.io/envoyproxy/envoy:latest", want: "docker.io/envoyproxy/envoy:latest"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := images.ToFullURL(tc.input)
			assert.Equal(t, tc.want, got)
		})
	}
}
