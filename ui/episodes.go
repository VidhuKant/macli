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
  "log"
  "github.com/MikunoNaka/macli/mal"
  p "github.com/manifoldco/promptui"
)

func EpisodeInput(animeId int) {
  validate := func(input string) error {
    _, err := strconv.ParseFloat(input, 64)
    return err
  }

  template := &p.PromptTemplates {
    Prompt: "{{ . }} ",
    Valid: "{{ . }} ",
    Invalid: "{{ . | red }} ",
    Success: "{{ . }} ",
  }

  prompt := p.Prompt {
    // Label: "Set Episode Number: (Increment/Decrement With +1, -2, etc)",
    Label: "Set Episode Number: ",
    Templates: template,
    Validate:  validate,
  }

  res, err := prompt.Run()
  if err != nil {
    log.Println("Error Running EpisodeInput Prompt.", err)
    return
  }

  mal.SetEpisodes(animeId, res)
}

func ChapterInput(mangaId int) {
  validate := func(input string) error {
    _, err := strconv.ParseFloat(input, 64)
    return err
  }

  template := &p.PromptTemplates {
    Prompt: "{{ . }} ",
    Valid: "{{ . }} ",
    Invalid: "{{ . | red }} ",
    Success: "{{ . }} ",
  }

  prompt := p.Prompt {
    Label: "Set Chapter Number: (Increment/Decrement With +1, -2, etc)",
    Templates: template,
    Validate:  validate,
  }

  res, err := prompt.Run()
  if err != nil {
    log.Println("Error Running ChapterInput Prompt.", err)
    return
  }

  mal.SetChapters(mangaId, res)
}
