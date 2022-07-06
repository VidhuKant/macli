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
  "os/user"
  "os/exec"
  "runtime"
  "errors"
  "fmt"
  "github.com/zalando/go-keyring"
)

var serviceName string = "macli"
var userName string

func init() {
  // get user and set username
  currentUser, err := user.Current()
  if err != nil {
    fmt.Println("Error getting current user's info", err.Error())
    os.Exit(1)
  }
  userName = currentUser.Username
}

// asks for all the details
func Login() {
  /* check if an auth token already exists
   * if there is an error with keyring, askClientId would handle it
   * can safely ignore error here */
  existingToken, _ := keyring.Get(serviceName, userName)
  if existingToken != "" {
    if !confirmInput("Already logged in. Log in again? [Y/n] ", true) {
      fmt.Println("Login process aborted")
      os.Exit(0)
    }
  }

  clientId := askClientId()
  challenge := codeChallenge()
  link := generateLink(clientId, challenge)

  openInBrowser(link)
  listen(clientId, challenge)
}

func generateLink(clientId, challenge string) string {
  return "https://myanimelist.net/v1/oauth2/authorize?response_type=code&client_id=" + clientId + "&code_challenge=" + challenge
}

func openInBrowser(url string) {
  fmt.Println("Attempting to launch \033[36m" + url + "\033[0m in your default web browser. If it doesn't launch please manually copy-paste the link.")

  var err error
  switch runtime.GOOS {
  case "linux":
    err = exec.Command("xdg-open", url).Start()
  case "windows":
    err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
  case "darwin":
    err = exec.Command("open", url).Start()
  default:
    err = errors.New("<failed to detect platform>")
  }

  if err != nil {
    fmt.Println("There was an error while launching your browser.")
    fmt.Println("Please manually copy and paste the above URL into your web browser.")
    fmt.Println(err)
  }
}

func Logout() {
  existingToken, _ := keyring.Get(serviceName, userName)
  deleteToken()
  deleteExpiresIn()
  deleteRefreshToken()
  if existingToken != "" {
    fmt.Println("Deleted user credentials.")
  }

  // only ask to delete Client ID if it actually exists
  existingClientId, _ := getClientId()
  if existingClientId != "" && confirmInput("Delete your Client ID? [y/N] ", false) {
    deleteClientId()
    fmt.Println("Deleted Client ID.")
  }
}
