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
	"github.com/MikunoNaka/macli/ui"
	"github.com/MikunoNaka/macli/mal"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var episodesCmd = &cobra.Command{
	Use:   "episodes",
	Short: "Set the number of episodes watched",
	Long: "Set the number of episodes watched" +
	"\n" +
    "Example Usage:\n" +
	" - \x1b[33m`macli episodes <anime-name>`\x1b[0m For interactive prompt (anime-name can be omitted)\n" +
	" - \x1b[33m`macli episodes -s 4 <anime-name>`\x1b[0m to set the episodes to 4\n" +
	" - \x1b[33m`macli episodes -s +1 <anime-name>`\x1b[0m to increment the episodes by 1\n" +
	" - \x1b[33m`macli episodes -s -2 <anime-name>`\x1b[0m to decrement the episodes by 2\n",
	Run: func(cmd *cobra.Command, args []string) {
    mal.Init()
		searchInput := strings.Join(args, " ")
	    if searchInput == "" {
		    var promptText string
			if queryOnlyMode {
				promptText = "Search Anime to Get Amount of Episodes Watched for: "
			} else {
				promptText = "Search Anime To Set Episodes For: "
			}
	    	searchInput = ui.TextInput(promptText, "Search can't be blank.")
	    }

		var (
			epInput string
			err error
		)
		if !queryOnlyMode {
		    epInput, err = cmd.Flags().GetString("set-value")
		    if err != nil {
		    	fmt.Println("Error while reading \x1b[33m--set-value\x1b[0m flag.", err.Error())
		    	os.Exit(1)
		    }
		}

	    anime := ui.AnimeSearch("Select Anime:", searchInput)
		selectedAnime := mal.GetAnimeData(anime.Id, []string{"my_list_status", "num_episodes"})
		prevEpWatched := selectedAnime.MyListStatus.EpWatched

		if queryOnlyMode {
			fmt.Printf("\x1b[35m%s\x1b[0m Episodes :: \x1b[1;36m%d\x1b[0m / \x1b[1;33m%d\x1b[0m Watched\n", anime.Title, prevEpWatched, selectedAnime.NumEpisodes)
			os.Exit(0)
		}

		if epInput == "" {
			ui.EpisodeInput(selectedAnime)
		} else {
			resp := mal.SetEpisodes(anime.Id, prevEpWatched, epInput)
		    fmt.Println(ui.CreateEpisodeUpdateConfirmationMessage(anime.Title, prevEpWatched, resp.EpWatched))
		}


	},
}

func init() {
	rootCmd.AddCommand(episodesCmd)
    episodesCmd.Flags().StringP("set-value", "s", "", "Number of episodes")
    episodesCmd.Flags().IntVarP(&ui.PromptLength, "prompt-length", "l", 5, "Length of select prompt")
    episodesCmd.Flags().BoolVarP(&queryOnlyMode, "query", "q", false, "Query only (don't update data)")
    episodesCmd.Flags().IntVarP(&mal.SearchLength, "search-length", "n", 10, "Amount of search results to load")
    episodesCmd.Flags().BoolVarP(&mal.SearchNSFW, "search-nsfw", "", false, "Include NSFW-rated items in search results")
    episodesCmd.Flags().IntVarP(&mal.SearchOffset, "search-offset", "o", 0, "Offset for the search results")
}
