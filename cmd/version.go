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

	"github.com/spf13/cobra"
)

const ascii = `
  _  __     _                            _     _____                           _   
 | |/ /    | |            /\            | |   |_   _|                         | |  
 | ' /_   _| |__   ___   /  \   _ __ ___| |__   | |  _ __  ___ _ __   ___  ___| |_ 
 |  <| | | | '_ \ / _ \ / /\ \ | '__/ __| '_ \  | | | '_ \/ __| '_ \ / _ \/ __| __|
 | . \ |_| | |_) |  __// ____ \| | | (__| | | |_| |_| | | \__ \ |_) |  __/ (__| |_ 
 |_|\_\__,_|_.__/ \___/_/    \_\_|  \___|_| |_|_____|_| |_|___/ .__/ \___|\___|\__|
                                                              | |                  
                                                              |_|                  `

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Check the version of KubeArchInspect you are running.",
		Long:  `Check the version of KubeArchInspect you are running.`,
		Run:   versionCmdRun,
	}
)

func versionCmdRun(_ *cobra.Command, _ []string) {
	fmt.Println(ascii)
	fmt.Printf("Version: %s\nCommit: %s\nBuilt at: %s\n", version, commit, date)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
