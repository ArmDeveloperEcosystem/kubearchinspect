package images

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/containers/image/v5/manifest"
	"github.com/containers/image/v5/transports/alltransports"
	"github.com/containers/image/v5/types"
)

func getDockerConfigPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".docker", "config.json")
}

// CheckLinuxArm64Support checks for the existance of an arm64 linux image in the manifest
func CheckLinuxArm64Support(imgName string) (bool, error) {
	sys := &types.SystemContext{
		ArchitectureChoice:       "arm64",
		OSChoice:                 "linux",
		DockerCompatAuthFilePath: getDockerConfigPath(),
	}

	ref, err := alltransports.ParseImageName(fmt.Sprintf("docker://%s", imgName))
	if err != nil {
		return false, fmt.Errorf("error parsing image name: %w", err)
	}

	imageSource, err := ref.NewImageSource(context.Background(), sys)
	if err != nil {
		return false, fmt.Errorf("error getting image source: %w", err)
	}
	defer imageSource.Close()

	rawManifest, mimeType, err := imageSource.GetManifest(context.Background(), nil)
	if err != nil {
		return false, fmt.Errorf("error getting manifest: %w", err)
	}

	if manifest.MIMETypeIsMultiImage(mimeType) {
		manifestList, err := manifest.ListFromBlob(rawManifest, mimeType)
		if err != nil {
			return false, err
		}

		// This call will error if it cannot find a instance that supports linux arm64
		_, err = manifestList.ChooseInstance(sys)
		return err == nil, nil
	} else {
		// m, err := manifest.FromBlob(rawManifest, mimeType)
		// if err != nil {
		// 	return false, nil
		// }
		// mInfo, err := m.Inspect(nil)
		// if err != nil {
		// 	return false, nil
		// }
		// return mInfo.Architecture == "arm64", nil
		return false, fmt.Errorf("image manifests not supported")
	}
}

func CheckLatestLinuxArm64Support(imgName string) (bool, error) {
	split := strings.Split(imgName, ":")
	if len(split) < 2 {
		return false, fmt.Errorf("invalid image name")
	}
	latestImageName := split[0] + ":latest"
	return CheckLinuxArm64Support(latestImageName)
}
