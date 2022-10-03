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
  "github.com/spf13/viper"
  "errors"
)

var clientSuffix string = "-client-id"

func getClientId() (string, error) {
  var id string
  var err error


  if NoSysKeyring {
    id = viper.GetString("auth.client_id")
    if id == "" {
      err = errors.New("secret not found in keyring")
    }
  } else {
    id, err = keyring.Get(serviceName + clientSuffix, userName)
  }

  return id, err
}

func setClientId(clientId string) {
  if NoSysKeyring {
    defer viper.WriteConfig()
    viper.Set("auth.client_id", clientId)
    return
  }

  err := keyring.Set(serviceName + clientSuffix, userName, clientId)
  if err != nil {
    fmt.Println("Error while writing Client ID to keychain", err)
    os.Exit(1)
  }
}

func deleteClientId() {
  if NoSysKeyring {
    defer viper.WriteConfig()
    viper.Set("auth.client_id", "")
  }

  err := keyring.Delete(serviceName + clientSuffix, userName)
  // if secret doesnt exist dont show error
  if err != nil && !NoSysKeyring {
    if err.Error() != "secret not found in keyring" {
      fmt.Println("Error while deleting Client ID", err.Error())
      os.Exit(1)
    }
  }
}

// if client id isn't in keyring
// it will ask the user to enter/create one
func askClientId(storeClientId bool) string {
  clientId, err := getClientId()
  if err != nil {
    if err.Error() == "secret not found in keyring" {
      fmt.Println("Looks like you don't have any Client ID saved.")
      fmt.Println("If you don't have a MyAnimeList Client ID, please go to \x1b[34mhttps://myanimelist.net/apiconfig\x1b[0m and create one.")
      fmt.Println("Remember to set the App Redirect Url to \x1b[33mhttp://localhost:8000\x1b[0m. Other details don't matter.")

      // get clientId from user input
      clientId = secretInput("Enter your Client ID: ", "Client ID Can't be blank")
      validateClientId(clientId)
      if storeClientId {
        setClientId(clientId)
      }
    } else {
      fmt.Println("Error while reading Client ID from keychain:", err)
      os.Exit(1)
    }
  }

  return clientId
}

func validateClientId(clientId string) {
  /* I'm not sure if ALL client IDs are 32 characters
   * but that's most likely the case */
  if len(clientId) != 32 {
    fmt.Println("\x1b[33mWarning:\x1b[0m The Client ID you have entered doesn't have 32 characters.")
    fmt.Println("It's not confirmed but MyAnimeList Client IDs have 32 characters. If you think this is a mistake, you can manually verify your Client ID.")
    fmt.Println("macli doesn't have a way to verify a Client ID. If you think you entered it correctly you can move on with the login process. If you have problems do consider re-entering the Client ID.")

    if confirmInput("Show entered Client ID? [Y/n] ", true) {
      fmt.Println("The Client ID you just entered:", clientId)
      if !confirmInput("Is this correct? [Y/n] ", true) {
        fmt.Println("Please verify your Client ID and run \x1b[33m`macli login`\x1b[0m again.")
        os.Exit(1)
      }
    }
  }
}
