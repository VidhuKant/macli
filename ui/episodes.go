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

package ui

import (
  "strconv"
  "fmt"
  "os"
  "errors"
  "github.com/MikunoNaka/macli/mal"
  a "github.com/MikunoNaka/MAL2Go/v2/anime"
  m "github.com/MikunoNaka/MAL2Go/v2/manga"
  p "github.com/manifoldco/promptui"
)

// very short name I know
func CreateEpisodeUpdateConfirmationMessage(title string, prevEpNum, epNum int) string {
  return fmt.Sprintf("Set Episodes Watched for \x1b[35m%s\x1b[0m from \x1b[1;33m%d\x1b[0m to \x1b[1;36m%d\x1b[0m.", title, prevEpNum, epNum)
}


func EpisodeInput(anime a.Anime) {
  // fetch number of total episodes, number of watched episodes from the API
  animeData := mal.GetAnimeData(anime.Id, []string{"num_episodes", "my_list_status"})
  epWatchedNum := animeData.MyListStatus.EpWatched
  epTotalNum := animeData.NumEpisodes

  validate := func(input string) error {
    if _, err := strconv.ParseFloat(input, 64); err != nil {
      return errors.New("Input must be a number.")
    }
    return nil
  }

  template := &p.PromptTemplates {
    Valid: "\x1b[0m{{ . | magenta }}",
    Invalid: "\x1b[0m{{ . | magenta }}\x1b[31m ",
    Success: "{{ . | cyan }}",
  }

  prompt := p.Prompt {
    Label: fmt.Sprintf("Set Episode Number (%d/%d watched): ", epWatchedNum, epTotalNum),
    Templates: template,
    Validate:  validate,
  }

  res, err := prompt.Run()
  if err != nil {
    fmt.Println("Error Running episode input Prompt.", err.Error())
    os.Exit(1)
  }

  resp := mal.SetEpisodes(anime.Id, epWatchedNum, res)
  fmt.Println(CreateEpisodeUpdateConfirmationMessage(anime.Title, epWatchedNum, resp.EpWatched))
}

func ChapterInput(manga m.Manga) {
  // fetch number of total chapters, number of read chapters from the API
  animeData := mal.GetMangaData(manga.Id, []string{"num_chapters", "my_list_status"})
  chReadNum := animeData.MyListStatus.ChaptersRead
  chTotalNum := animeData.NumChapters

  validate := func(input string) error {
    if _, err := strconv.ParseFloat(input, 64); err != nil {
      return errors.New("Input must be a number.")
    }
    return nil
  }

  template := &p.PromptTemplates {
    Valid: "\x1b[0m{{ . | magenta }}",
    Invalid: "\x1b[0m{{ . | magenta }}\x1b[31m ",
    Success: "{{ . | cyan }}",
  }

  prompt := p.Prompt {
    Label: fmt.Sprintf("Set Chapter Number (%d/%d read): ", chReadNum, chTotalNum),
    Templates: template,
    Validate:  validate,
  }

  res, err := prompt.Run()
  if err != nil {
    fmt.Println("Error Running chapter input Prompt.", err.Error())
    os.Exit(1)
  }

  // TODO: read resp and show confirmation message
  mal.SetChapters(manga.Id, chReadNum, res)
}
