/*
Copyright 2024 Arm Limited

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
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeTagIfDigestExists(t *testing.T) {

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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, removeTagIfDigestExists(tt.imgName))
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				assert.Equal(t, tt.want, GetLatestImage(tt.imgName))
			})
		})
	}
}
