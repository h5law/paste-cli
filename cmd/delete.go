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

// deleteCmd represents the delete command
var (
	delUuid      string
	delAccessKey string

	deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete a paste",
		Long: `Delete a paste with the given UUID from a paste-server instance provided the
access key provided matches.`,
		Run: func(cmd *cobra.Command, args []string) {
			resp, err := api.DeletePaste()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			if resp != "" {
				fmt.Println(resp)
			} else {
				fmt.Println("Paste deleted")
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringVarP(
		&delUuid,
		"uuid",
		"u",
		"",
		"UUID of paste to delete",
	)
	deleteCmd.MarkFlagRequired("uuid")
	deleteCmd.Flags().StringVarP(
		&delAccessKey,
		"access-key",
		"a",
		"",
		"Access key needed to delete paste",
	)
	deleteCmd.MarkFlagRequired("access-key")

	viper.BindPFlag("del-accessKey", deleteCmd.Flags().Lookup("access-key"))
	viper.BindPFlag("del-uuid", deleteCmd.Flags().Lookup("uuid"))
	viper.SetDefault("del-accessKey", "")
	viper.SetDefault("del-uuid", "")
}
