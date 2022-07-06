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
  "github.com/MikunoNaka/macli/mal"
  a "github.com/MikunoNaka/MAL2Go/v3/anime"
  m "github.com/MikunoNaka/MAL2Go/v3/manga"
  p "github.com/manifoldco/promptui"
)

type StatusOption struct {
  Label  string
  Status string
}

// to print dropped in red color, etc
func GetColorCodeByStatus(status string) string {
  switch status {
    case "watching", "reading":
      return "\x1b[32m"
    case "completed":
      return "\x1b[34m"
    case "on_hold":
      return "\x1b[33m"
    case "dropped":
      return "\x1b[31m"
    case "plan_to_watch", "plan_to_read":
      return "\x1b[36m"
    case "":
      return "\x1b[37m"
    default:
      return ""
  }
}

func FormatStatus(status string) string {
  switch status {
    case "watching":
      return "Watching"
    case "reading":
      return "Reading"
    case "completed":
      return "Completed"
    case "on_hold":
      return "On Hold"
    case "dropped":
      return "Dropped"
    case "plan_to_watch":
      return "Plan to Watch"
    case "plan_to_read":
      return "Plan to Read"
    case "":
      return "Not in List"
    default:
      return ""
  }
}

// very short name I know
func CreateStatusUpdateConfirmationMessage(title, status, prevStatus string) string {
  if prevStatus != "" {
    return fmt.Sprintf("\x1b[35m%s\x1b[0m Status :: %s%s\x1b[0m -> %s%s\x1b[0m", title, GetColorCodeByStatus(prevStatus), FormatStatus(prevStatus), GetColorCodeByStatus(status), FormatStatus(status))
  }
  return "\x1b[35m" + title + "\x1b[0m Status :: " + GetColorCodeByStatus(status) + FormatStatus(status) + "\x1b[0m"
}

func AnimeStatusMenu(anime a.Anime) {
  options := []StatusOption {
    {"Watching", "watching"},
    {"Completed", "completed"},
    {"On Hold", "on_hold"},
    {"Dropped", "dropped"},
    {"Plan to Watch", "plan_to_watch"},
  }

  // highlight current status (if any)
  animeStatus := anime.MyListStatus.Status
  if animeStatus != "" {
    for i := range options {
      if options[i].Status == animeStatus {
        options[i].Label = options[i].Label + " \x1b[35m\U00002714\x1b[0m"
      }
    }
  }

  template := &p.SelectTemplates {
    Label: "{{ .Label }}",
    Active: "{{ .Label | magenta }}",
    Inactive: "{{ .Label }}",
    Selected: "{{ .Label | cyan }}",
  }

  // returns true if input == anime title
  searcher := func(input string, index int) bool {
    status := strings.Replace(strings.ToLower(options[index].Label), " ", "", -1)
    input = strings.Replace(strings.ToLower(input), " ", "", -1)
    return strings.Contains(status, input)
  }

  promptLabel := "Set Status: "
  if animeStatus != "" {
    promptLabel = promptLabel + "(current - " + animeStatus + ")"
  }

  prompt := p.Select {
    Label: promptLabel,
    Items: options,
    Templates: template,
    Searcher: searcher,
    Size: 5,
  }

  res, _, err := prompt.Run()
  if err != nil {
    fmt.Println("Error running status prompt.", err.Error())
    os.Exit(1)
  }

  resp := mal.SetAnimeStatus(anime.Id, options[res].Status)
  fmt.Println(CreateStatusUpdateConfirmationMessage(anime.Title, resp.Status, anime.MyListStatus.Status))
}

func MangaStatusMenu(manga m.Manga) {
  options := []StatusOption {
    {"Reading", "reading"},
    {"Completed", "completed"},
    {"On Hold", "on_hold"},
    {"Dropped", "dropped"},
    {"Plan to Read", "plan_to_read"},
  }

  // highlight current status (if any)
  mangaStatus := manga.MyListStatus.Status
  if mangaStatus != "" {
    for i := range options {
      if options[i].Status == mangaStatus {
        options[i].Label = options[i].Label + " \x1b[35m\U00002714\x1b[0m"
      }
    }
  }

  template := &p.SelectTemplates {
    Label: "{{ .Label }}",
    Active: "{{ .Label | magenta }}",
    Inactive: "{{ .Label }}",
    Selected: "{{ .Label | cyan }}",
  }

  // returns true if input == anime title
  searcher := func(input string, index int) bool {
    status := strings.Replace(strings.ToLower(options[index].Label), " ", "", -1)
    input = strings.Replace(strings.ToLower(input), " ", "", -1)
    return strings.Contains(status, input)
  }

  promptLabel := "Set Status: "
  if mangaStatus != "" {
    promptLabel = promptLabel + "(current - " + mangaStatus + ")"
  }

  prompt := p.Select {
    Label: promptLabel,
    Items: options,
    Templates: template,
    Searcher: searcher,
    Size: 5,
  }

  res, _, err := prompt.Run()
  if err != nil {
    fmt.Println("Error running status prompt.", err.Error())
    os.Exit(1)
  }

  resp := mal.SetMangaStatus(manga.Id, options[res].Status)
  fmt.Println(CreateStatusUpdateConfirmationMessage(manga.Title, resp.Status, manga.MyListStatus.Status))
}
