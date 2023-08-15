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
	"vidhukant.com/mg"
  "fmt"
)

func SetAnimeScore(animeId, score int) mg.AnimeUpdateResponse {
	var res mg.AnimeUpdateResponse
  err := MALClient.UpdateAnime(&res, animeId, map[string]interface{}{mg.Score: score})
  if err != nil {
    fmt.Println("MyAnimeList returned error while updating anime score:", err)
  }
  return res
}

func SetMangaScore(mangaId, score int) mg.MangaUpdateResponse {
	var res mg.MangaUpdateResponse
  err := MALClient.UpdateManga(&res, mangaId, map[string]interface{}{mg.Score: score})
  if err != nil {
    fmt.Println("MyAnimeList returned error while updating manga score:", err)
  }
  return res
}
