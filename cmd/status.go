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
	"strings"
	"github.com/MikunoNaka/macli/ui"
	"github.com/MikunoNaka/macli/mal"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Set an anime/manga's status",
	Long: `
-- help/description to be added later
`,
	Run: func(cmd *cobra.Command, args []string) {
		mal.Init() // needs to be manually called else it won't let you login
		searchInput := strings.Join(args, " ")

		statusInput, err := cmd.Flags().GetString("status")
		if err != nil {
			fmt.Println("Error while reading status flag.", err.Error())
		}

		mangaMode, err := cmd.Flags().GetBool("manga")
		if err != nil {
			fmt.Println("Error while reading manga flag.", err.Error())
		}

		if mangaMode {
			setMangaStatus(statusInput, searchInput)
		} else {
			setAnimeStatus(statusInput, searchInput)
		}

	},
}

func setAnimeStatus(statusInput, searchInput string) {
	if searchInput == "" {
		searchInput = ui.TextInput("Search Anime To Update: ", "Search can't be blank.")
	}

	anime := ui.AnimeSearch("Select Anime:", searchInput)

	if statusInput == "" {
		ui.AnimeStatusMenu(anime)
	} else {
		mal.SetAnimeStatus(anime.Id, statusInput)
		fmt.Printf("Successfully set \"%s\" to \"%s\"\n", anime.Title, statusInput)
	}
}

func setMangaStatus(statusInput, searchInput string) {
	if searchInput == "" {
		searchInput = ui.TextInput("Search Manga To Update: ", "Search can't be blank.")
	}

	manga := ui.MangaSearch("Select Manga:", searchInput)

	if statusInput == "" {
		ui.MangaStatusMenu(manga)
	} else {
		mal.SetAnimeStatus(manga.Id, statusInput)
		fmt.Printf("Successfully set \"%s\" to \"%s\"\n", manga.Title, statusInput)
	}
}

func init() {
	rootCmd.AddCommand(statusCmd)
    statusCmd.Flags().StringP("status", "s", "", "status to be set")
}
