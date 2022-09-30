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
	"github.com/MikunoNaka/macli/util"
	"github.com/MikunoNaka/macli/mal"
	a "github.com/MikunoNaka/MAL2Go/v4/anime"

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
		conf, err := util.BindSearchConfig(cmd.Flags())
		if err != nil {
			fmt.Println("Error while parsing flags.", err.Error())
			os.Exit(1)
		}
		mal.SearchLength = conf.SearchLength
		mal.SearchOffset = conf.SearchOffset
		mal.SearchNSFW = conf.SearchNSFW
		ui.PromptLength = conf.PromptLength
    mal.Init()

		var selectedAnime a.Anime
		if entryId > 0 {
			selectedAnime = mal.GetAnimeData(entryId, []string{"my_list_status", "num_episodes"})
		}

		searchInput := strings.Join(args, " ")
	    if searchInput == "" && entryId < 1 {
		    var promptText string
			if queryOnlyMode {
				promptText = "Search Anime to Get Amount of Episodes Watched for: "
			} else {
				promptText = "Search Anime To Set Episodes For: "
			}
	    	searchInput = ui.TextInput(promptText, "Search can't be blank.")
	    }

		var epInput string
		if !queryOnlyMode {
		    epInput, err = cmd.Flags().GetString("set-value")
		    if err != nil {
		    	fmt.Println("Error while reading \x1b[33m--set-value\x1b[0m flag.", err.Error())
		    	os.Exit(1)
		    }
		}

		if entryId < 1 {
			anime := ui.AnimeSearch("Select Anime:", searchInput)
			selectedAnime = mal.GetAnimeData(anime.Id, []string{"my_list_status", "num_episodes"})
		}

		if queryOnlyMode {
		    fmt.Printf("%s / \x1b[1;33m%d\n", ui.CreateEpisodeUpdateConfirmationMessage(selectedAnime.Title, selectedAnime.MyListStatus.EpWatched, -1), selectedAnime.NumEpisodes)
			os.Exit(0)
		}

		if epInput == "" {
			ui.EpisodeInput(selectedAnime)
		} else {
			resp := mal.SetEpisodes(selectedAnime.Id, selectedAnime.MyListStatus.EpWatched, epInput)
		    fmt.Println(ui.CreateEpisodeUpdateConfirmationMessage(selectedAnime.Title, resp.EpWatched, selectedAnime.MyListStatus.EpWatched))
		}
	},
}

func init() {
	rootCmd.AddCommand(episodesCmd)
    episodesCmd.Flags().StringP("set-value", "s", "", "Number of episodes")
    episodesCmd.Flags().BoolVarP(&queryOnlyMode, "query", "q", false, "Query only (don't update data)")
    episodesCmd.Flags().IntVarP(&entryId, "id", "i", -1, "Manually specify the ID of anime/manga (overrides search)")
    episodesCmd.Flags().StringVarP(&mal.Secret, "authentication-token", "t", "", "MyAnimeList authentication token to use (overrides system keyring if any)")

    episodesCmd.Flags().IntP("prompt-length", "l", 5, "Length of select prompt")
    episodesCmd.Flags().IntP("search-length", "n", 10, "Amount of search results to load")
    episodesCmd.Flags().IntP("search-offset", "o", 0, "Offset for the search results")
    episodesCmd.Flags().BoolP("search-nsfw", "", false, "Include NSFW-rated items in search results")
}
