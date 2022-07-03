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
  "strings"
  "fmt"
  "os"
  p "github.com/manifoldco/promptui"
  a "github.com/MikunoNaka/MAL2Go/v2/anime"
  m "github.com/MikunoNaka/MAL2Go/v2/manga"
  "github.com/MikunoNaka/macli/mal"
)

type AnimeAction struct {
  Label       string
  Description string
  Method      func(a.Anime)
}

type MangaAction struct {
  Label       string
  Description string
  Method      func(m.Manga)
}

func AnimeActionMenu(animeIsAdded bool) func(a.Anime) {
  options := []AnimeAction {
    {"Set Status",          "Set status for an anime (watching, dropped, etc)", AnimeStatusMenu},
    {"Set Episodes",        "Set number of episodes watched", EpisodeInput},
    {"Set Score",           "Set score", AnimeScoreInput},
    // {"Set Re-watching",      "Set if re-watching", AnimeStatusMenu},
    // {"Set Times Re-watched", "Set number of times re-watched", AnimeStatusMenu},
  }

  // if anime not in list
  if animeIsAdded {
    options = append(
      options,
      AnimeAction{"Delete Anime", "Delete Anime From Your MyAnimeList List.", mal.DeleteAnime},
    )
  }

  template := &p.SelectTemplates {
    Label: "{{ .Label }}",
    Active: "{{ .Label | magenta }} {{ .Description | faint }}",
    Inactive: "{{ .Label }}",
    Selected: "{{ .Label | magenta }}",
    Details: `
-------------------
{{ .Description }}
`,
  }

  // returns true if input == anime title
  searcher := func(input string, index int) bool {
    action := strings.Replace(strings.ToLower(options[index].Label), " ", "", -1)
    input = strings.Replace(strings.ToLower(input), " ", "", -1)
    return strings.Contains(action, input)
  }

  prompt := p.Select {
    Label: "Select Action: ",
    Items: options,
    Templates: template,
    Searcher: searcher,
    Size: PromptLength,
  }

  res, _, err := prompt.Run()
  if err != nil {
    fmt.Println("Error running actions menu.", err.Error())
    os.Exit(1)
  }

  return options[res].Method
}

func MangaActionMenu(mangaIsAdded bool) func(m.Manga) {
  options := []MangaAction {
    {"Set Status",        "Set status for a manga (reading, dropped, etc)", MangaStatusMenu},
    {"Set Chapters",      "Set number of chapters read", ChapterInput},
    {"Set Score",         "Set score", MangaScoreInput},
    // {"Set Re-reading",    "Set if re-reading", MangaStatusMenu},
    // {"Set Times Re-read", "Set number of times re-read", MangaStatusMenu},
  }

  // if manga not in list
  if mangaIsAdded {
    options = append(
      options,
      MangaAction{"Delete Manga", "Delete Manga From Your MyAnimeList List.", mal.DeleteManga},
    )
  }

  template := &p.SelectTemplates {
    Label: "{{ .Label }}",
    Active: "{{ .Label | magenta }} {{ .Description | faint }}",
    Inactive: "{{ .Label }}",
    Selected: "{{ .Label | magenta }}",
    Details: `
-------------------
{{ .Description }}
`,
  }

  // returns true if input == anime title
  searcher := func(input string, index int) bool {
    action := strings.Replace(strings.ToLower(options[index].Label), " ", "", -1)
    input = strings.Replace(strings.ToLower(input), " ", "", -1)
    return strings.Contains(action, input)
  }

  prompt := p.Select {
    Label: "Select Action: ",
    Items: options,
    Templates: template,
    Searcher: searcher,
    Size: PromptLength,
  }

  res, _, err := prompt.Run()
  if err != nil {
    fmt.Println("Error running actions menu.", err.Error())
    os.Exit(1)
  }

  return options[res].Method
}
