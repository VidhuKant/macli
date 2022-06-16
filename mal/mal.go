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
  "github.com/MikunoNaka/macli/auth"
  a "github.com/MikunoNaka/MAL2Go/anime"
  m "github.com/MikunoNaka/MAL2Go/manga"
  u "github.com/MikunoNaka/MAL2Go/user"
  ua "github.com/MikunoNaka/MAL2Go/user/anime"
  um "github.com/MikunoNaka/MAL2Go/user/manga"
)

var animeClient a.Client
var mangaClient m.Client
var userClient u.Client
var userAnimeClient ua.Client
var userMangaClient um.Client

func init() {
  secret := auth.GetToken()

  // initialise MAL2Go Client(s)
  animeClient.AuthToken = "Bearer " + secret
  mangaClient.AuthToken = "Bearer " + secret
  userClient.AuthToken = "Bearer " + secret
  userAnimeClient.AuthToken = "Bearer " + secret
  userMangaClient.AuthToken = "Bearer " + secret
}
