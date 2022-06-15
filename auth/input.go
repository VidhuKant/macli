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
  "errors"
  p "github.com/manifoldco/promptui"
)

// because importing macli/ui causes import cycle
func secretInput(label, errMessage string) string {
  validate := func(input string) error {
    if input == "" {
      return errors.New(errMessage)
    }
    return nil
  }

  template := &p.PromptTemplates {
    Valid: "{{ . | cyan }}",
    Invalid: "{{ . | cyan }}",
    Success: "{{ . | blue }}",
  }

  prompt := p.Prompt {
    Label: label,
    Templates: template,
    Validate: validate,
    Mask: '*',
  }

  res, err := prompt.Run()
  if err != nil {
    fmt.Println("Failed to run secret input prompt.", err.Error())
    os.Exit(1)
  }

  return res
}
