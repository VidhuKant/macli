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
  "strconv"
  a "github.com/MikunoNaka/MAL2Go/v2/user/anime"
  m "github.com/MikunoNaka/MAL2Go/v2/user/manga"
)

func SetEpisodes(animeId, prevValue int, ep string) a.UpdateResponse {
  epInt, err := strconv.Atoi(ep)
  if err != nil {
    fmt.Println("Error while parsing episode input", err)
    os.Exit(1)
  }

  var epValue int
  switch ep[0:1] {
  case "+", "-":
    // works both for increment and decrement
    epValue = prevValue + epInt
  default:
    epValue = epInt
  }

  res, err := userAnimeClient.SetWatchedEpisodes(animeId, epValue)
  if err != nil {
    fmt.Println("MyAnimeList returned error while updating episodes:", err)
    os.Exit(1)
  }
  return res
}

func SetChapters(mangaId, prevValue int, ch string) m.UpdateResponse {
  chInt, err := strconv.Atoi(ch)
  if err != nil {
    fmt.Println("Error while parsing chapter input", err)
    os.Exit(1)
  }

  var chValue int
  switch ch[0:1] {
  case "+", "-":
    // works both for increment and decrement
    chValue = prevValue + chInt
  default:
    chValue = chInt
  }

  res, err := userMangaClient.SetChaptersRead(mangaId, chValue)
  if err != nil {
    fmt.Println("MyAnimeList returned error while updating chapters:", err)
    os.Exit(1)
  }
  return res
}
