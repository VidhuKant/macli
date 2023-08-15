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
var volumesCmd = &cobra.Command{
	Use:   "volumes",
	Short: "Set the number of volumes read",
	Long: "Set the number of volumes read" +
	"\n" +
    "Example Usage:\n" +
	" - \x1b[33m`macli volumes <manga-name>`\x1b[0m For interactive prompt (manga-name can be omitted)\n" +
	" - \x1b[33m`macli volumes -s 4 <manga-name>`\x1b[0m to set the volumes to 4\n" +
	" - \x1b[33m`macli volumes -s +1 <manga-name>`\x1b[0m to increment the volumes by 1\n" +
	" - \x1b[33m`macli volumes -s -2 <manga-name>`\x1b[0m to decrement the volumes by 2\n" +
	" - \x1b[33m`macli volumes <manga-name> -S 1`\x1b[0m automatically selects the first search result\n",
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := util.BindSearchConfig(cmd.Flags())
		if err != nil {
			fmt.Println("Error while parsing flags.", err.Error())
			os.Exit(1)
		}
		mal.SearchLength = conf.SearchLength
		mal.SearchOffset = conf.SearchOffset
 		mal.AutoSel = conf.AutoSel
		mal.SearchNSFW = conf.SearchNSFW
		ui.PromptLength = conf.PromptLength
    mal.Init()

		var selectedManga mg.Manga
		if entryId > 0 {
			selectedManga = mal.GetMangaData(entryId, []string{"my_list_status", "num_volumes"})
		}

		searchInput := strings.Join(args, " ")
	    if searchInput == "" && entryId < 1 {
		    var promptText string
			if queryOnlyMode {
				promptText = "Search Manga to Get Amount of Volumes Read For: "
			} else {
				promptText = "Search Manga to Set Volumes For: "
			}
	    	searchInput = ui.TextInput(promptText, "Search can't be blank.")
	    }

		var volInput string
		if !queryOnlyMode {
		  volInput, err = cmd.Flags().GetString("set-value")
		  if err != nil {
		  	fmt.Println("Error while reading \x1b[33m--set-value\x1b[0m flag.", err.Error())
		  	os.Exit(1)
		  }
		}

		if entryId < 1 {
			manga := ui.MangaSearch("Select Manga:", searchInput)
			selectedManga = mal.GetMangaData(manga.Id, []string{"my_list_status", "num_volumes"})
		}

		if queryOnlyMode {
		    fmt.Printf("%s / \x1b[1;33m%d\n", ui.CreateVolumeUpdateConfirmationMessage(selectedManga.Title, selectedManga.MyListStatus.VolumesRead, -1), selectedManga.NumVolumes)
			os.Exit(0)
		}

		if volInput == "" {
			ui.VolumeInput(selectedManga)
		} else {
			resp := mal.SetVolumes(selectedManga.Id, selectedManga.MyListStatus.VolumesRead, volInput)
		    fmt.Println(ui.CreateVolumeUpdateConfirmationMessage(selectedManga.Title, resp.VolumesRead, selectedManga.MyListStatus.VolumesRead))
		}
	},
}

func init() {
	rootCmd.AddCommand(volumesCmd)
    volumesCmd.Flags().StringP("set-value", "s", "", "Number of voulmes")
    volumesCmd.Flags().BoolVarP(&queryOnlyMode, "query", "q", false, "Query only (don't update data)")
    volumesCmd.Flags().IntVarP(&entryId, "id", "i", -1, "Manually specify the ID of manga (overrides search)")
    volumesCmd.Flags().StringVarP(&mal.Secret, "authentication-token", "t", "", "MyAnimeList authentication token to use (overrides system keyring if any)")

    volumesCmd.Flags().IntP("auto-select", "S", 0, "Automatically select nth value")
    volumesCmd.Flags().IntP("prompt-length", "l", 5, "Length of select prompt")
    volumesCmd.Flags().IntP("search-length", "n", 10, "Amount of search results to load")
    volumesCmd.Flags().IntP("search-offset", "o", 0, "Offset for the search results")
    volumesCmd.Flags().BoolP("search-nsfw", "", false, "Include NSFW-rated items in search results")
}
