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
	failedIcon  = "\xE2\x9D\x97"
	warningIcon = "\xE2\x9D\x8C"
)

var imagesCmd = &cobra.Command{
	Use:   "images",
	Short: "Check which images in your cluster support arm64.",
	Long:  `Check which images in your cluster support arm64.`,
	Run:   imagesCmdRun,
}

func imagesCmdRun(cmd *cobra.Command, args []string) {
	k8sClient, err := k8s.NewKubernetesClient()
	if err != nil {
		log.Fatal(err)
	}
	imageList, err := k8sClient.GetAllImages()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Legends:\n%s - Supports arm64, %s - Do not support arm64, %s - Some error occurred\n", successIcon, warningIcon, failedIcon)
	fmt.Print("------------------------------------------------------------------------\n\n")
	for _, image := range imageList {
		var icon string
		supportsArm, err := images.CheckLinuxArm64Support(images.ToFullUrl(image))
		if err != nil {
			icon = warningIcon
		} else if supportsArm {
			icon = successIcon
		} else {
			icon = failedIcon
		}
		fmt.Printf("%s %s\n", image, icon)
	}
}

func init() {
	rootCmd.AddCommand(imagesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// imagesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// imagesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
