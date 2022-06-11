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
  p "github.com/manifoldco/promptui"
  mal "github.com/MikunoNaka/macli/mal"
)

// only search animes probably only now
func SearchAndGetID(label, searchString string) int {
  // TODO: load promptLength from config
  promptLength := 5

  animes := mal.SearchAnime(searchString)

  template := &p.SelectTemplates {
    Label: "{{ . }}?",
    Active: "{{ .Title | magenta }}",
    Inactive: "{{ .Title }}",
    Selected: "{{ .Title }}",
    Details: `
--------- {{ .Title }} ----------
More Details To Be Added Later
`,
  }

  // returns true if input == anime title
  searcher := func(input string, index int) bool {
    title := strings.Replace(strings.ToLower(animes[index].Title), " ", "", -1)
    input = strings.Replace(strings.ToLower(input), " ", "", -1)
    return strings.Contains(title, title)
  }

  prompt := p.Select {
    Label: label,
    Items: animes,
    Templates: template,
    Searcher: searcher,
    Size: promptLength,
  }

  animeIndex, _, err := prompt.Run()
  if err != nil {
    log.Println(err)
    return 0
  }

  return animes[animeIndex].Id
}
