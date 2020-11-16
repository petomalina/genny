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
	"fmt"
	"github.com/petomalina/genny/internal/perform"
	"github.com/spf13/cobra"
)

// protocCmd represents the protoc command
var protocCmd = &cobra.Command{
	Use:   "protoc",
	Short: "A brief description of your command",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		protomodules := ""
		for _, m := range conf.ProtoModules {
			protomodules += " -I" + m
		}

		for _, svc := range conf.APIs {
			err := perform.Command(
				"protoc",
				[]string{
					"-Iproto -I." + protomodules,
					fmt.Sprintf("proto/%s.proto", svc),
				},
				perform.Dry(),
			)
			if err != nil {
				return err
			}
		}

		return nil
	},
}

func init() {
	runCmd.AddCommand(protocCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// protocCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// protocCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
