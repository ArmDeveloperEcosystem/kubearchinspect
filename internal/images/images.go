package images

import (
	"context"
	"fmt"

	"github.com/containers/image/v5/manifest"
	"github.com/containers/image/v5/transports/alltransports"
	"github.com/containers/image/v5/types"
)

func DoesSupportLinuxArm64(imgName string) (bool, error) {
	sys := &types.SystemContext{
		ArchitectureChoice: "arm64",
		OSChoice:           "linux",
	}

	ref, err := alltransports.ParseImageName(fmt.Sprintf("docker://%s", imgName))
	if err != nil {
		return false, err
	}

	imageSource, err := ref.NewImageSource(context.Background(), sys)
	if err != nil {
		return false, err
	}
	defer imageSource.Close()

	rawManifest, mimeType, err := imageSource.GetManifest(context.Background(), nil)
	if err != nil {
		return false, err
	}

	manifestList, err := manifest.ListFromBlob(rawManifest, mimeType)
	if err != nil {
		return false, err
	}

	_, err = manifestList.ChooseInstance(sys)
	return err == nil, nil
}
