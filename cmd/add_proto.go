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
	"github.com/petomalina/genny/internal/perform"
	"github.com/petomalina/genny/internal/types"
	"github.com/spf13/cobra"
	"path/filepath"
	"strings"
)

// protoCmd represents the proto command
var protoCmd = &cobra.Command{
	Use:   "proto <path> [include-path]",
	Short: "Adds a protobuf dependency as a submodule",
	Long: `For an example:

# installs api-common-protos as is as third party protobuf
genny add proto github.com/googleapis/api-common-protos

# installs protoc-gen-validate from envoy as is as third party protobuf
genny add proto https://github.com/envoyproxy/protoc-gen-validate

# installs protobuf default library and targets its 'src' for includes.
# in case the 'src' is omitted, you will need to import them as src/google/protobuf
# instead of google/protobuf, so we strongly recommend it'
genny add proto github.com/protocolbuffers/protobuf src`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please provide a proto repository to include")
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		moduleName := args[0]
		// normalize the git name with https
		moduleNameGit := moduleName
		if !strings.HasPrefix(moduleNameGit, "https://") {
			moduleNameGit = "https://" + moduleNameGit
		}

		// path that should actually be included
		includePath := ""
		if len(args) > 1 {
			includePath = strings.Join(args[1:], " ")
		}

		err := perform.Command("git", []string{
			"submodule",
			"add",
			moduleNameGit,
			fmt.Sprintf("./apis/3rdparty/%s", filepath.Base(moduleName)),
		}, perform.Logger(logger))
		if err != nil {
			return err
		}

		conf.ProtoModules = append(
			conf.ProtoModules,
			types.ProtoModule{
				Repository: moduleNameGit,
				Path: filepath.Join(
					"3rdparty",
					filepath.Base(moduleName),
					includePath,
				),
				IncludePath: includePath,
			},
		)
		return writeConfig()
	},
}

func init() {
	addCmd.AddCommand(protoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// protoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// protoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
