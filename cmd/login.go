/*
macli - Unofficial CLI-Based MyAnimeList Client
Copyright © 2022 Vidhu Kant Sharma <vidhukant@vidhukant.xyz>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/MikunoNaka/macli/auth"
)

var loginCmd = &cobra.Command {
	Use:   "login",
	Short: "Login with your MyAnimeList account",
	Long: "To authenticate with macli, a Client ID is required.\n" +
	"If you have logged in before and ran `macli logout`, you may not need to enter your Client ID again unless you specifically deleted it.\n" +
	"\n" +
	"\x1b[31;1mHow to generate a Client ID:\x1b[0m\n" +
	" - Go to \x1b[36mhttps://myanimelist.net/apiconfig\x1b[0m\n" +
	" - Click on \x1b[33m\"Create ID\"\x1b[0m\n" +
	" - Inside the form you can set all the details to whatever you'd like\n" +
	" - For macli to work properly, you only need to set \x1b[33m\"App Redirect Url\"\x1b[0m to \x1b[36mhttp://localhost:8000\x1b[0m\n" +
	" - After that, hit submit, then copy your Client ID, run `macli login` and paste in your Client ID.\n" +
	" - \x1b[31mIf after running `macli login` it opens a dialogue box in the browser asking for credentials,\n   and not the MyAnimeList login page, that means you have entered your Client ID wrong.\x1b[0m\n" +
	"",
	Run: func(cmd *cobra.Command, args []string) {
		auth.Login()
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	// TODO: save given token to keyring
    // rootCmd.Flags().StringVarP(&mal.Secret, "authentication-token", "t", "", "MyAnimeList authentication token to use (overrides system keyring if any)")
}
