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
  a "github.com/MikunoNaka/MAL2Go/v3/user/anime"
  m "github.com/MikunoNaka/MAL2Go/v3/user/manga"
  "github.com/MikunoNaka/macli/util"
)

func SetEpisodes(animeId, prevValue int, ep string) a.UpdateResponse {
  res, err := userAnimeClient.SetWatchedEpisodes(animeId, util.ParseNumeric(ep, prevValue))
  if err != nil {
    fmt.Println("MyAnimeList returned error while updating episodes:", err)
    os.Exit(1)
  }
  return res
}

func SetChapters(mangaId, prevValue int, ch string) m.UpdateResponse {
  res, err := userMangaClient.SetChaptersRead(mangaId, util.ParseNumeric(ch, prevValue))
  if err != nil {
    fmt.Println("MyAnimeList returned error while updating chapters:", err)
    os.Exit(1)
  }
  return res
}
