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
	Short: "prints authenticated user's info",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		userInfo := mal.GetUserInfo()

		fmt.Printf("\x1b[1;34mUsername: \t%s\n\x1b[0m", userInfo.Name)
		fmt.Printf("\x1b[1;34mProfile Picture: \t%s\n\x1b[0m", userInfo.Picture)
		fmt.Printf("\x1b[1;34mGender: \t%s\n\x1b[0m", userInfo.Gender)
		fmt.Printf("\x1b[1;34mLocation: \t%s\n\x1b[0m", userInfo.Location)
		fmt.Printf("\x1b[1;34mBirthday: \t%s\n\x1b[0m", userInfo.Birthday)
		fmt.Printf("\x1b[1;34mTime Zone: \t%s\n\x1b[0m", userInfo.TimeZone)
		fmt.Printf("\x1b[1;34mJoined At: \t%s\n\x1b[0m", userInfo.JoinedAt)
		fmt.Printf("\x1b[1;34mUser ID: \t%d\n\x1b[0m", userInfo.Id)

		if userInfo.IsSupporter {
		  fmt.Printf("\x1b[33mYou are a MyAnimeList Supporter.\n\x1b[0m")
		}
	},
}

func init() {
	rootCmd.AddCommand(userInfoCmd)
}
