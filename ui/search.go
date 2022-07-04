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

var AnimeSearchFields []string = []string {
  "num_episodes", "synopsis",
  "alternative_titles", "start_date",
  "end_date", "mean", "start_season",
  "rank", "media_type", "status",
  "average_episode_duration",
  "rating", "studios",
}

// only search animes probably only now
func AnimeSearch(label, searchString string) a.Anime {
  animes := mal.SearchAnime(searchString, AnimeSearchFields)

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
{{ "Number of Episodes:" | blue | bold }} {{ if .NumEpisodes }}{{ .NumEpisodes }}{{ else }}{{ "unknown" | faint }}{{ end }}
{{ "English Title:" | blue | bold }} {{ if .AltTitles.En }}{{ .AltTitles.En }}{{ else }}{{ "none" | faint }}{{ end }}
{{ "Japanese Title:" | blue | bold }} {{ if .AltTitles.Ja }}{{ .AltTitles.Ja }}{{ else }}{{ "none" | faint }}{{ end }}
{{ "Original Run:" | blue | bold }} {{ if .StartDate }}{{ .StartDate | cyan }}{{ else }}{{ "unknown" | faint }}{{ end }} - {{ if .EndDate }}{{ .EndDate | yellow }}{{ else }}{{ "unknown" | faint }}{{end}} {{ if .StartSeason.Year }}({{ .StartSeason.Name }} {{ .StartSeason.Year }}){{ else }}{{ end }}
{{ "Mean Score:" | blue | bold }} {{ if .MeanScore }}{{ .MeanScore }}{{ else }}{{ "unknown" | faint }}{{ end }}
{{ "Rank:" | blue | bold }} {{ if .Rank }}{{ .Rank }}{{ else }}{{ "unknown" | faint }}{{ end }}
{{ "Type:" | blue | bold }} {{ .MediaType }}
{{ "Status:" | blue | bold }} {{ .Status }}
{{ "Average Duration:" | blue | bold }} {{ if .DurationSeconds }}{{ .DurationSeconds }} minutes{{ else }}{{ "unknown" | faint }}{{ end }}
{{ "Rating:" | blue | bold }} {{ if .Rating }}{{ .Rating }}{{ else }}{{ "unknown" | faint }}{{ end }}
{{ "Studios:" | blue | bold }} {{ if .Studios }}{{ range .Studios }}{{ .Name }}{{ end }}{{ else }}{{ "unknown" | faint }}{{ end }}
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

var MangaSearchFields []string = []string {
  "num_chapters", "num_volumes", "synopsis",
  "alternative_titles", "start_date",
  "end_date", "mean", "rank",
  "media_type", "status",
}

func MangaSearch(label, searchString string) m.Manga {
  mangas := mal.SearchManga(searchString, MangaSearchFields)

  template := &p.SelectTemplates {
    Label: "{{ . }}",
    Active: "{{ .Title | magenta }}",
    Inactive: "{{ .Title }}",
    Selected: "{{ .Title | blue }}",
    Details: `
--------- {{ .Title }} ----------
{{ "Number of Volumes:" | blue | bold }} {{ .NumVolumes }}
{{ "Number of Chapters:" | blue | bold }} {{ .NumChapters }}
{{ "English Title:" | blue | bold }} {{ .AltTitles.En }}
{{ "Japanese Title:" | blue | bold }} {{ .AltTitles.Ja }}
{{ "Original Run:" | blue | bold }} {{ .StartDate }} - {{ .EndDate }}
{{ "Mean Score:" | blue | bold }} {{ .MeanScore }}
{{ "Rank:" | blue | bold }} {{ .Rank }}
{{ "Type:" | blue | bold }} {{ .MediaType }}
{{ "Status:" | blue | bold }} {{ .Status }}
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
