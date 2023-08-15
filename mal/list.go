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
  "os"
)

// TODO: return all the list items using loop
func AnimeList(user, status, sort string) []mg.Anime {
	var res []mg.Anime
  _, err := MALClient.GetAnimeList(&res, &mg.ListParams{
		Username: user,
		Status: status,
		Sort: sort,
		Limit: 1000,
		Offset: 0,
		NSFW: SearchNSFW,
		Fields: []string{"title", "num_episodes", "media_type"},
	})

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  return res
}

// TODO: return all the list items using loop
func MangaList(user, status, sort string) []m.Manga {
	var res []mg.Manga
  _, err := MALClient.GetMangaList(&res, &mg.ListParams{
		Username: user,
		Status: status,
		Sort: sort,
		Limit: 1000,
		Offset: 0,
		NSFW: SearchNSFW,
		Fields: []string{"title", "num_chapters", "num_volumes", "media_type"},
	})

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  return res
}
