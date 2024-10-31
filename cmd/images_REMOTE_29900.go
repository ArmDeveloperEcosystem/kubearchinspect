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
	"os"
	"sort"
	"strings"

	"github.com/ArmDeveloperEcosystem/kubearchinspect/internal/images"
	"github.com/ArmDeveloperEcosystem/kubearchinspect/internal/k8s"
	"github.com/spf13/cobra"
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

func containsAnyOf(input string, suggestions []string) bool {
	for _, suggestion := range suggestions {
		if strings.Contains(input, suggestion) {
			return true
		}
	}
	return false
}

func getFriendlyErrorMessage(err error) string {
	if err == nil {
		return ""
	}

	errorMessage := err.Error()
	switch {
	case containsAnyOf(errorMessage, []string{"authentication", "auth", "authorized"}):
		return "|| Authentication required. Please check your Docker credentials and your permissions."
	case containsAnyOf(errorMessage, []string{"no image found"}):
		return "|| Image not found."
	default:
		return "|| An unknown error occurred. Please run with debug -d for more details."
	}
}

func imagesCmdRun(cmd *cobra.Command, args []string) {
	debug, err := cmd.Flags().GetBool("debug")
	saveLog := cmd.Flags().Changed("logfile")

	if saveLog {

		path, err := cmd.Flags().GetString("logfile")
		file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

		if err != nil {
			log.Fatal(err)
		}

		log.SetOutput(file)
	}

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
	fmt.Printf("Legend:\n-------\n%s - arm64 supported\n%s - arm64 supported (with update)\n%s - arm64 not supported\n%s - error occurred\n", successIcon, upgradeIcon, failedIcon, errorIcon)
	fmt.Print("------------------------------------------------------------------------------------------------\n\n")

	sort.Strings(imageList)
	for _, image := range imageList {
		var icon string
		supportsArm, err := images.CheckLinuxArm64Support(image)
		if err != nil {
			icon = errorIcon
			if debug {
				fmt.Printf("error: %s\n", err)
			}
			if saveLog {
				log.Println(icon, " image: ", image, "||", "error: ", err)
			}
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

		if debug {
			fmt.Printf("%s %s\n", icon, image)
		} else {
			fmt.Printf("%s %s %s\n", icon, image, getFriendlyErrorMessage(err))
		}
	}
}

func init() {
	rootCmd.AddCommand(imagesCmd)

	imagesCmd.Flags().BoolP("debug", "d", false, "Enable debug mode")
	imagesCmd.Flags().StringP("logfile", "l", "", "log errors")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// imagesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// imagesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
