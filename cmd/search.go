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
	"os"
	"strings"

	"github.com/MikunoNaka/macli/mal"
	"github.com/MikunoNaka/macli/ui"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command {
	Use:   "search",
	Short: "Search for an anime/manga",
	Long: "Search for an anime or a manga on MyAnimeList\n" +
	"\n" +
    "Example Usage:\n" +
	"\t\x1b[33m`macli search <anime-name>`\x1b[0m searches for an anime\n" +
	"\t\x1b[33m`macli search -m <manga-name>`\x1b[0m searches for a manga\n" +
	"\t\x1b[33m`macli search`\x1b[0m interactively asks for an anime to search for (same for manga with -m/--manga flag)\n",
	Run: func(cmd *cobra.Command, args []string) {
		mal.Init()
		// read searchInput from command
		searchInput := strings.Join(args, " ")
		mangaMode, err := cmd.Flags().GetBool("manga")
		if err != nil {
			fmt.Println("Error while reading \x1b[33m--manga\x1b[0m flag.", err.Error())
			os.Exit(1)
		}

		if mangaMode {
			searchManga(searchInput)
		} else {
			searchAnime(searchInput)
		}
	},
}

func searchManga(searchInput string) {
	if entryId > 0 {
	    selectedManga := mal.GetMangaData(entryId, []string{"my_list_status", "title"})
		fmt.Println("Selected: \x1b[35m" + selectedManga.Title + "\x1b[0m")
		ui.MangaActionMenu(selectedManga.MyListStatus.Status != "")(selectedManga)
		os.Exit(0)
	}
	if searchInput == "" {
		searchInput = ui.TextInput("Search Manga: ", "Search can't be blank.")
	}
    manga := ui.MangaSearch("Select Manga:", searchInput)
	selectedManga := mal.GetMangaData(manga.Id, []string{"my_list_status", "num_chapters"})
	mangaIsAdded := selectedManga.MyListStatus.Status != ""
	ui.MangaActionMenu(mangaIsAdded)(selectedManga)
}

func searchAnime(searchInput string) {
	if entryId > 0 {
	    selectedAnime := mal.GetAnimeData(entryId, []string{"my_list_status"})
		fmt.Println("Selected: \x1b[35m" + selectedAnime.Title + "\x1b[0m")
		ui.AnimeActionMenu(selectedAnime.MyListStatus.Status != "")(selectedAnime)
		os.Exit(0)
	}
	if searchInput == "" {
		searchInput = ui.TextInput("Search Anime: ", "Search can't be blank.")
	}
	anime := ui.AnimeSearch("Select Anime:", searchInput)
	selectedAnime := mal.GetAnimeData(anime.Id, []string{"my_list_status", "num_episodes"})
	animeIsAdded := selectedAnime.MyListStatus.Status != ""
	ui.AnimeActionMenu(animeIsAdded)(selectedAnime)
}

func init() {
	rootCmd.AddCommand(searchCmd)
    searchCmd.Flags().IntVarP(&ui.PromptLength, "prompt-length", "l", 5, "Length of select prompt")
    searchCmd.Flags().BoolVarP(&mangaMode, "manga", "m", false, "Use manga mode")
    searchCmd.Flags().IntVarP(&mal.SearchLength, "search-length", "n", 10, "Amount of search results to load")
    searchCmd.Flags().IntVarP(&mal.SearchOffset, "search-offset", "o", 0, "Offset for the search results")
    searchCmd.Flags().BoolVarP(&mal.SearchNSFW, "search-nsfw", "", false, "Include NSFW-rated items in search results")
    searchCmd.Flags().BoolVarP(&queryOnlyMode, "query", "q", false, "Query only (don't update data)")
    searchCmd.Flags().IntVarP(&entryId, "id", "i", -1, "Manually specify the ID of anime/manga (overrides search)")
}
