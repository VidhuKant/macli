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
	"github.com/MikunoNaka/macli/ui"
	"github.com/MikunoNaka/macli/auth"
)

var loginCmd = &cobra.Command {
	Use:   "login",
	Short: "Login with your MyAnimeList client secret",
	Long: `
Currently, macli doesn't support logging in.
You need to manually generate an access token/client secret to authorise
macli with your MyAnimeList account.

An easy way to generate a token is to use my python script:
https://github.com/MikunoNaka/mal-authtoken-generator
`,
	Run: func(cmd *cobra.Command, args []string) {
		secret := ui.PasswordInput("Enter your client secret: ", "Client secret can't be empty")
		auth.Login(secret)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
