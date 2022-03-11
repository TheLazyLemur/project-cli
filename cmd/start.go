/*
Copyright © 2021 Dan Rousseau <danrousseau@protonmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/TheLazyLemur/project-cli/config"
	"github.com/TheLazyLemur/project-cli/data"
	commands "github.com/TheLazyLemur/project-cli/utils"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Open project",
	Long:  `Open a previously registered project alias`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := config.GetConfig()
		if err != nil {
			log.Fatal("error: " + err.Error())
		}

		text, err := ioutil.ReadFile(c.StoreDirectory + alias)

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
