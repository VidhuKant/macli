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
  a "github.com/MikunoNaka/MAL2Go/v3/anime"
  m "github.com/MikunoNaka/MAL2Go/v3/manga"
)

// only search animes probably only now
func AnimeSearch(label, searchString string) a.Anime {
  animes := mal.SearchAnime(searchString)

  for i, anime := range animes {
    animes[i].DurationSeconds = anime.DurationSeconds / 60

    /* I cant find a way to add functions to the details template
     * So I am formatting the studios as one string
     * and setting as the first studio name. pretty hacky. */
    if len(anime.Studios) > 0 {
      var studiosFormatted string
      for j, studio := range anime.Studios {
        studiosFormatted = studiosFormatted + studio.Name
        // setting other studio names as ""
        animes[i].Studios[j].Name = ""
        if j != len(anime.Studios) - 1 {
          studiosFormatted = studiosFormatted + ", "
        }
      }
      animes[i].Studios[0].Name = studiosFormatted
    }

    var ratingFormatted string
    switch anime.Rating {
      case "g":
        ratingFormatted = "G - All Ages"
      case "pg":
        ratingFormatted = "PG - Children"
      case "pg_13":
        ratingFormatted = "PG13 - Teens 13 and Older"
      case "r":
        ratingFormatted = "R - 17+ (violence & profanity)"
      case "r+":
        ratingFormatted = "R+ - Profanity & Mild Nudity"
      case "rx":
        ratingFormatted = "Rx - Hentai"
      default:
        ratingFormatted = anime.Rating
    }
    animes[i].Rating = ratingFormatted
  }

  template := &p.SelectTemplates {
    Label: "{{ . }}",
    Active: "{{ .Title | magenta }}",
    Inactive: "{{ .Title }}",
    Selected: "{{ .Title | blue }}",
    // TODO: format and maybe color code details
    Details: `
--------- {{ .Title }} ----------
{{ "Number of Episodes:" | blue | bold }} {{ .NumEpisodes }}
{{ "English Title:" | blue | bold }} {{ .AltTitles.En }}
{{ "Japanese Title:" | blue | bold }} {{ .AltTitles.Ja }}
{{ "Original Run:" | blue | bold }} {{ .StartDate }} - {{ .EndDate }} ({{ .StartSeason.Name }} {{ .StartSeason.Year }})
{{ "Mean Score:" | blue | bold }} {{ .MeanScore }}
{{ "Rank:" | blue | bold }} {{ .Rank }}
{{ "Type:" | blue | bold }} {{ .MediaType }}
{{ "Status:" | blue | bold }} {{ .Status }}
{{ "Average Duration:" | blue | bold }} {{ .DurationSeconds }} minutes
{{ "Rating:" | blue | bold }} {{ .Rating }}
{{ "Studios:" | blue | bold }} {{ range .Studios }}{{ .Name }}{{ end }}
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
    Size: PromptLength,
  }

  animeIndex, _, err := prompt.Run()
  if err != nil {
    fmt.Println("Error running search menu.", err.Error())
    os.Exit(1)
  }

  return animes[animeIndex]
}

func MangaSearch(label, searchString string) m.Manga {
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
    Size: PromptLength,
  }

  mangaIndex, _, err := prompt.Run()
  if err != nil {
    fmt.Println("Error running search menu.", err.Error())
    os.Exit(1)
  }

  return mangas[mangaIndex]
}
