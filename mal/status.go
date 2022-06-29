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
)

func SetAnimeStatus(animeId int, status string) {
  resp, err := userAnimeClient.SetStatus(animeId, status)
  if err != nil {
    fmt.Println("Error while parsing status:", err.Error())
    os.Exit(1)
  }
  // TODO: do something with resp
  fmt.Println(resp)

  // not needed with MAL2Go v2.... probably
  // if resp.Error != "" {
  //   fmt.Println("MyAnimeList reported error on setting anime status", resp.Error, resp.Message)
  //   os.Exit(1)
  // }
}

func SetMangaStatus(mangaId int, status string) {
  resp, err := userMangaClient.SetStatus(mangaId, status)
  if err != nil {
    fmt.Println("Error while parsing status:", err.Error())
    os.Exit(1)
  }
  // TODO: do something with resp
  fmt.Println(resp)

  // not needed with MAL2Go v2.... probably
  // if resp.Error != "" {
  //   fmt.Println("MyAnimeList reported error on setting manga status", resp.Error, resp.Message)
  //   os.Exit(1)
  // }
}
