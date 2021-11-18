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

	"github.com/TheLazyLemur/project-cli/data"
	commands "github.com/TheLazyLemur/project-cli/utils"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Open project",
	Long:  `Open a previosuly specified project`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(alias)
		text, err := ioutil.ReadFile("/home/dan/.config/project-cli/" + alias)

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(text))

		var data data.Entry
		if err := json.Unmarshal(text, &data); err != nil {
			fmt.Println("failed to unmarshal:", err)
		} else {
			fmt.Println(data.Alias)
			commands.OpenEditor(data.Editor, data.Directory)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().StringVarP(&alias, "alias", "a", "", "Project to open")
}
