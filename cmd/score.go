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
	"strconv"
	"errors"
	"github.com/MikunoNaka/macli/ui"
	"github.com/MikunoNaka/macli/mal"
	"github.com/MikunoNaka/macli/util"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var scoreCmd = &cobra.Command{
	Use:   "score",
	Short: "Set an anime/manga's status",
	Long: "Set an anime's status\n" +
	"\n" +
    "Example Usage:\n" +
	" - \x1b[33m`macli status <anime-name>`\x1b[0m For interactive prompt (anime-name can be omitted)\n" +
	" - \x1b[33m`macli status -s \x1b[34mwatching|plan_to_watch|dropped|on_hold|completed\x1b[33m <anime-name>`\x1b[0m to specify status from command\n",
	Run: func(cmd *cobra.Command, args []string) {
		mal.Init()
		searchInput := strings.Join(args, " ")

		scoreInput, err := cmd.Flags().GetString("set-value")
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
			setMangaScore(scoreInput, searchInput)
		} else {
			setAnimeScore(scoreInput, searchInput)
		}

	},
}

func validateScore(input string, currentScore int) (int, error) {
	parsedScore := util.ParseNumeric(input, currentScore)
	if parsedScore > 10 {
        return currentScore, errors.New("\x1b[31mScore out of range (" + strconv.Itoa(parsedScore) + " > 10)\x1b[0m")
	}
	if parsedScore < 0 {
        return currentScore, errors.New("\x1b[31mScore out of range (" + strconv.Itoa(parsedScore) + " < 0)\x1b[0m")
	}
	return parsedScore, nil
}

func setAnimeScore(scoreInput, searchInput string) {
	if searchInput == "" {
	    var promptText string
		if queryOnlyMode {
			promptText = "Search Anime to Get Score of: "
		} else {
			promptText = "Search Anime to Set Score of: "
		}
	   	searchInput = ui.TextInput(promptText, "Search can't be blank.")
	}

	anime := ui.AnimeSearch("Select Anime: ", searchInput)
	selectedAnime := mal.GetAnimeData(anime.Id, []string{"my_list_status"})
	currentScore := selectedAnime.MyListStatus.Score

	if queryOnlyMode {
		fmt.Printf("\x1b[35m%s\x1b[0m Score :: %s\n", anime.Title, ui.FormatScore(currentScore))
		os.Exit(0)
	}

    if scoreInput == "" {
		ui.AnimeScoreInput(selectedAnime)
    } else {
		score, err := validateScore(scoreInput, selectedAnime.MyListStatus.Score)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
    	resp := mal.SetAnimeScore(anime.Id, score)
    	fmt.Println(ui.CreateScoreUpdateConfirmationMessage(anime.Title, currentScore, resp.Score))
    }
}

func setMangaScore(scoreInput, searchInput string) {
	if searchInput == "" {
	    var promptText string
		if queryOnlyMode {
			promptText = "Search Manga to Get Score of: "
		} else {
			promptText = "Search Manga to Set Score of: "
		}
	   	searchInput = ui.TextInput(promptText, "Search can't be blank.")
	}

	manga := ui.MangaSearch("Select Manga: ", searchInput)
	selectedManga := mal.GetMangaData(manga.Id, []string{"my_list_status"})
	currentScore := selectedManga.MyListStatus.Score

	if queryOnlyMode {
		fmt.Printf("\x1b[35m%s\x1b[0m Score :: %s\n", manga.Title, ui.FormatScore(currentScore))
		os.Exit(0)
	}

    if scoreInput == "" {
		ui.MangaScoreInput(selectedManga)
    } else {
		score, err := validateScore(scoreInput, selectedManga.MyListStatus.Score)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
    	resp := mal.SetMangaScore(manga.Id, score)
    	fmt.Println(ui.CreateScoreUpdateConfirmationMessage(manga.Title, currentScore, resp.Score))
    }
}

func init() {
	rootCmd.AddCommand(scoreCmd)
    scoreCmd.Flags().StringP("set-value", "s", "", "Score to be set")
}
