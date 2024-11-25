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
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/containers/image/v5/image"
	"github.com/containers/image/v5/transports/alltransports"
	"github.com/containers/image/v5/types"
)

func containsAnyOf(input string, suggestions []string) bool {
	for _, suggestion := range suggestions {
		if strings.Contains(input, suggestion) {
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
	return filepath.Join(home, ".docker", "config.json")
}

func GetFriendlyErrorMessage(err error, pods []string) string {
	if err == nil {
		return ""
	}

	errorMessage := err.Error()
	switch {
	case containsAnyOf(errorMessage, []string{"authentication", "auth", "authorized"}):
		return " Authentication error. A private image could not be checked, check the docker credentials are present and up to date."
	case containsAnyOf(errorMessage, []string{"no image found", "image not found"}):
		return " Image not found. Some pods like `" + pods[0] + "` are using an image that no longer exists."
	case containsAnyOf(errorMessage, []string{"no such host"}):
		return " communication error. Could not communication with the registry, make sure the registry host exists."
	default:
		return " An unknown error occurred. Please run in debug mode using the flag '-d' for more details."
	}
}

// CheckLinuxArm64Support checks for the existance of an arm64 linux image in the manifest
func CheckLinuxArm64Support(imgName string) (bool, error) {
	sys := &types.SystemContext{
		ArchitectureChoice:       "arm64",
		OSChoice:                 "linux",
		DockerCompatAuthFilePath: getDockerConfigPath(),
	}

	// Docker references with both a tag and digest are currently not supported
	imgName = removeTagIfDigestExists(imgName)

	ref, err := alltransports.ParseImageName(fmt.Sprintf("docker://%s", imgName))
	if err != nil {
		return false, fmt.Errorf("error parsing image name: %w", err)
	}

	src, err := ref.NewImageSource(context.Background(), sys)
	if err != nil {
		return false, fmt.Errorf("error getting image source: %w", err)
	}
	defer src.Close()

	img, err := image.FromUnparsedImage(context.TODO(), sys, image.UnparsedInstance(src, nil))
	if err != nil {
		return false, fmt.Errorf("error parsing manifest: %w", err)
	}

	imgInspect, err := img.Inspect(context.TODO())
	if err != nil {
		return false, fmt.Errorf("error inspecting image: %w", err)
	}

	return imgInspect.Architecture == "arm64", nil
}

func removeTagIfDigestExists(imgName string) string {
	// check for empty string
	if imgName == "" {
		return imgName
	}

	if strings.Contains(imgName, "@") {
		parts := strings.Split(imgName, "@")
		// Check if the first part contains a colon, indicating a tag
		if strings.Contains(parts[0], ":") {
			subParts := strings.Split(parts[0], ":")
			// Reconstruct the image name without the tag
			imgName = subParts[0] + "@" + parts[1]
		}
	}
	return imgName
}

func GetLatestImage(imgName string) string {

	// check for empty string
	if imgName == "" {
		return imgName
	}

	// Remove everything after '@' or ':'
	parts := strings.FieldsFunc(imgName, func(c rune) bool {
		return c == '@' || c == ':'
	})

	return parts[0] + ":latest"
}
