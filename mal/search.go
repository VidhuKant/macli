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

package mal

import (
  "fmt"
  "os"
  a "github.com/MikunoNaka/MAL2Go/v3/anime"
  m "github.com/MikunoNaka/MAL2Go/v3/manga"
)

func SearchAnime(searchString string) []a.Anime {
  fields := []string{"title", "id", "my_list_status"}

  res, err := animeClient.SearchAnime(searchString, SearchLength, SearchOffset, SearchNSFW, fields)
  if err != nil {
    fmt.Println("MyAnimeList reported error while searching:", err.Error())
    os.Exit(1)
  }

  return res
}

func SearchManga(searchString string) []m.Manga {
  fields := []string{"title", "id", "my_list_status"}

  res, err := mangaClient.SearchManga(searchString, SearchLength, SearchOffset, SearchNSFW, fields)
  if err != nil {
    fmt.Println("MyAnimeList reported error while searching:", err.Error())
    os.Exit(1)
  }

  return res
}
