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
	"strings"
	"vidhukant.com/macli/ui"
	"vidhukant.com/macli/util"
	"vidhukant.com/macli/mal"
  "vidhukant.com/mg"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var chaptersCmd = &cobra.Command{
	Use:   "chapters",
	Short: "Set the number of chapters read",
	Long: "Set the number of chapters read" +
	"\n" +
    "Example Usage:\n" +
	" - \x1b[33m`macli chapters <manga-name>`\x1b[0m For interactive prompt (anime-name can be omitted)\n" +
	" - \x1b[33m`macli chapters -s 4 <manga-name>`\x1b[0m to set the chapters to 4\n" +
	" - \x1b[33m`macli chapters -s +1 <manga-name>`\x1b[0m to increment the chapters by 1\n" +
	" - \x1b[33m`macli chapters -s -2 <manga-name>`\x1b[0m to decrement the chapters by 2\n" +
	" - \x1b[33m`macli chapters <manga-name> -S 1`\x1b[0m automatically selects the first search result\n",
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := util.BindSearchConfig(cmd.Flags())
		if err != nil {
			fmt.Println("Error while parsing flags.", err.Error())
			os.Exit(1)
		}
		mal.SearchLength = conf.SearchLength
		mal.SearchOffset = conf.SearchOffset
		mal.SearchNSFW = conf.SearchNSFW
 		mal.AutoSel = conf.AutoSel
		ui.PromptLength = conf.PromptLength
    mal.Init()

		var selectedManga mg.Manga
		if entryId > 0 {
			selectedManga = mal.GetMangaData(entryId, []string{"my_list_status", "num_chapters"})
		}

		searchInput := strings.Join(args, " ")
	    if searchInput == "" && entryId < 1 {
		    var promptText string
			if queryOnlyMode {
				promptText = "Search Manga to Get Amount of Chapters Read For: "
			} else {
				promptText = "Search Manga to Set Chapters For: "
			}
	    	searchInput = ui.TextInput(promptText, "Search can't be blank.")
	    }

		var chInput string
		if !queryOnlyMode {
		  chInput, err = cmd.Flags().GetString("set-value")
		  if err != nil {
		  	fmt.Println("Error while reading \x1b[33m--set-value\x1b[0m flag.", err.Error())
		  	os.Exit(1)
		  }
		}

		if entryId < 1 {
			manga := ui.MangaSearch("Select Manga:", searchInput)
			selectedManga = mal.GetMangaData(manga.Id, []string{"my_list_status", "num_chapters"})
		}

		if queryOnlyMode {
		    fmt.Printf("%s / \x1b[1;33m%d\n", ui.CreateChapterUpdateConfirmationMessage(selectedManga.Title, selectedManga.MyListStatus.ChaptersRead, -1), selectedManga.NumChapters)
			os.Exit(0)
		}

		if chInput == "" {
			ui.ChapterInput(selectedManga)
		} else {
			resp := mal.SetChapters(selectedManga.Id, selectedManga.MyListStatus.ChaptersRead, chInput)
		    fmt.Println(ui.CreateChapterUpdateConfirmationMessage(selectedManga.Title, resp.ChaptersRead, selectedManga.MyListStatus.ChaptersRead))
		}
	},
}

func init() {
	rootCmd.AddCommand(chaptersCmd)
    chaptersCmd.Flags().StringP("set-value", "s", "", "Number of chapters")
    chaptersCmd.Flags().BoolVarP(&queryOnlyMode, "query", "q", false, "Query only (don't update data)")
    chaptersCmd.Flags().IntVarP(&entryId, "id", "i", -1, "Manually specify the ID of anime/manga (overrides search)")
    chaptersCmd.Flags().StringVarP(&mal.Secret, "authentication-token", "t", "", "MyAnimeList authentication token to use (overrides system keyring if any)")

    chaptersCmd.Flags().IntP("auto-select", "S", 0, "Automatically select nth value")
    chaptersCmd.Flags().IntP("prompt-length", "l", 5, "Length of select prompt")
    chaptersCmd.Flags().IntP("search-length", "n", 10, "Amount of search results to load")
    chaptersCmd.Flags().IntP("search-offset", "o", 0, "Offset for the search results")
    chaptersCmd.Flags().BoolP("search-nsfw", "", false, "Include NSFW-rated items in search results")
}
