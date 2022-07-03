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
  m "github.com/MikunoNaka/MAL2Go/v3/manga"
  a "github.com/MikunoNaka/MAL2Go/v3/anime"
)

// because MAL2Go/Anime.SearchAnime won't give us all the data sometimes
func GetAnimeData(animeId int, fields []string) a.Anime {
  data, err := animeClient.GetAnimeById(animeId, fields)
  if err != nil {
    fmt.Println("Error while fetching data about anime:", err)
    os.Exit(1)
  }
  return data
}

// because MAL2Go/Manga.SearchManga won't give us all the data sometimes
func GetMangaData(mangaId int, fields []string) m.Manga {
  data, err := mangaClient.GetMangaById(mangaId, fields)
  if err != nil {
    fmt.Println("Error while fetching data about manga:", err)
    os.Exit(1)
  }
  return data
}
