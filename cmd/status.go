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
	"github.com/MikunoNaka/macli/ui"
	"github.com/MikunoNaka/macli/mal"
	a "github.com/MikunoNaka/MAL2Go/v4/anime"
	m "github.com/MikunoNaka/MAL2Go/v4/manga"

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
    mal.Init()
		searchInput := strings.Join(args, " ")

		statusInput, err := cmd.Flags().GetString("set-value")
		if err != nil {
			fmt.Println("Error while reading \x1b[33m--set-value\x1b[0m flag.", err.Error())
			os.Exit(1)
		}

		mangaMode, err := cmd.Flags().GetBool("manga")
		if err != nil {
			fmt.Println("Error while reading \x1b[33m--manga\x1b[0m flag.", err.Error())
			os.Exit(1)
		}

		if mangaMode {
			setMangaStatus(statusInput, searchInput)
		} else {
			setAnimeStatus(statusInput, searchInput)
		}

	},
}

func setAnimeStatus(statusInput, searchInput string) {
	var selectedAnime a.Anime
	if entryId > 0 {
	    selectedAnime = mal.GetAnimeData(entryId, []string{"my_list_status"})
	}

	if searchInput == "" && entryId < 1 {
	    var promptText string
		if queryOnlyMode {
			promptText = "Search Anime to Get Status of: "
		} else {
			promptText = "Search Anime to Set Status of: "
		}
	   	searchInput = ui.TextInput(promptText, "Search can't be blank.")
	}

	if entryId < 1 {
		anime := ui.AnimeSearch("Select Anime:", searchInput)
		selectedAnime = mal.GetAnimeData(anime.Id, []string{"my_list_status"})
	}

	if queryOnlyMode {
    	fmt.Println(ui.CreateStatusUpdateConfirmationMessage(selectedAnime.Title, selectedAnime.MyListStatus.Status, ""))
		os.Exit(0)
	}

    if statusInput == "" {
    	ui.AnimeStatusMenu(selectedAnime)
    } else {
    	resp := mal.SetAnimeStatus(selectedAnime.Id, statusInput)
    	fmt.Println(ui.CreateStatusUpdateConfirmationMessage(selectedAnime.Title, resp.Status, selectedAnime.MyListStatus.Status))
    }
}

func setMangaStatus(statusInput, searchInput string) {
	var selectedManga m.Manga
	if entryId > 0 {
	    selectedManga = mal.GetMangaData(entryId, []string{"my_list_status"})
	}

	if searchInput == "" && entryId < 1 {
	    var promptText string
		if queryOnlyMode {
			promptText = "Search Manga to Get Status of: "
		} else {
			promptText = "Search Manga to Set Status of: "
		}
	   	searchInput = ui.TextInput(promptText, "Search can't be blank.")
	}

	if entryId < 1 {
	  manga := ui.MangaSearch("Select Manga:", searchInput)
	  selectedManga = mal.GetMangaData(manga.Id, []string{"my_list_status"})
	}

	if queryOnlyMode {
    	fmt.Println(ui.CreateStatusUpdateConfirmationMessage(selectedManga.Title, selectedManga.MyListStatus.Status, ""))
		os.Exit(0)
	}

    if statusInput == "" {
    	ui.MangaStatusMenu(selectedManga)
    } else {
    	resp := mal.SetMangaStatus(selectedManga.Id, statusInput)
    	fmt.Println(ui.CreateStatusUpdateConfirmationMessage(selectedManga.Title, selectedManga.MyListStatus.Status, resp.Status))
    }
}

func init() {
	rootCmd.AddCommand(statusCmd)
    statusCmd.Flags().StringP("set-value", "s", "", "status to be set")
    statusCmd.Flags().IntVarP(&ui.PromptLength, "prompt-length", "l", 5, "Length of select prompt")
    statusCmd.Flags().IntVarP(&mal.SearchLength, "search-length", "n", 10, "Amount of search results to load")
    statusCmd.Flags().IntVarP(&mal.SearchOffset, "search-offset", "o", 0, "Offset for the search results")
    statusCmd.Flags().BoolVarP(&mal.SearchNSFW, "search-nsfw", "", false, "Include NSFW-rated items in search results")
    statusCmd.Flags().BoolVarP(&mangaMode, "manga", "m", false, "Use manga mode")
    statusCmd.Flags().BoolVarP(&queryOnlyMode, "query", "q", false, "Query only (don't update data)")
    statusCmd.Flags().IntVarP(&entryId, "id", "i", -1, "Manually specify the ID of anime/manga (overrides search)")
    statusCmd.Flags().StringVarP(&mal.Secret, "authentication-token", "t", "", "MyAnimeList authentication token to use (overrides system keyring if any)")
}
