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
  "net/http"
  "net/url"
  "encoding/json"
  "os"
  "fmt"
	// "io/ioutil"
)

func listen(clientId, verifier string) {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query()

	code, codePresent := query["code"]
	if !codePresent {
	  // TODO: check if error message present
	  fmt.Println("Error: response from MyAnimeList doesn't contain required code.")
	  os.Exit(1)
	}

	accessToken, refreshToken, expiresIn := requestToken(clientId, verifier, code[0])

	if accessToken != "" {
	  w.WriteHeader(200)
	  w.Write([]byte("<h1>You have successfully logged into macli.</h1>"))
	}

	setToken(accessToken)
	setRefreshToken(refreshToken)
	setExpiresIn(expiresIn)
  })

  err := http.ListenAndServe(":8000", nil)
  if err != nil {
    fmt.Println("There was an error initialising the server", err.Error())
    os.Exit(1)
  }
}

func requestToken(clientId, verifier, code string) (string, string, string) {
  data := url.Values{
	"client_id": {clientId},
	"code_verifier": {verifier},
	"grant_type": {"authorization_code"},
	"code": {code},
  }

  resp, err := http.PostForm("https://myanimelist.net/v1/oauth2/token", data)
  if err != nil {
    fmt.Println("Error while requesting an access token:", err)
	os.Exit(1)
  }

  var res map[string]interface{}
  json.NewDecoder(resp.Body).Decode(&res)

  return fmt.Sprintf("%v", res["access_token"]), fmt.Sprintf("%v", res["refresh_token"]), fmt.Sprintf("%v", res["expires_in"])
}
