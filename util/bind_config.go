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

/* viper.BindPFlag won't work if
 * multiple commands have the same 
 * flags. so this is my hacky
 * way to do that stuff myself
 */

package util

import (
  "github.com/spf13/viper"
  "github.com/spf13/pflag"
)

type SearchConfig struct {
	PromptLength int
	SearchLength int
	SearchOffset int
	AutoSel      int
	SearchNSFW   bool
}

type ListConfig struct {
	ResultsLength int
	ResultsOffset int
	IncludeNSFW   bool
}

// handles prompt-length, search-length, search-offset and search-nsfw
func BindSearchConfig(flags *pflag.FlagSet) (SearchConfig, error) {
	var (
		conf SearchConfig
		err error
	)

	if flags.Lookup("auto-select").Changed {
		conf.AutoSel, err = flags.GetInt("auto-select")
		if err != nil {return conf, err}
	} else {
		conf.AutoSel = viper.GetInt("searching.auto_select_n")
	}

	if flags.Lookup("prompt-length").Changed {
		conf.PromptLength, err = flags.GetInt("prompt-length")
		if err != nil {return conf, err}
	} else {
		conf.PromptLength = viper.GetInt("searching.prompt_length")
	}

	if flags.Lookup("search-length").Changed {
		conf.SearchLength, err = flags.GetInt("search-length")
		if err != nil {return conf, err}
	} else {
		conf.SearchLength = viper.GetInt("searching.search_length")
	}

	if flags.Lookup("search-offset").Changed {
		conf.SearchOffset, err = flags.GetInt("search-offset")
		if err != nil {return conf, err}
	} else {
		conf.SearchOffset = viper.GetInt("searching.search_offset")
	}

	if flags.Lookup("search-nsfw").Changed {
		conf.SearchNSFW, err = flags.GetBool("search-nsfw")
		if err != nil {return conf, err}
	} else {
		conf.SearchNSFW = viper.GetBool("searching.search_nsfw")
	}

	return conf, nil
}

// handles results-length, results-offset, include-nsfw
func BindListConfig(flags *pflag.FlagSet) (ListConfig, error) {
	var (
		conf ListConfig
		err error
	)

	if flags.Lookup("results-length").Changed {
		conf.ResultsLength, err = flags.GetInt("results-length")
		if err != nil {return conf, err}
	} else {
		conf.ResultsLength = viper.GetInt("lists.list_length")
	}

	if flags.Lookup("results-offset").Changed {
		conf.ResultsOffset, err = flags.GetInt("results-offset")
		if err != nil {return conf, err}
	} else {
		conf.ResultsOffset = viper.GetInt("lists.list_offset")
	}

	if flags.Lookup("include-nsfw").Changed {
		conf.IncludeNSFW, err = flags.GetBool("include-nsfw")
		if err != nil {return conf, err}
	} else {
		conf.IncludeNSFW = viper.GetBool("lists.include_nsfw_results")
	}

	return conf, nil
}
