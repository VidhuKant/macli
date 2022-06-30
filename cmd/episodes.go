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
var episodesCmd = &cobra.Command{
	Use:   "episodes",
	Short: "Set an anime/manga's status",
	Long: "Set an anime's status\n" +
	"\n" +
    "Example Usage:\n" +
	" - \x1b[33m`macli status <anime-name>`\x1b[0m For interactive prompt (anime-name can be omitted)\n" +
	" - \x1b[33m`macli status -s \x1b[34mwatching|plan_to_watch|dropped|on_hold|completed\x1b[33m <anime-name>`\x1b[0m to specify status from command\n",
	Run: func(cmd *cobra.Command, args []string) {
		searchInput := strings.Join(args, " ")
	    if searchInput == "" {
	    	searchInput = ui.TextInput("Search Anime To Set Episodes For: ", "Search can't be blank.")
	    }

		epInput, err := cmd.Flags().GetString("set-value")
		if err != nil {
			fmt.Println("Error while reading --set-value flag.", err.Error())
		}

	    anime := ui.AnimeSearch("Select Anime:", searchInput)
		animeData := mal.GetAnimeData(anime.Id, []string{"my_list_status"})
		prevEpWatched := animeData.MyListStatus.EpWatched

		if epInput == "" {
			ui.EpisodeInput(anime)
		} else {
			resp := mal.SetEpisodes(anime.Id, prevEpWatched, epInput)
		    fmt.Println(ui.CreateEpisodeUpdateConfirmationMessage(anime.Title, prevEpWatched, resp.EpWatched))
		}


	},
}

func init() {
	rootCmd.AddCommand(episodesCmd)
    episodesCmd.Flags().StringP("set-value", "s", "", "status to be set")
}
