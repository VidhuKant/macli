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
  a "github.com/MikunoNaka/MAL2Go/v2/anime"
  m "github.com/MikunoNaka/MAL2Go/v2/manga"
  u "github.com/MikunoNaka/MAL2Go/v2/user"
  ua "github.com/MikunoNaka/MAL2Go/v2/user/anime"
  um "github.com/MikunoNaka/MAL2Go/v2/user/manga"
)

var animeClient a.Client
var mangaClient m.Client
var userClient u.Client
var userAnimeClient ua.Client
var userMangaClient um.Client

// init() would kill the program prematurely on `macli login` command
func Init() {
  secret := auth.GetToken()
  tk := "Bearer " + secret

  // initialise MAL2Go Client(s)
  animeClient.AuthToken = tk
  mangaClient.AuthToken = tk
  userClient.AuthToken = tk
  userAnimeClient.AuthToken = tk
  userMangaClient.AuthToken = tk
}
