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

package auth

import (
  "os"
  "fmt"
  "github.com/zalando/go-keyring"
)

func GetToken() string {
  secret, err := keyring.Get(serviceName, userName)
  if err != nil {
    fmt.Println("\x1b[31mError while reading access token from keychain:", err.Error(), "\x1b[0m")
    fmt.Println("Run `macli login` first to authenticate with your MyAnimeList API Token")
    os.Exit(1)
  }

  return secret
}

func setToken(secret string) {
  err := keyring.Set(serviceName, userName, secret)
  if err != nil {
    fmt.Println("Error while writing access token to keychain", err)
    os.Exit(1)
  }
}

func deleteToken() {
  err := keyring.Delete(serviceName, userName)
  // TODO: if secret doesnt exist dont show error
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
