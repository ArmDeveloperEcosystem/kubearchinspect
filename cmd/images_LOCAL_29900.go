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

package cmd

import (
	"fmt"
	"log"
	"sort"

	"github.com/spf13/cobra"

	"github.com/ArmDeveloperEcosystem/kubearchinspect/internal/images"
	"github.com/ArmDeveloperEcosystem/kubearchinspect/internal/k8s"
)

const (
	successIcon = "\xE2\x9C\x85"
	errorIcon   = "\xF0\x9F\x9A\xAB"
	failedIcon  = "\xE2\x9D\x8C"
	upgradeIcon = "\xF0\x9F\x86\x99"
)

var imagesCmd = &cobra.Command{
	Use:   "images",
	Short: "Check which images in your cluster support arm64.",
	Long:  `Check which images in your cluster support arm64.`,
	Run:   imagesCmdRun,
}

func imagesCmdRun(_ *cobra.Command, _ []string) {

	k8sClient, err := k8s.NewKubernetesClient()
	if err != nil {
		log.Fatal(err)
	}

	imageList, err := k8sClient.GetAllImages()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(
		"Legend:\n-------\n%s - arm64 supported\n%s - arm64 supported (with update)\n%s - arm64 not supported\n%s - error occurred\n%s",
		successIcon,
		upgradeIcon,
		failedIcon,
		errorIcon,
		"------------------------------------------------------------------------------------------------\n\n",
	)

	sort.Strings(imageList)
	for _, image := range imageList {
		var (
			icon             string
			supportsArm, err = images.CheckLinuxArm64Support(image)
		)

		switch {
		case err != nil:
			if debugEnabled {
				fmt.Printf("error: %s\n", err)
			}
			icon = errorIcon
		case supportsArm:
			icon = successIcon
		default:
			latestSupportsArm, _ := images.CheckLatestLinuxArm64Support(image)
			if latestSupportsArm {
				icon = upgradeIcon
			} else {
				icon = failedIcon
			}
		}

		fmt.Printf("%s %s\n", icon, image)
	}
}

func init() {
	rootCmd.AddCommand(imagesCmd)
}
