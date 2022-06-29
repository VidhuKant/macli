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

func EpisodeInput(anime a.Anime) {
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
    Label: "Set Episode Number: ",
    Templates: template,
    Validate:  validate,
  }

  // print current episode number if any
  epNum := anime.MyListStatus.EpWatched
  if epNum != 0 {
    fmt.Printf("\x1b[33mYou currently have watched %d episodes.\n\x1b[0m", epNum)
  }

  res, err := prompt.Run()
  if err != nil {
    fmt.Println("Error Running episode input Prompt.", err.Error())
    os.Exit(1)
  }

  mal.SetEpisodes(anime.Id, res)
}

func ChapterInput(manga m.Manga) {
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
    Label: "Set Chapter Number: ",
    Templates: template,
    Validate:  validate,
  }

  // print current chapter number if any
  chNum := manga.MyListStatus.ChaptersRead
  if chNum != 0 {
    fmt.Printf("\x1b[33mYou currently have read %d chapters.\n\x1b[0m", chNum)
  }

  res, err := prompt.Run()
  if err != nil {
    fmt.Println("Error Running chapter input Prompt.", err.Error())
    os.Exit(1)
  }

  mal.SetChapters(manga.Id, res)
}
