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

	"github.com/Arm-Debug/armer/internal/images"
	"github.com/Arm-Debug/armer/internal/k8s"
	"github.com/spf13/cobra"
)

const (
	successIcon = "\xE2\x9C\x85"
	warningIcon = "\xE2\x9D\x97"
	failedIcon  = "\xE2\x9D\x8C"
	upgradeIcon = "\xE2\xAC\x86"
)

var imagesCmd = &cobra.Command{
	Use:   "images",
	Short: "Check which images in your cluster support arm64.",
	Long:  `Check which images in your cluster support arm64.`,
	Run:   imagesCmdRun,
}

func imagesCmdRun(cmd *cobra.Command, args []string) {
	debug, err := cmd.Flags().GetBool("debug")
	if err != nil {
		log.Fatal(err)
	}

	k8sClient, err := k8s.NewKubernetesClient()
	if err != nil {
		log.Fatal(err)
	}
	imageList, err := k8sClient.GetAllImages()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Legends:\n%s - Supports arm64, %s - Does not support arm64, %s - Upgrade for arm64 support, %s - Some error occurred\n", successIcon, failedIcon, upgradeIcon, warningIcon)
	fmt.Print("------------------------------------------------------------------------------------------------\n\n")
	for _, image := range imageList {
		var icon string
		supportsArm, err := images.CheckLinuxArm64Support(image)
		if err != nil {
			if debug {
				fmt.Printf("error: %s\n", err)
			}
			icon = warningIcon
		} else if supportsArm {
			icon = successIcon
		} else {
			latestSupportsArm, _ := images.CheckLatestLinuxArm64Support(image)
			if latestSupportsArm {
				icon = upgradeIcon
			} else {
				icon = failedIcon
			}
		}
		fmt.Printf("%s %s\n", image, icon)
	}
}

func init() {
	rootCmd.AddCommand(imagesCmd)

	imagesCmd.Flags().BoolP("debug", "d", false, "Enable debug mode")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// imagesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// imagesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
