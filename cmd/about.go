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
)

var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "More information about macli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("macli is an unofficial MyAnimeClient for use inside the terminal.\n" +
			"\n" +
			"\x1b[34mmacli  Copyright (C) 2022  Vidhu Kant Sharma <vidhukant@vidhukant.xyz>\n" +
			"This program comes with ABSOLUTELY NO WARRANTY;\n" +
			"This is free software, and you are welcome to redistribute it\n" +
			"under certain conditions; For details refer to the GNU General Public License.\n" +
			"You should have received a copy of the GNU General Public License\n" +
			"along with this program.  If not, see <https://www.gnu.org/licenses/>.\x1b[0m\n" +
			"\n" +
			"\x1b[35mPlease report any bugs on the GitHub page: https://github.com/MikunoNaka/macli\n" +
			"or through email: vidhukant@vidhukant.xyz\x1b[0m\n\n",
		)
		fmt.Println("Visit \x1b[36mhttps://www.vidhukant.xyz/docs/macli/\x1b[0m for more info.")
		fmt.Println("Run \x1b[33m`macli --help`\x1b[0m for instructions.")
	},
}

func init() {
	rootCmd.AddCommand(aboutCmd)
}
