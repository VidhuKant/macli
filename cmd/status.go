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
	Long: "Set an anime's status\n" +
	"\n" +
    "Example Usage:\n" +
	" - \x1b[33m`macli status <anime-name>`\x1b[0m For interactive prompt (anime-name can be omitted)\n" +
	" - \x1b[33m`macli status -s \x1b[34mwatching|plan_to_watch|dropped|on_hold|completed\x1b[33m <anime-name>`\x1b[0m to specify status from command\n",
	Run: func(cmd *cobra.Command, args []string) {
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
		resp := mal.SetAnimeStatus(anime.Id, statusInput)
		fmt.Println(ui.CreateStatusUpdateConfirmationMessage(anime.Title, resp.Status))
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
		resp := mal.SetAnimeStatus(manga.Id, statusInput)
		fmt.Println(ui.CreateStatusUpdateConfirmationMessage(manga.Title, resp.Status))
	}
}

func init() {
	rootCmd.AddCommand(statusCmd)
    statusCmd.Flags().StringP("status", "s", "", "status to be set")
}
