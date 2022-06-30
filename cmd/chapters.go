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
var chaptersCmd = &cobra.Command{
	Use:   "chapters",
	Short: "Set the number of chapters read",
	Long: "Set the number of chapters read" +
	"\n" +
    "Example Usage:\n" +
	" - \x1b[33m`macli chapters <anime-name>`\x1b[0m For interactive prompt (anime-name can be omitted)\n" +
	" - \x1b[33m`macli chapters -s 4 <anime-name>`\x1b[0m to set the chapters to 4\n" +
	" - \x1b[33m`macli chapters -s +1 <anime-name>`\x1b[0m to increment the chapters by 1\n" +
	" - \x1b[33m`macli chapters -s -2 <anime-name>`\x1b[0m to decrement the chapters by 2\n",
	Run: func(cmd *cobra.Command, args []string) {
		searchInput := strings.Join(args, " ")
	    if searchInput == "" {
	    	searchInput = ui.TextInput("Search Manga To Set Chapters For: ", "Search can't be blank.")
	    }

		chInput, err := cmd.Flags().GetString("set-value")
		if err != nil {
			fmt.Println("Error while reading --set-value flag.", err.Error())
		}

	    manga := ui.MangaSearch("Select Manga:", searchInput)
		mangaData := mal.GetMangaData(manga.Id, []string{"my_list_status"})
		prevChRead := mangaData.MyListStatus.ChaptersRead

		if chInput == "" {
			ui.ChapterInput(manga)
		} else {
			resp := mal.SetChapters(manga.Id, prevChRead, chInput)
		    fmt.Println(ui.CreateChapterUpdateConfirmationMessage(manga.Title, prevChRead, resp.ChaptersRead))
		}
	},
}

func init() {
	rootCmd.AddCommand(chaptersCmd)
    chaptersCmd.Flags().StringP("set-value", "s", "", "Number of chapters")
}
