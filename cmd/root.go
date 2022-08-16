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
	"os"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	queryOnlyMode , mangaMode bool
	entryId int
)

var rootCmd = &cobra.Command{
	Use:   "macli",
	Short: "macli - Unofficial CLI-Based MyAnimeList Client.",
	Long: "macli is an unofficial MyAnimeList Client for use inside the terminal.",
}

func init() {
	viper.SetConfigName("macli")
	viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    viper.AddConfigPath("$HOME/.config/macli")
    viper.AddConfigPath("/etc/macli")

	// dont show error if file not found
	// macli doesnt need a config file to work properly
    if err := viper.ReadInConfig(); err != nil {
		// error if config file found but has errors
	    if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Println("Error while reading macli config file:", err)
			os.Exit(1)
	    }
	}

	authConfig := viper.Get("auth").(map[string]interface{})
	defConfig := viper.Get("defaults").(map[string]interface{})

	fmt.Println("Config File Contents:", authConfig["token"], defConfig["prompt_length"])
	os.Exit(0)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
