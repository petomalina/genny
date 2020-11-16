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
	"github.com/petomalina/genny/internal/cmd_api"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Creates a cmd_new gRPC API definition",
	Long: `API definitions must be versioned as following:
<apiname>/<version> or <apiname>/<apisubname>/<version>

For an example:

monitoring/v1
monitoring/v2
monitoring/dashboard/v1

A base file with a gRPC service definition will be added
to the generated proto, e.g:

monitoring/v1/monitoring.proto (that will include the gRPC service def.)`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please specify the name of the API")
		}

		if !strings.Contains(args[0], "/") {
			return errors.New("your api is violating the <name>/<version> format")
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		apiPath := args[0]
		fmt.Println("Adding a new API:", apiPath)

		// this strips the version from the API, e.g.
		// documents/v1 becomes only documents,
		// documents/sheets/v1 becomes only sheets
		apiName := filepath.Base(filepath.Dir(apiPath))

		apiVersion := filepath.Base(apiPath)

		err := os.MkdirAll(filepath.Join("apis/proto", apiPath), os.ModePerm)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(filepath.Join("apis/proto", apiPath, apiName+".proto"),
			[]byte(cmd_api.ServiceDefinition(apiName, apiVersion)),
			os.ModePerm,
		)
		if err != nil {
			return err
		}

		viper.Set("apis", append(viper.GetStringSlice("apis"), apiPath))
		return viper.WriteConfig()
	},
}

func init() {
	addCmd.AddCommand(apiCmd)
}
