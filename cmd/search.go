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
	"strings"
	"github.com/spf13/cobra"
	"github.com/MikunoNaka/macli/ui"
)

var searchCmd = &cobra.Command {
	Use:   "search",
	Short: "Search for an anime/manga",
	Long: `
-- help/description to be added later
`,
	Run: func(cmd *cobra.Command, args []string) {
		// read searchInput from command
		searchInput := strings.Join(args, " ")
		mangaMode, _ := cmd.Flags().GetBool("manga")

		if mangaMode {
			searchManga(searchInput)
		} else {
			searchAnime(searchInput)
		}
	},
}

func searchManga(searchInput string) {
	if searchInput == "" {
		searchInput = ui.TextInput("Search Manga: ", "Search can't be blank.")
	}
    manga := ui.MangaSearch("Select Manga:", searchInput)
	mangaIsAdded := manga.MyListStatus.Status != ""
	ui.MangaActionMenu(mangaIsAdded)(manga)
}

func searchAnime(searchInput string) {
	if searchInput == "" {
		searchInput = ui.TextInput("Search Anime: ", "Search can't be blank.")
	}
	anime := ui.AnimeSearch("Select Anime:", searchInput)
	animeIsAdded := anime.MyListStatus.Status != ""
	ui.AnimeActionMenu(animeIsAdded)(anime)
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
