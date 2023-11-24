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
	"context"
  "net/http"
  "net/url"
	"html"
  "encoding/json"
  "os"
  "fmt"
	"errors"
)

func listen(clientId, verifier string) {
	mux := http.NewServeMux()
	server := &http.Server{Addr: ":8000", Handler: mux}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		code, codeExists := r.URL.Query()["code"]

		if codeExists {
			err := getToken(clientId, verifier, code[0])
			if err != nil {
			  fmt.Fprintf(w, "<p>An error occoured while logging in: %s</p>", html.EscapeString(err.Error()))

				fmt.Println("Error while requesting an access token from MyAnimeList:", err)
			} else {
				fmt.Fprintf(w, "<p>Login successful! You may close this tab now.</p>")

				fmt.Println("\x1b[32mYou have successfully logged into macli.\x1b[0m")
				fmt.Println("\x1b[32mYou can close the web browser tab now.\x1b[0m")
			}
		} else {
			fmt.Fprintf(w, "<p>An error occoured while logging in (invalid request)</p>")

			fmt.Println("An error occoured while logging in (invalid request)")
		}

		go server.Shutdown(context.Background())
  })

  err := server.ListenAndServe()
  if err != nil && err != http.ErrServerClosed {
    fmt.Println("There was an error initialising the server", err.Error())
    os.Exit(1)
  }
}

func getToken(clientId, verifier, code string) error {
  data := url.Values{
		"client_id": {clientId},
		"code_verifier": {verifier},
		"grant_type": {"authorization_code"},
		"code": {code},
  }

  resp, err := http.PostForm("https://myanimelist.net/v1/oauth2/token", data)
  if err != nil {
		return err
  }

  var res map[string]interface{}
  json.NewDecoder(resp.Body).Decode(&res)

	if res["error"] != nil {
		return errors.New(fmt.Sprintf("%v (%v)", res["message"], res["error"]))
	}

	setToken(fmt.Sprintf("%v", res["access_token"]))
	//setRefreshToken(refreshToken)
	//setExpiresIn(expiresIn)
  return nil
}
