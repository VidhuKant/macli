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
	"os"
	"fmt"
	"github.com/MikunoNaka/macli/ui"
	"github.com/MikunoNaka/macli/util"
	"github.com/MikunoNaka/macli/mal"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print out user's animelist",
	Long: "To be added", // +
	// "\n" +
  //   "Example Usage:\n" +
	// " - \x1b[33m`macli chapters <anime-name>`\x1b[0m For interactive prompt (anime-name can be omitted)\n" +
	// " - \x1b[33m`macli chapters -s 4 <anime-name>`\x1b[0m to set the chapters to 4\n" +
	// " - \x1b[33m`macli chapters -s +1 <anime-name>`\x1b[0m to increment the chapters by 1\n" +
	// " - \x1b[33m`macli chapters -s -2 <anime-name>`\x1b[0m to decrement the chapters by 2\n",
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := util.BindListConfig(cmd.Flags())
		if err != nil {
			fmt.Println("Error while parsing flags.", err.Error())
			os.Exit(1)
		}
		mal.SearchLength = conf.ResultsLength
		mal.SearchOffset = conf.ResultsOffset
		mal.SearchNSFW = conf.IncludeNSFW
    mal.Init()

    status, err := cmd.Flags().GetString("status")
    if err != nil {
      fmt.Println("error while reading \x1b[33m--status\x1b[0m flag:", err)
      os.Exit(1)
    }

    user, err := cmd.Flags().GetString("user")
    if err != nil {
      fmt.Println("error while reading \x1b[33m--user\x1b[0m flag:", err)
      os.Exit(1)
    }

    sort, err := cmd.Flags().GetString("sort")
    if err != nil {
      fmt.Println("error while reading \x1b[33m--sort\x1b[0m flag:", err)
      os.Exit(1)
    }

	  if mangaMode {
		  ui.MangaList(mal.MangaList(user, status, sort, conf.IncludeNSFW))
	  } else {
		  ui.AnimeList(mal.AnimeList(user, status, sort, conf.IncludeNSFW))
	  }
	},
}

func init() {
  rootCmd.AddCommand(listCmd)
  listCmd.Flags().StringP("status", "", "", "Status (leave blank for all)")
  listCmd.Flags().StringP("user", "", "@me", "User (@me or blank for self)")
  listCmd.Flags().StringP("sort", "", "list_score", "Sort the list")
  listCmd.Flags().BoolVarP(&mangaMode, "manga", "m", false, "Use manga mode")
  listCmd.Flags().StringVarP(&mal.Secret, "authentication-token", "t", "", "MyAnimeList authentication token to use (overrides system keyring if any)")

  listCmd.Flags().IntP("results-length", "n", 10, "Amount of results to load")
  listCmd.Flags().IntP("results-offset", "o", 0, "Offset for the results")
  listCmd.Flags().BoolP("include-nsfw", "", false, "Include NSFW-rated items in search results")
}
