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
  "github.com/spf13/viper"
  a "github.com/MikunoNaka/MAL2Go/v4/anime"
  m "github.com/MikunoNaka/MAL2Go/v4/manga"
  u "github.com/MikunoNaka/MAL2Go/v4/user"
  ua "github.com/MikunoNaka/MAL2Go/v4/user/anime"
  um "github.com/MikunoNaka/MAL2Go/v4/user/manga"
)

var (
  Secret string
  animeClient a.Client
  mangaClient m.Client
  userClient u.Client
  userAnimeClient ua.Client
  userMangaClient um.Client

  SearchLength, SearchOffset int
  SearchNSFW bool
)

// init() would kill the program prematurely on `macli login` command
func Init() {
  if Secret == "" {
    Secret = auth.GetToken()
  }
  tk := "Bearer " + Secret

  /* NOTE: currently, macli is checking wether the specified
   * search length, etc is the default value (5) or not. if it is not
   * then it wont do anything. if it is, then if a config file
   * exists the value in the config file will be used
   * this works but flags won't be able to take precedence
   *
   * i.e if the value in config file is 6 but I want to set it to 5 through
   * flags, it will see that the value is the default value so it'll use
   * the value in the macli.yaml file which is 6. in this case the
   * flags aren't taking precedence. fix that! */
  // load config file vars (if any)
  confSearchLength := viper.Get("searching.search_length")
  confSearchOffset := viper.Get("searching.search_offset")
  confSearchNsfw   := viper.Get("searching.search_nsfw")
  confSecret       := viper.Get("auth.token")

  // if SearchLength is the default value just use the one in config file if any
  if confSearchLength != nil && SearchLength == 10 {
    SearchLength = confSearchLength.(int)
  }
  // if SearchOffset is the default value just use the one in config file if any
  if confSearchOffset != nil && SearchOffset == 0 {
    SearchOffset = confSearchOffset.(int)
  }
  // if SearchNsfw is the default value just use the one in config file if any
  if confSearchNsfw != nil && SearchNSFW == false {
    SearchNSFW = confSearchNsfw.(bool)
  }

  /* the secret stored in the config file
   * takes precedence on the system keyring */
  if confSecret != nil && confSecret != "" {
    Secret = confSecret.(string)
  }

  // initialise MAL2Go Client(s)
  animeClient.AuthToken = tk
  mangaClient.AuthToken = tk
  userClient.AuthToken = tk
  userAnimeClient.AuthToken = tk
  userMangaClient.AuthToken = tk
}
