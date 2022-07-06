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
  // "os"
  a "github.com/MikunoNaka/MAL2Go/v4/anime"
  m "github.com/MikunoNaka/MAL2Go/v4/manga"
)

func DeleteAnime(anime a.Anime) {
  res, err := userAnimeClient.DeleteAnime(anime.Id)
  if err != nil {
    fmt.Println("Error While Deleting " + anime.Title + ":", err)
  }
  if res != "200" {
    fmt.Println("Error: MyAnimeList Returned " + res + " while deleting " + anime.Title)
  }
}

func DeleteManga(manga m.Manga) {
  res, err := userMangaClient.DeleteManga(manga.Id)
  if err != nil {
    fmt.Println("Error While Deleting " + manga.Title + ":", err)
  }
  if res != "200" {
    fmt.Println("Error: MyAnimeList Returned " + res + " while deleting " + manga.Title)
  }
}
