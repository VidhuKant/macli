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

You should have received a copy of the GNU General Public Lice
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package auth

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"github.com/zalando/go-keyring"
	"os"
	"os/exec"
	"os/user"
	"runtime"
)

// set with ldflags, allows for direct login without setting up client id
var VendoredClientId string
var serviceName string = "macli"
var userName string

func init() {
	fmt.Println(VendoredClientId)
	// get user and set username
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error getting current user's info", err.Error())
		os.Exit(1)
	}
	userName = currentUser.Username
}

// asks for all the details
func Login(tk, clientId string, storeClientId bool) {
	// check if an auth token already exists
	var existingToken string
	if NoSysKeyring {
		existingToken = viper.GetString("auth.token")
	} else {
		/* if there is an error with keyring, askClientId would handle it
		 * can safely ignore error here */
		existingToken, _ = keyring.Get(serviceName, userName)
	}

	if existingToken != "" {
		if !confirmInput("Already logged in. Log in again? [Y/n] ", true) {
			fmt.Println("Login process aborted")
			os.Exit(0)
		}
	}

	if clientId == "" {
		if VendoredClientId == "" {
			clientId = askClientId(storeClientId)
		} else {
			clientId = VendoredClientId
		}
	} else {
		validateClientId(clientId)
		if storeClientId {
			setClientId(clientId)
		}
	}

	if tk != "" {
		setToken(tk)
		fmt.Println("\x1b[32mYou have successfully logged into macli.\x1b[0m")
		fmt.Println("\x1b[32mYou can close the web browser tab now.\x1b[0m")
		return
	}

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
