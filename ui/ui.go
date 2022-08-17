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

package ui

import (
    "github.com/spf13/viper"
)

var PromptLength int

/* NOTE: currently, macli is checking wether the specified
 * prompt length is the default value (5) or not. if it is not
 * then it wont do anything. if it is, then if a config file
 * exists the value in the config file will be used
 * this works but flags won't be able to take precedence
 *
 * i.e if the value in config file is 6 but I want to set it to 5 through
 * flags, it will see that the value is the default value so it'll use
 * the value in the macli.yaml file which is 6. in this case the
 * flags aren't taking precedence. fix that! */
func init() {
	// read prompt length from config file
	confPromptLength := viper.Get("searching.prompt_length")

    // if PromptLength is the default value just use the one in config file if any
    if confPromptLength != nil && PromptLength == 5 {
      PromptLength = confPromptLength.(int)
    }
}
