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
	if searchInput == "" {
	    var promptText string
		if queryOnlyMode {
			promptText = "Search Anime to Get Status of: "
		} else {
			promptText = "Search Anime to Set Status of: "
		}
	   	searchInput = ui.TextInput(promptText, "Search can't be blank.")
	}

	anime := ui.AnimeSearch("Select Anime:", searchInput)
	selectedAnime := mal.GetAnimeData(anime.Id, []string{"my_list_status"})

	if queryOnlyMode {
		status := selectedAnime.MyListStatus.Status
		// fmt.Printf("Anime: \x1b[35m%s\x1b[0m, Status: %s%s\x1b[0m\n", anime.Title, ui.GetColorCodeByStatus(status), ui.FormatStatus(status))
		fmt.Printf("\x1b[35m%s\x1b[0m :: %s%s\x1b[0m\n", anime.Title, ui.GetColorCodeByStatus(status), ui.FormatStatus(status))
		os.Exit(0)
	}

    if statusInput == "" {
    	ui.AnimeStatusMenu(selectedAnime)
    } else {
    	resp := mal.SetAnimeStatus(anime.Id, statusInput)
    	fmt.Println(ui.CreateStatusUpdateConfirmationMessage(anime.Title, resp.Status))
    }
}

func setMangaStatus(statusInput, searchInput string) {
	if searchInput == "" {
	    var promptText string
		if queryOnlyMode {
			promptText = "Search Manga to Get Status of: "
		} else {
			promptText = "Search Manga to Set Status of: "
		}
	   	searchInput = ui.TextInput(promptText, "Search can't be blank.")
	}

	manga := ui.MangaSearch("Select Manga:", searchInput)
	selectedManga := mal.GetMangaData(manga.Id, []string{"my_list_status"})

	if queryOnlyMode {
		status := selectedManga.MyListStatus.Status
		// fmt.Printf("Manga: \x1b[35m%s\x1b[0m, Status: %s%s\x1b[0m\n", manga.Title, ui.GetColorCodeByStatus(status), ui.FormatStatus(status))
		fmt.Printf("\x1b[35m%s\x1b[0m :: %s%s\x1b[0m\n", manga.Title, ui.GetColorCodeByStatus(status), ui.FormatStatus(status))
		os.Exit(0)
	}

    if statusInput == "" {
    	ui.MangaStatusMenu(selectedManga)
    } else {
    	resp := mal.SetMangaStatus(selectedManga.Id, statusInput)
		fmt.Println(resp.Status)
    	fmt.Println(ui.CreateStatusUpdateConfirmationMessage(manga.Title, resp.Status))
    }
}

func init() {
	rootCmd.AddCommand(statusCmd)
    statusCmd.Flags().StringP("set-value", "s", "", "status to be set")
    statusCmd.PersistentFlags().IntVarP(&ui.PromptLength, "prompt-length", "l", 5, "Length of select prompt")
    statusCmd.PersistentFlags().IntVarP(&mal.SearchLength, "search-length", "n", 10, "Amount of search results to load")
    statusCmd.PersistentFlags().IntVarP(&mal.SearchOffset, "search-offset", "o", 0, "Offset for the search results")
    statusCmd.PersistentFlags().BoolVarP(&mal.SearchNSFW, "search-nsfw", "", false, "Include NSFW-rated items in search results")
    statusCmd.PersistentFlags().BoolVarP(&mangaMode, "manga", "m", false, "Use manga mode")
    statusCmd.PersistentFlags().BoolVarP(&queryOnlyMode, "query", "q", false, "Query only (don't update data)")
}
