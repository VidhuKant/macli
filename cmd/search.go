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

package cmd

import (
	"fmt"

	"errors"
	"github.com/spf13/cobra"
	"log"
    p "github.com/manifoldco/promptui"

	"github.com/MikunoNaka/macli/ui"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for an anime.",
	Long: `
-- help/description to be added later
`,
	Run: func(cmd *cobra.Command, args []string) {
        validate := func(input string) error {
			if input == "" {
				return errors.New("Search can't be blank")
			}

			return nil
	    }

		prompt := p.Prompt {
		    Label: "Search Anime: ",
			Validate: validate,
		}

		res, err := prompt.Run()
		if err != nil {
			log.Fatal("Failed to run prompt.", err)
		}

		fmt.Println(ui.SearchAndGetID("Select Anime", res))
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
