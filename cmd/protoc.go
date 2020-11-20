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
	"path/filepath"
)

// protocCmd represents the protoc command
var protocCmd = &cobra.Command{
	Use:   "protoc",
	Short: "A brief description of your command",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		protomodules := ""
		for _, m := range conf.ProtoModules {
			protomodules += " -I" + filepath.Join(m.Path, m.IncludePath)
		}

		for _, svc := range conf.APIs {
			err := perform.Command(
				"protoc",
				[]string{
					"-Iproto -I." + protomodules,
					fmt.Sprintf("proto/%s.proto", svc),
				},
				perform.Dry(), perform.Logger(logger),
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
}
