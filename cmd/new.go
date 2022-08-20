/*
Copyright © 2022 Harry Law <hrryslw@pm.me>

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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// newCmd represents the new command
var (
	filePath  string
	fileType  string
	expiresIn int

	newCmd = &cobra.Command{
		Use:   "new",
		Short: "Create a new paste",
		Long: `Create a new paste from os.Stdin or using the file flag and send
the content to the paste-server.

Running this command will return the UUID, expiration date and
access key for the paste created.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("new called")
		},
	}
)

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().StringVarP(
		&filePath,
		"file",
		"f",
		"",
		"Path to file for upload",
	)
	newCmd.Flags().StringVarP(
		&fileType,
		"filetype",
		"t",
		"plaintext",
		"Filetype of paste",
	)
	newCmd.Flags().IntVarP(
		&expiresIn,
		"expires",
		"e",
		14,
		"Number of days before paste expries (1-30)",
	)

	viper.BindPFlag("file", newCmd.Flags().Lookup("file"))
	viper.BindPFlag("filetype", newCmd.Flags().Lookup("filetype"))
	viper.BindPFlag("expiresIn", newCmd.Flags().Lookup("expires"))
	viper.SetDefault("file", "")
	viper.SetDefault("filetype", "plaintext")
	viper.SetDefault("expiresIn", 14)
}
