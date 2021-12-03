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
	"fmt"
	"io/ioutil"
	"log"

	"github.com/TheLazyLemur/project-cli/config"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all created project aliases",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		c, err := config.GetConfig()
		if err != nil {
			log.Fatal("error: " + err.Error())
		}

		files, err := ioutil.ReadDir(c.StoreDirectory)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			n := file.Name()
			if !file.IsDir() && n != "config.json" {
				fmt.Println(file.Name())
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
