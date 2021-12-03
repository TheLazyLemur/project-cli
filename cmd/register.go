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
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/TheLazyLemur/project-cli/config"
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

		c, err := config.GetConfig()
		if err != nil {
			log.Fatal("error: " + err.Error())
		}

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

		if _, err := os.Stat(c.StoreDirectory + alias); errors.Is(err, os.ErrNotExist) {

			err = ioutil.WriteFile(c.StoreDirectory+alias, data, 0)
			if err != nil {
				log.Fatal(err)
			}

			err = os.Chmod(c.StoreDirectory+alias, 0700)

		}
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringVarP(&project, "projectd", "p", "", "Project directory")
	registerCmd.Flags().StringVarP(&alias, "alias", "a", "", "Alias name")
	registerCmd.Flags().StringVarP(&editor, "editor", "e", "", "Editor to open project with")
}
