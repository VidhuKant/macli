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
  "strconv"
)

func GetSeasonalAnime(season mg.Season, sort string) []mg.Anime {
	var res []mg.Anime
  err := MALClient.GetSeasonalAnime(&res, &mg.SeasonalParams{
		Year: strconv.Itoa(season.Year),
		Season: season.Name,
		Sort: sort,
		Limit: SearchLength,
		Offset: SearchOffset,
		NSFW: SearchNSFW,
		Fields: []string{"title", "id"},
	})

  if err != nil {
    fmt.Println("MyAnimeList reported error while getting seasonal animes:", err.Error())
    os.Exit(1)
  }

  return res
}
