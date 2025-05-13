/*
Copyright 2025 Arm Limited

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

var (
	kubeConfigPath string
	kubeContext    string

	imagesCmd = &cobra.Command{
		Use:   "images",
		Short: "Check which images in your cluster support arm64",
		Long:  `Check which images in your cluster support arm64`,
		Run:   imagesCmdRun,
	}
)

func imagesCmdRun(_ *cobra.Command, _ []string) {

	var loggingEnabled = len(loggingFile) > 0

	if loggingEnabled {
		file, err := os.OpenFile(loggingFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}

		log.SetOutput(file)
	}

	k8sClient, err := k8s.NewKubernetesClient(kubeConfigPath, kubeContext, debugEnabled)
	if err != nil {
		log.Fatal(err)
	}

	imageMap, err := k8sClient.GetAllImages()
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

	type result struct {
		image string
		icon  string
		err   error
	}

	results := make(chan result, len(imageMap))

	for image := range imageMap {
		go func(image string) {
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

			results <- result{image: image, icon: icon, err: err}
		}(image)
	}

	var resultList []result
	for range imageMap {
		resultList = append(resultList, <-results)
	}

	sort.Slice(resultList, func(i, j int) bool {
		return resultList[i].image < resultList[j].image
	})

	for _, res := range resultList {
		fmt.Printf("%s %s %s\n", res.icon, res.image, images.GetFriendlyErrorMessage(res.err))
		if res.err != nil && debugEnabled {
			fmt.Printf("Error: %s\n", res.err)
			fmt.Printf("Pods: %s\n", imageMap[res.image])
		}
		if res.err != nil && loggingEnabled {
			log.Println(res.icon, " image: ", res.image, " error: ", res.err, "\n", imageMap[res.image])
		}
	}

}

func init() {
	rootCmd.AddCommand(imagesCmd)

	// Add flags
	imagesCmd.PersistentFlags().StringVarP(&kubeConfigPath, "kube-config-path", "c", "", "Path to your Kube config file. (Default: '~/.kube/config')")
	imagesCmd.PersistentFlags().StringVarP(&kubeContext, "kube-context", "", "", "The Kubernetes context to be used. (Default: Current context in the config)")
}
