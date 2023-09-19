/*
macli - Unofficial CLI-Based MyAnimeList Client
Copyright © 2022-2023 Vidhu Kant Sharma <vidhukant@vidhukant.com>

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
  "runtime"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	queryOnlyMode, mangaMode bool
	entryId int
)

var rootCmd = &cobra.Command{
	Use: "macli",
  Version: "v1.20.5" + " " + runtime.GOOS + "/" + runtime.GOARCH,
	Short: "macli - Unofficial CLI-Based MyAnimeList Client.",
	Long:
      "macli is an unofficial MyAnimeClient for use inside the terminal.\n" +
      "Run \x1b[33m`macli --help`\x1b[0m for instructions.\n\n" +
      "\x1b[34mmacli  Copyright (C) 2022-2023  Vidhu Kant Sharma <vidhukant@vidhukant.com>\n" +
      "This program comes with ABSOLUTELY NO WARRANTY;\n" +
      "This is free software, and you are welcome to redistribute it\n" +
      "under certain conditions; For details refer to the GNU General Public License.\n" +
      "You should have received a copy of the GNU General Public License\n" +
      "along with this program.  If not, see <https://www.gnu.org/licenses/>.\x1b[0m\n" +
      "\n",
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
  viper.SetConfigName("macli")
  viper.AddConfigPath(".")
  viper.AddConfigPath("$HOME/.config")
  viper.AddConfigPath("/etc")

	// dont show error if file not found
	// macli doesnt need a config file to work properly
  if err := viper.ReadInConfig(); err != nil {
    // error if config file found but has errors

    if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
      // if config file isn't found
      fmt.Println("Error while reading macli config file:", err)
      fmt.Println("Exiting... Please check the macli config file.")
      os.Exit(1)
    }
	}

  viper.SetDefault("searching.prompt_length", 5)
  viper.SetDefault("searching.search_length", 10)
  viper.SetDefault("searching.search_offset", 0)
  viper.SetDefault("searching.search_nsfw", false)

  viper.SetDefault("lists.list_offset", 0)
  viper.SetDefault("lists.list_length", 15)
  viper.SetDefault("lists.include_nsfw_results", false)

  viper.SetDefault("auth.save_client_id", "yes")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
