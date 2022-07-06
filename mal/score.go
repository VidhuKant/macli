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
  a "github.com/MikunoNaka/MAL2Go/v4/user/anime"
  m "github.com/MikunoNaka/MAL2Go/v4/user/manga"
)

func SetAnimeScore(animeId, score int) a.UpdateResponse {
  resp, err := userAnimeClient.SetScore(animeId, score)
  if err != nil {
    fmt.Println("MyAnimeList returned error while updating anime score:", err)
  }
  return resp
}

func SetMangaScore(mangaId, score int) m.UpdateResponse {
  resp, err := userMangaClient.SetScore(mangaId, score)
  if err != nil {
    fmt.Println("MyAnimeList returned error while updating manga score:", err)
  }
  return resp
}
