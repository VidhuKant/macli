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
  "vidhukant.com/macli/mal"
  "vidhukant.com/mg"
  p "github.com/manifoldco/promptui"
)

// very short name I know
func CreateEpisodeUpdateConfirmationMessage(title string, epNum, prevEpNum int) string {
  if prevEpNum >= 0 {
    return fmt.Sprintf("\x1b[35m%s\x1b[0m Episodes Watched :: \x1b[1;33m%d\x1b[0m -> \x1b[1;36m%d\x1b[0m", title, prevEpNum, epNum)
  }
  return fmt.Sprintf("\x1b[35m%s\x1b[0m Episodes Watched :: \x1b[1;36m%d\x1b[0m", title, epNum)
}

func CreateChapterUpdateConfirmationMessage(title string, chNum, prevChNum int) string {
  if prevChNum >= 0 {
    return fmt.Sprintf("\x1b[35m%s\x1b[0m Chapters Read :: \x1b[1;33m%d\x1b[0m -> \x1b[1;36m%d\x1b[0m", title, prevChNum, chNum)
  }
  return fmt.Sprintf("\x1b[35m%s\x1b[0m Chapters Read :: \x1b[1;36m%d\x1b[0m", title, chNum)
}

func CreateVolumeUpdateConfirmationMessage(title string, volNum, prevVolNum int) string {
  if prevVolNum >= 0 {
    return fmt.Sprintf("\x1b[35m%s\x1b[0m Volumes Read :: \x1b[1;33m%d\x1b[0m -> \x1b[1;36m%d\x1b[0m", title, prevVolNum, volNum)
  }
  return fmt.Sprintf("\x1b[35m%s\x1b[0m Volumes Read :: \x1b[1;36m%d\x1b[0m", title, volNum)
}

func EpisodeInput(anime mg.Anime) {
  epWatchedNum := anime.MyListStatus.EpWatched
  epTotalNum := anime.NumEpisodes

  validate := func(input string) error {
    if _, err := strconv.ParseFloat(input, 64); err != nil {
      return errors.New("Input must be a number.")
    }
    return nil
  }

  template := &p.PromptTemplates {
    Valid: "\x1b[0m{{ . | magenta }}",
    Invalid: "\x1b[0m{{ . | magenta }}\x1b[31m",
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
  fmt.Println(CreateEpisodeUpdateConfirmationMessage(anime.Title, resp.EpWatched, epWatchedNum))
}

func ChapterInput(manga mg.Manga) {
  chReadNum := manga.MyListStatus.ChaptersRead
  chTotalNum := manga.NumChapters

  validate := func(input string) error {
    if _, err := strconv.ParseFloat(input, 64); err != nil {
      return errors.New("Input must be a number.")
    }
    return nil
  }

  template := &p.PromptTemplates {
    Valid: "\x1b[0m{{ . | magenta }}",
    Invalid: "\x1b[0m{{ . | magenta }}\x1b[31m",
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

  resp := mal.SetChapters(manga.Id, chReadNum, res)
  fmt.Println(CreateChapterUpdateConfirmationMessage(manga.Title, resp.ChaptersRead, chReadNum))
}

func VolumeInput(manga mg.Manga) {
  volReadNum := manga.MyListStatus.VolumesRead
  volTotalNum := manga.NumVolumes

  validate := func(input string) error {
    if _, err := strconv.ParseFloat(input, 64); err != nil {
      return errors.New("Input must be a number.")
    }
    return nil
  }

  template := &p.PromptTemplates {
    Valid: "\x1b[0m{{ . | magenta }}",
    Invalid: "\x1b[0m{{ . | magenta }}\x1b[31m",
    Success: "{{ . | cyan }}",
  }

  prompt := p.Prompt {
    Label: fmt.Sprintf("Set Volume Number (%d/%d read): ", volReadNum, volTotalNum),
    Templates: template,
    Validate:  validate,
  }

  res, err := prompt.Run()
  if err != nil {
    fmt.Println("Error Running volume input Prompt.", err.Error())
    os.Exit(1)
  }

  resp := mal.SetVolumes(manga.Id, volReadNum, res)
  fmt.Println(CreateVolumeUpdateConfirmationMessage(manga.Title, resp.VolumesRead, volReadNum))
}
