/*
Copyright © 2021 Dan Rousseau danrousseau@protonmai.com

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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/TheLazyLemur/project-cli/data"
	"github.com/spf13/cobra"
)

var alias string
var project string
var editor string

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new project",
	Long:  `Register a new project providing it the project alias name, project directory and preferred editor`,
	Run: func(cmd *cobra.Command, args []string) {

		newEntry := &data.Entry{
			Alias:     alias,
			Editor:    editor,
			Directory: project,
		}

		b, err := json.Marshal(newEntry)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(b))

		val := string(b)
		data := []byte(val)

		err = ioutil.WriteFile("/home/dan/.config/project-cli/"+alias, data, 0)

		if err != nil {
			log.Fatal(err)
		}

		err = os.Chmod("/home/dan/.config/project-cli/"+alias, 0700)

	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringVarP(&project, "projectd", "p", "", "Project directory")
	registerCmd.Flags().StringVarP(&alias, "alias", "a", "", "Alias name")
	registerCmd.Flags().StringVarP(&editor, "editor", "e", "", "Editor to open project with")
}
