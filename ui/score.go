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
  "errors"
  "fmt"
  "os"
  "github.com/MikunoNaka/macli/mal"
  "github.com/MikunoNaka/macli/util"
  m "github.com/MikunoNaka/MAL2Go/v4/manga"
  a "github.com/MikunoNaka/MAL2Go/v4/anime"
  p "github.com/manifoldco/promptui"
)

// to print dropped in red color, etc
func GetColorCodeByScore(score int) string {
  switch score {
    case 0:
      return "\x1b[37m"
    case 1:
      return "\x1b[31m"
    case 2:
      return "\x1b[1;31m"
    case 3:
      return "\x1b[35m"
    case 4:
      return "\x1b[1;35m"
    case 5:
      return "\x1b[33m"
    case 6:
      return "\x1b[1;33m"
    case 7:
      return "\x1b[36m"
    case 8:
      return "\x1b[1;36m"
    case 9:
      return "\x1b[32m"
    case 10:
      return "\x1b[1;32m"
    default:
      return ""
  }
}

func FormatScore(score int) string {
  return fmt.Sprintf("%s%d\x1b[0m", GetColorCodeByScore(score), score)
}

// very short name I know
func CreateScoreUpdateConfirmationMessage(title string, score, prevScore int) string {
  if prevScore >= 0 {
    return fmt.Sprintf("\x1b[35m%s\x1b[0m Score :: %s -> %s", title, FormatScore(prevScore), FormatScore(score))
  }
  return fmt.Sprintf("\x1b[35m%s\x1b[0m Score :: %s", title, FormatScore(score))
}

func AnimeScoreInput(anime a.Anime) {
  currentScore := anime.MyListStatus.Score
  validate := func(input string) error {
    i, err := strconv.ParseFloat(input, 64)
    if err != nil || i < -10 || i > 10 {
      return errors.New("Input must be a number within 0-10.")
    }
    newScore := util.ParseNumeric(input, currentScore)
    if newScore < 0 {
      return errors.New("Score out of range (" + strconv.Itoa(newScore) + " < 0)")
    }
    if newScore > 10 {
      return errors.New("Score out of range (" + strconv.Itoa(newScore) + " > 10)")
    }
    return nil
  }

  template := &p.PromptTemplates {
    Valid: "\x1b[0m{{ . | magenta }}",
    Invalid: "\x1b[0m{{ . | magenta }}\x1b[31m",
    Success: "{{ . | cyan }}",
  }

  prompt := p.Prompt {
    Label: fmt.Sprintf("Set Score (Current: %d): ", currentScore),
    Templates: template,
    Validate:  validate,
  }

  res, err := prompt.Run()
  if err != nil {
    fmt.Println("Error Running score input Prompt.", err.Error())
    os.Exit(1)
  }

  resp := mal.SetAnimeScore(anime.Id, util.ParseNumeric(res, currentScore))
  fmt.Println(CreateScoreUpdateConfirmationMessage(anime.Title, resp.Score ,currentScore))
}

func MangaScoreInput(manga m.Manga) {
  currentScore := manga.MyListStatus.Score

  validate := func(input string) error {
    i, err := strconv.ParseFloat(input, 64)
    if err != nil || i < -10 || i > 10 {
      return errors.New("Input must be a number within 0-10.")
    }
    newScore := util.ParseNumeric(input, currentScore)
    if newScore < 0 {
      return errors.New("Score out of range (" + strconv.Itoa(newScore) + " < 0)")
    }
    if newScore > 10 {
      return errors.New("Score out of range (" + strconv.Itoa(newScore) + " > 10)")
    }
    return nil
  }

  template := &p.PromptTemplates {
    Valid: "\x1b[0m{{ . | magenta }}",
    Invalid: "\x1b[0m{{ . | magenta }}\x1b[31m",
    Success: "{{ . | cyan }}",
  }

  prompt := p.Prompt {
    Label: fmt.Sprintf("Set Score (Current: %d): ", currentScore),
    Templates: template,
    Validate:  validate,
  }

  res, err := prompt.Run()
  if err != nil {
    fmt.Println("Error Running score input Prompt.", err.Error())
    os.Exit(1)
  }

  resp := mal.SetMangaScore(manga.Id, util.ParseNumeric(res, currentScore))
  fmt.Println(CreateScoreUpdateConfirmationMessage(manga.Title, resp.Score ,currentScore))
}
