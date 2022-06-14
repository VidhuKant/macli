/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strings"
	"github.com/MikunoNaka/macli/ui"
	"github.com/MikunoNaka/macli/mal"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Set an anime/manga's status",
	Long: `
-- help/description to be added later
`,
	Run: func(cmd *cobra.Command, args []string) {
		searchInput := strings.Join(args, " ")

		statusInput, err := cmd.Flags().GetString("status")
		if err != nil {
			fmt.Println("Error while reading status flag.", err.Error())
		}

		mangaMode, err := cmd.Flags().GetBool("manga")
		if err != nil {
			fmt.Println("Error while reading manga flag.", err.Error())
		}

		if mangaMode {
			// setMangaStatus(statusInput, searchInput)
			fmt.Println("Manga mode coming soon")
		} else {
			setAnimeStatus(statusInput, searchInput)
		}

	},
}

func setAnimeStatus(statusInput, searchInput string) {
	if searchInput == "" {
		searchInput = ui.TextInput("Search Anime To Update", "Search can't be blank.")
	}

	anime := ui.AnimeSearch("Select Anime:", searchInput)

	if statusInput == "" {
		ui.StatusMenu(anime)
	} else {
		mal.SetStatus(anime.Id, statusInput)
		fmt.Printf("Successfully set \"%s\" to \"%s\"", anime.Title, statusInput)
	}
}

func init() {
	rootCmd.AddCommand(statusCmd)
    statusCmd.Flags().StringP("status", "s", "", "status to be set")
}
