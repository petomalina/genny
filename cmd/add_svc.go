/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// svcCmd represents the svc command
var svcCmd = &cobra.Command{
	Use:   "svc <name> [<api>/<version>...]",
	Short: "Adds a cmd_new microservice [into a specified cluster]",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please specify name of the service")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		apiNames := args[1:]

		fmt.Println("svc called:" + name + ";" + strings.Join(apiNames, " "))
	},
}

func init() {
	addCmd.AddCommand(svcCmd)

	svcCmd.Flags().Bool("http", false, "Allow HTTP JSON transcoding gateway")

	svcCmd.Flags().Bool("pubsub", false, "Allow Pub/Sub integration in xrpc")

	svcCmd.Flags().Bool("webrpc", false, "Allow WebRPC transcoder")

	svcCmd.Flags().Bool("cloude-events", false, "Allow Cloud Events integration")
}
