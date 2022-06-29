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
  mal "github.com/MikunoNaka/macli/mal"
  a "github.com/MikunoNaka/MAL2Go/v2/anime"
  m "github.com/MikunoNaka/MAL2Go/v2/manga"
)

// only search animes probably only now
func AnimeSearch(label, searchString string) a.Anime {
  // TODO: load promptLength from config
  promptLength := 5
  animes := mal.SearchAnime(searchString)

  template := &p.SelectTemplates {
    Label: "{{ . }}",
    Active: "{{ .Title | magenta }}",
    Inactive: "{{ .Title }}",
    Selected: "{{ .Title | blue }}",
    Details: `
--------- {{ .Title }} ----------
More Details To Be Added Later
`,
  }

  // returns true if input == anime title
  searcher := func(input string, index int) bool {
    title := strings.Replace(strings.ToLower(animes[index].Title), " ", "", -1)
    input = strings.Replace(strings.ToLower(input), " ", "", -1)
    return strings.Contains(title, input)
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
    fmt.Println("Error running search menu.", err.Error())
    os.Exit(1)
  }

  return animes[animeIndex]
}

func MangaSearch(label, searchString string) m.Manga {
  // TODO: load promptLength from config
  promptLength := 5
  mangas := mal.SearchManga(searchString)

  template := &p.SelectTemplates {
    Label: "{{ . }}",
    Active: "{{ .Title | magenta }}",
    Inactive: "{{ .Title }}",
    Selected: "{{ .Title | blue }}",
    Details: `
--------- {{ .Title }} ----------
More Details To Be Added Later
`,
  }

  // returns true if input == anime title
  searcher := func(input string, index int) bool {
    title := strings.Replace(strings.ToLower(mangas[index].Title), " ", "", -1)
    input = strings.Replace(strings.ToLower(input), " ", "", -1)
    return strings.Contains(title, input)
  }

  prompt := p.Select {
    Label: label,
    Items: mangas,
    Templates: template,
    Searcher: searcher,
    Size: promptLength,
  }

  mangaIndex, _, err := prompt.Run()
  if err != nil {
    fmt.Println("Error running search menu.", err.Error())
    os.Exit(1)
  }

  return mangas[mangaIndex]
}
