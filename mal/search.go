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

func SearchAnime(searchString string, fields []string) []mg.Anime {
  searchLength, searchOffset := SearchLength, SearchOffset
  if AutoSel > 0 {
    searchLength = 1
    searchOffset = AutoSel - 1
  }

	var res []mg.Anime
  err := MALClient.SearchAnime(&res, &mg.SearchParams{
		Limit: searchLength,
		Offset: searchOffset,
		NSFW: SearchNSFW,
		SearchString: searchString,
		Fields: append([]string{"title", "id"}, fields...),
  })

  if err != nil {
    fmt.Println("MyAnimeList reported error while searching:", err.Error())
    os.Exit(1)
  }

  return res
}

func SearchManga(searchString string, fields []string) []mg.Manga {
  searchLength, searchOffset := SearchLength, SearchOffset
  if AutoSel > 0 {
    searchLength = 1
    searchOffset = AutoSel - 1
  }

	var res []mg.Manga
  err := MALClient.SearchManga(&res, &mg.SearchParams{
		Limit: searchLength,
		Offset: searchOffset,
		NSFW: SearchNSFW,
		SearchString: searchString,
		Fields: append([]string{"title", "id"}, fields...),
  })

  if err != nil {
    fmt.Println("MyAnimeList reported error while searching:", err.Error())
    os.Exit(1)
  }

  return res
}
