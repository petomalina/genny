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
	"github.com/petomalina/genny/internal/cmd_new"
	"github.com/petomalina/genny/internal/perform"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
)

// newCmd represents the cmd_new command
var newCmd = &cobra.Command{
	Use:   "new <name>",
	Short: "Creates a new genny application scaffold without services",
	Long:  `A new application scaffold without services will be created`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please specify a name of the project")
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		gomodule := args[0]
		projectName := filepath.Base(gomodule)
		conf.Project = projectName

		folders := []string{
			fmt.Sprintf("%s/infrastructure", projectName),
			fmt.Sprintf("%s/services", projectName),
			fmt.Sprintf("%s/apis/proto", projectName),
			fmt.Sprintf("%s/apis/go-sdk", projectName),
			fmt.Sprintf("%s/apis/3rdparty", projectName),
		}

		fmt.Println("Creating project structure ...")
		// create all boilerplate folders needed for other commands
		for _, folder := range folders {
			if err := os.MkdirAll(folder, os.ModePerm); err != nil {
				return err
			}
		}

		files := map[string]string{
			fmt.Sprintf("%s/README.md", projectName):     fmt.Sprintf(`# %s`, projectName),
			fmt.Sprintf("%s/.gitignore", projectName):    cmd_new.DefaultGitignoreContents(),
			fmt.Sprintf("%s/apis/Makefile", projectName): cmd_new.DefaultMakefileContents(projectName),
		}

		fmt.Println("Creating project files ...")
		// create all boilerplate files
		for name, content := range files {
			if err := ioutil.WriteFile(name, []byte(content), os.ModePerm); err != nil {
				return err
			}
		}

		err := perform.Command("go", []string{"mod", "init", gomodule}, perform.Dir(projectName))
		if err != nil {
			return err
		}

		// initialize an empty git repo in the directory
		err = perform.Command("git", []string{"init", projectName})
		if err != nil {
			return err
		}

		fmt.Println(fmt.Sprintf("A new project '%s' was created, run 'cd %s' before you continue", projectName, projectName))

		return writeConfig(projectName)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
