/*
Copyright © 2022 Harry Law <hrryslw@pm.me>
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice,
   this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its contributors
   may be used to endorse or promote products derived from this software
   without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
POSSIBILITY OF SUCH DAMAGE.
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
