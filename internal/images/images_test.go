/*
Copyright 2026 Arm Limited

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package images

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveTagIfDigestExists(t *testing.T) {

	tests := []struct {
		name    string
		imgName string
		want    string
	}{
		{
			name:    "Both digest and tag exist",
			imgName: "testingImage:latest@sha256:1234567890abcdefghijklmnopqrstuvwxyzabcdefg",
			want:    "testingImage@sha256:1234567890abcdefghijklmnopqrstuvwxyzabcdefg",
		},
		{
			name:    "Only digest exists",
			imgName: "testingImage@sha256:1234567890abcdefghijklmnopqrstuvwxyzabcdefg",
			want:    "testingImage@sha256:1234567890abcdefghijklmnopqrstuvwxyzabcdefg",
		},
		{
			name:    "Only tag exists",
			imgName: "testingImage:latest",
			want:    "testingImage:latest",
		},
		{
			name:    "Empty string",
			imgName: "",
			want:    "",
		},
		{
			name:    "No tag or digest",
			imgName: "testingImage",
			want:    "testingImage",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, removeTagIfDigestExists(tt.imgName))
		})
	}
}

func TestGetFriendlyErrorMessage(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want string
	}{
		{
			name: "nil error",
			err:  nil,
			want: "",
		},
		{
			name: "authentication error",
			err:  fmt.Errorf("authentication required"),
			want: " Authentication Error. The private image could not be queried, please check the docker credentials are present and correct.",
		},
		{
			name: "auth error",
			err:  fmt.Errorf("auth failed"),
			want: " Authentication Error. The private image could not be queried, please check the docker credentials are present and correct.",
		},
		{
			name: "authorized error",
			err:  fmt.Errorf("not authorized"),
			want: " Authentication Error. The private image could not be queried, please check the docker credentials are present and correct.",
		},
		{
			name: "image not found",
			err:  fmt.Errorf("image not found in registry"),
			want: " Image not found. One or more pods are using an image that no longer exists.",
		},
		{
			name: "no image found",
			err:  fmt.Errorf("no image found"),
			want: " Image not found. One or more pods are using an image that no longer exists.",
		},
		{
			name: "manifest unknown",
			err:  fmt.Errorf("manifest unknown"),
			want: " Image not found. One or more pods are using an image that no longer exists.",
		},
		{
			name: "no such host",
			err:  fmt.Errorf("no such host"),
			want: " Communication error. Unable to communicate with the registry, please ensure the registry host is available.",
		},
		{
			name: "unknown error",
			err:  fmt.Errorf("something unexpected happened"),
			want: " An unknown error occurred. Please run in debug mode using the flag '-d' for more details.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, GetFriendlyErrorMessage(tt.err))
		})
	}
}

func TestCheckManifestListForArm64(t *testing.T) {
	dockerListMediaType := "application/vnd.docker.distribution.manifest.list.v2+json"
	ociIndexMediaType := "application/vnd.oci.image.index.v1+json"

	tests := []struct {
		name     string
		mimeType string
		manifest []byte
		want     bool
		wantErr  bool
	}{
		{
			name:     "Docker manifest list with arm64",
			mimeType: dockerListMediaType,
			manifest: []byte(`{"schemaVersion":2,"mediaType":"application/vnd.docker.distribution.manifest.list.v2+json","manifests":[{"mediaType":"application/vnd.docker.distribution.manifest.v2+json","digest":"sha256:aaa","platform":{"architecture":"amd64","os":"linux"}},{"mediaType":"application/vnd.docker.distribution.manifest.v2+json","digest":"sha256:bbb","platform":{"architecture":"arm64","os":"linux"}}]}`),
			want:     true,
		},
		{
			name:     "Docker manifest list without arm64",
			mimeType: dockerListMediaType,
			manifest: []byte(`{"schemaVersion":2,"mediaType":"application/vnd.docker.distribution.manifest.list.v2+json","manifests":[{"mediaType":"application/vnd.docker.distribution.manifest.v2+json","digest":"sha256:aaa","platform":{"architecture":"amd64","os":"linux"}}]}`),
			want:     false,
		},
		{
			name:     "Docker manifest list with arm but not arm64",
			mimeType: dockerListMediaType,
			manifest: []byte(`{"schemaVersion":2,"mediaType":"application/vnd.docker.distribution.manifest.list.v2+json","manifests":[{"mediaType":"application/vnd.docker.distribution.manifest.v2+json","digest":"sha256:aaa","platform":{"architecture":"arm","os":"linux"}}]}`),
			want:     false,
		},
		{
			name:     "Docker manifest list with arm64 windows (not linux)",
			mimeType: dockerListMediaType,
			manifest: []byte(`{"schemaVersion":2,"mediaType":"application/vnd.docker.distribution.manifest.list.v2+json","manifests":[{"mediaType":"application/vnd.docker.distribution.manifest.v2+json","digest":"sha256:aaa","platform":{"architecture":"arm64","os":"windows"}}]}`),
			want:     false,
		},
		{
			name:     "Docker manifest list invalid JSON",
			mimeType: dockerListMediaType,
			manifest: []byte(`not-valid-json`),
			wantErr:  true,
		},
		{
			name:     "OCI index with arm64",
			mimeType: ociIndexMediaType,
			manifest: []byte(`{"schemaVersion":2,"mediaType":"application/vnd.oci.image.index.v1+json","manifests":[{"mediaType":"application/vnd.oci.image.manifest.v1+json","digest":"sha256:aaa","platform":{"architecture":"amd64","os":"linux"}},{"mediaType":"application/vnd.oci.image.manifest.v1+json","digest":"sha256:bbb","platform":{"architecture":"arm64","os":"linux"}}]}`),
			want:     true,
		},
		{
			name:     "OCI index without arm64",
			mimeType: ociIndexMediaType,
			manifest: []byte(`{"schemaVersion":2,"mediaType":"application/vnd.oci.image.index.v1+json","manifests":[{"mediaType":"application/vnd.oci.image.manifest.v1+json","digest":"sha256:aaa","platform":{"architecture":"amd64","os":"linux"}}]}`),
			want:     false,
		},
		{
			name:     "OCI index invalid JSON",
			mimeType: ociIndexMediaType,
			manifest: []byte(`not-valid-json`),
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkManifestListForArm64(tt.manifest, tt.mimeType)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestGetLatestImage(t *testing.T) {
	tests := []struct {
		name    string
		imgName string
		want    string
	}{
		{
			name:    "Image with a tag",
			imgName: "nginx:v1",
			want:    "nginx:latest",
		},
		{
			name:    "Image with no tag",
			imgName: "nginx",
			want:    "nginx:latest",
		},
		{
			name:    "Image with no tag but with digest",
			imgName: "nginx@sha256:1114555554555",
			want:    "nginx:latest",
		},
		{
			name:    "Image with a tag and digest",
			imgName: "nginx:v2@sha256:1114555554555",
			want:    "nginx:latest",
		},
		{
			name:    "Empty string",
			imgName: "",
			want:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				assert.Equal(t, tt.want, GetLatestImage(tt.imgName))
			})
		})
	}
}
