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
  "log"
  "github.com/MikunoNaka/macli/mal"
  p "github.com/manifoldco/promptui"
)

type StatusOption struct {
  Label  string
  Status string
}

// only search animes probably only now
func StatusMenu(animeId int) {
  options := []StatusOption {
    {"Watching", "watching"},
    {"Completed", "completed"},
    {"On Hold", "on_hold"},
    {"Dropped", "dropped"},
    {"Plan to Watch", "plan_to_watch"},
  }

  template := &p.SelectTemplates {
    Label: "{{ .Label }}",
    Active: "{{ .Label | magenta }}",
    Inactive: "{{ .Label }}",
    Selected: "{{ .Label }}",
  }

  // returns true if input == anime title
  searcher := func(input string, index int) bool {
    status := strings.Replace(strings.ToLower(options[index].Label), " ", "", -1)
    input = strings.Replace(strings.ToLower(input), " ", "", -1)
    return strings.Contains(status, input)
  }

  prompt := p.Select {
    Label: "Set Status:",
    Items: options,
    Templates: template,
    Searcher: searcher,
    Size: 5,
  }

  res, _, err := prompt.Run()
  if err != nil {
    log.Println(err)
    return
  }

  mal.SetStatus(animeId, options[res].Status)
}
