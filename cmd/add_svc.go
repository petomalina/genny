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

	"github.com/spf13/cobra"
)

// svcCmd represents the svc command
var svcCmd = &cobra.Command{
	Use:   "svc <api>",
	Short: "Adds a cmd_new microservice [into a specified cluster]",
	Long: ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please specify api to be implemented for the service")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("svc called")
	},
}

func init() {
	addCmd.AddCommand(svcCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// svcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// svcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
