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
	"fmt"
	"github.com/spf13/cobra"
	"github.com/MikunoNaka/macli/mal"
)

var userInfoCmd = &cobra.Command {
	Use:   "user-info",
	Short: "Shows logged-in user's info",
	Long:
`Print logged-in user's profile information.
Currently, MyAnimeList doesn't allow reading of other users' profiles.
`,
	Run: func(cmd *cobra.Command, args []string) {
    mal.Init()
		userInfo := mal.GetUserInfo()

		fmt.Printf("\x1b[1;34mUsername: \x1b[0m%s\n", userInfo.Name)
		fmt.Printf("\x1b[1;34mProfile Picture: \x1b[0m%s\n\x1b[0m", userInfo.Picture)
		fmt.Printf("\x1b[1;34mGender: \x1b[0m%s\n", userInfo.Gender)
		fmt.Printf("\x1b[1;34mLocation: \x1b[0m%s\n", userInfo.Location)
		fmt.Printf("\x1b[1;34mBirthday: \x1b[0m%s\n", userInfo.Birthday)
		fmt.Printf("\x1b[1;34mTime Zone: \x1b[0m%s\n", userInfo.TimeZone)
		fmt.Printf("\x1b[1;34mJoined At: \x1b[0m%s\n", userInfo.JoinedAt)
		fmt.Printf("\x1b[1;34mUser ID: \x1b[0m%d\n", userInfo.Id)

		if userInfo.IsSupporter {
		  fmt.Printf("\x1b[33mYou are a MyAnimeList Supporter.\n\x1b[0m")
		}
	},
}

func init() {
	rootCmd.AddCommand(userInfoCmd)
}
