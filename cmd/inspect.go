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

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/ArmDeveloperEcosystem/kubearchinspect/internal/images"
)

var inspectCmd = &cobra.Command{
	Use:   "inspect <image>",
	Short: "Check whether a specific image supports arm64",
	Long:  `Check whether a specific container image supports the arm64 architecture`,
	Args:  cobra.ExactArgs(1),
	Run:   inspectCmdRun,
}

func inspectCmdRun(_ *cobra.Command, args []string) {

	image := args[0]

	fmt.Printf(
		legend,
		successIcon,
		upgradeIcon,
		failedIcon,
		errorIcon,
		"------------------------------------------------------------------------------------------------\n\n",
	)

	var (
		icon             string
		supportsArm, err = images.CheckLinuxArm64Support(image)
	)

	switch {
	case err != nil:
		icon = errorIcon
	case supportsArm:
		icon = successIcon
	default:
		latestImage := images.GetLatestImage(image)
		latestSupportsArm, _ := images.CheckLinuxArm64Support(latestImage)
		if latestSupportsArm {
			icon = upgradeIcon
		} else {
			icon = failedIcon
		}
	}

	fmt.Printf("%s %s %s\n", icon, image, images.GetFriendlyErrorMessage(err))
	if err != nil && debugEnabled {
		fmt.Printf("Error: %s\n", err)
	}
}

func init() {
	rootCmd.AddCommand(inspectCmd)
}
