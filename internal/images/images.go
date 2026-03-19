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
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/containers/image/v5/image"
	"github.com/containers/image/v5/manifest"
	"github.com/containers/image/v5/transports/alltransports"
	"github.com/containers/image/v5/types"
)

const (
	archArm64        = "arm64"
	osLinux          = "linux"
	dockerTransport  = "docker://"
	latestTag        = ":latest"
	dockerConfigFile = ".docker/config.json"

	errMsgAuthError = " Authentication Error. The private image could not be queried, please check the docker credentials are present and correct."
	errMsgNotFound  = " Image not found. One or more pods are using an image that no longer exists."
	errMsgCommError = " Communication error. Unable to communicate with the registry, please ensure the registry host is available."
	errMsgUnknown   = " An unknown error occurred. Please run in debug mode using the flag '-d' for more details."
)

var (
	authErrorKeywords = []string{"authentication", "auth", "authorized"}
	notFoundKeywords  = []string{"no image found", "image not found", "manifest unknown"}
	commErrorKeywords = []string{"no such host"}
)

func containsAnyOf(input string, keywords []string) bool {
	for _, kw := range keywords {
		if strings.Contains(input, kw) {
			return true
		}
	}
	return false
}

func getDockerConfigPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, dockerConfigFile)
}

func GetFriendlyErrorMessage(err error) string {
	if err == nil {
		return ""
	}

	msg := err.Error()
	switch {
	case containsAnyOf(msg, authErrorKeywords):
		return errMsgAuthError
	case containsAnyOf(msg, notFoundKeywords):
		return errMsgNotFound
	case containsAnyOf(msg, commErrorKeywords):
		return errMsgCommError
	default:
		return errMsgUnknown
	}
}

// CheckLinuxArm64Support checks for the existence of an arm64 linux image in the manifest.
func CheckLinuxArm64Support(imgName string) (bool, error) {
	sys := &types.SystemContext{
		DockerCompatAuthFilePath: getDockerConfigPath(),
	}

	// Docker references with both a tag and digest are currently not supported.
	imgName = removeTagIfDigestExists(imgName)

	ref, err := alltransports.ParseImageName(dockerTransport + imgName)
	if err != nil {
		return false, fmt.Errorf("error parsing image name: %w", err)
	}

	src, err := ref.NewImageSource(context.Background(), sys)
	if err != nil {
		return false, fmt.Errorf("error getting image source: %w", err)
	}
	defer func() { _ = src.Close() }()

	rawManifest, mimeType, err := src.GetManifest(context.Background(), nil)
	if err != nil {
		return false, fmt.Errorf("error getting manifest: %w", err)
	}

	// For manifest lists, directly inspect the platform entries rather than
	// resolving via system context (which picks the host platform, not arm64).
	if manifest.MIMETypeIsMultiImage(mimeType) {
		return checkManifestListForArm64(rawManifest, mimeType)
	}

	// Single-platform image: inspect to check architecture.
	img, err := image.FromUnparsedImage(context.Background(), sys, image.UnparsedInstance(src, nil))
	if err != nil {
		return false, fmt.Errorf("error parsing manifest: %w", err)
	}

	imgInspect, err := img.Inspect(context.Background())
	if err != nil {
		return false, fmt.Errorf("error inspecting image: %w", err)
	}

	return imgInspect.Architecture == archArm64, nil
}

func checkManifestListForArm64(rawManifest []byte, mimeType string) (bool, error) {
	switch mimeType {
	case manifest.DockerV2ListMediaType:
		list, err := manifest.Schema2ListFromManifest(rawManifest)
		if err != nil {
			return false, fmt.Errorf("error parsing manifest list: %w", err)
		}
		for _, m := range list.Manifests {
			if m.Platform.Architecture == archArm64 && m.Platform.OS == osLinux {
				return true, nil
			}
		}
	default: // OCI image index
		idx, err := manifest.OCI1IndexFromManifest(rawManifest)
		if err != nil {
			return false, fmt.Errorf("error parsing OCI index: %w", err)
		}
		for _, m := range idx.Manifests {
			if m.Platform != nil && m.Platform.Architecture == archArm64 && m.Platform.OS == osLinux {
				return true, nil
			}
		}
	}
	return false, nil
}

// removeTagIfDigestExists strips the tag from an image reference that has both a tag and a digest,
// since the containers/image library does not support such references.
func removeTagIfDigestExists(imgName string) string {
	if imgName == "" {
		return imgName
	}

	if strings.Contains(imgName, "@") {
		parts := strings.SplitN(imgName, "@", 2)
		if strings.Contains(parts[0], ":") {
			parts[0] = strings.SplitN(parts[0], ":", 2)[0]
		}
		return parts[0] + "@" + parts[1]
	}
	return imgName
}

// GetLatestImage returns the image reference with all tags/digests replaced by ":latest".
func GetLatestImage(imgName string) string {
	if imgName == "" {
		return imgName
	}

	// Trim everything from the first '@' or ':'.
	base := strings.FieldsFunc(imgName, func(c rune) bool {
		return c == '@' || c == ':'
	})[0]

	return base + latestTag
}
