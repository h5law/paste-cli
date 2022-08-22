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
	"os"

	"github.com/h5law/paste-cli/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents the get command
var (
	getUuid    string
	getVerbose bool

	getCmd = &cobra.Command{
		Use:   "get",
		Short: "Retrieve a paste",
		Long:  `Retrieve a paste from a paste-server instance with the given UUID`,
		Run: func(cmd *cobra.Command, args []string) {
			// Get response and load into struct
			resp, err := api.GetPaste()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// If verbose set create map for other values
			verbose := viper.GetBool("get-verbose")
			m := make(map[string]string)
			if verbose && resp.FileType != "" {
				m["filetype"] = resp.FileType
			}
			if verbose && resp.ExpiresAt != "" {
				m["expiresAt"] = resp.ExpiresAt
			}

			// Print map if verbose
			if verbose {
				fmt.Printf("uuid:      %s\n", viper.GetString("get-uuid"))
				fmt.Printf("filetype:  %s\n", m["filetype"])
				fmt.Printf("expiresAt: %s\n", m["expiresAt"])
				fmt.Println()
			}

			// Print content slice
			for _, v := range resp.Content {
				fmt.Println(v)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringVarP(
		&getUuid,
		"uuid",
		"u",
		"",
		"UUID of paste to fetch",
	)
	getCmd.MarkFlagRequired("uuid")

	getCmd.Flags().BoolVarP(
		&getVerbose,
		"verbose",
		"v",
		false,
		"Print detailed output",
	)

	viper.BindPFlag("get-uuid", getCmd.Flags().Lookup("uuid"))
	viper.BindPFlag("get-verbose", getCmd.Flags().Lookup("verbose"))
	viper.SetDefault("get-uuid", "")
	viper.SetDefault("get-verbose", false)
}
