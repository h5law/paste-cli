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
	"github.com/h5law/paste-cli/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// updateCmd represents the update command
var (
	updFilePath  string
	updFileType  string
	updUuid      string
	updAccessKey string
	updExpiresIn int

	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update paste with the UUID provided",
		Long: `Update a paste with the matching UUID automatically extending its time to
expire by 14 days unless told otherwise.`,
		Run: func(cmd *cobra.Command, args []string) {
			if pipe := utils.IsInputFromPipe(); pipe {
				viper.Set("file", "")
			}

			resp, err := api.UpdatePaste()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Printf("uuid:      \t%s\n", resp["uuid"])
			fmt.Printf("expiresAt: \t%s\n", resp["expiresAt"])
			fmt.Printf("url:       \t%s\n", resp["url"])
		},
	}
)

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringVarP(
		&updUuid,
		"uuid",
		"u",
		"",
		"UUID of paste to edit",
	)
	updateCmd.MarkFlagRequired("uuid")
	updateCmd.Flags().StringVarP(
		&updAccessKey,
		"access-key",
		"a",
		"",
		"Access key needed to update paste",
	)
	updateCmd.MarkFlagRequired("access-key")

	updateCmd.Flags().StringVarP(
		&updFilePath,
		"file",
		"f",
		"",
		"Path to file to update with",
	)

	updateCmd.Flags().StringVarP(
		&updFileType,
		"filetype",
		"t",
		"",
		"Filetype of paste",
	)
	updateCmd.Flags().IntVarP(
		&updExpiresIn,
		"expires",
		"e",
		0,
		"Number of days before paste expries (1-30)",
	)

	viper.BindPFlag("upd-file", updateCmd.Flags().Lookup("file"))
	viper.BindPFlag("upd-filetype", updateCmd.Flags().Lookup("filetype"))
	viper.BindPFlag("upd-expiresIn", updateCmd.Flags().Lookup("expires"))
	viper.BindPFlag("upd-accessKey", updateCmd.Flags().Lookup("access-key"))
	viper.BindPFlag("upd-uuid", updateCmd.Flags().Lookup("uuid"))
	viper.SetDefault("upd-file", "")
	viper.SetDefault("upd-filetype", "")
	viper.SetDefault("upd-expiresIn", 0)
	viper.SetDefault("upd-accessKey", "")
	viper.SetDefault("upd-uuid", "")
}
