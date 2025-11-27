package cmd

import (
	"fmt"
	"os"
	"trovador/interal/controller"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
	Aliases: []string{"ls", "l"},
	Short: "List all available media players",
	Long: "List all currently running media players that can be controller",
	Run: runList,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func runList(cmd *cobra.Command, args []string) {
	ctrl, err := controller.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting players: %v\n", err)
		os.Exit(1)
	}

	players, err := ctrl.GetPlayers()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting players: %v\n", err)
		os.Exit(1)
	}

	if len(players) == 0 {
		fmt.Println("No media players found.")
		fmt.Println("Start Spotify, a browser with media, or another MPRIS-compatible player.")
		return
	}

	fmt.Printf("Found %d player(s):\n\n", len(players))
	for i, player := range players {
		fmt.Printf("%d. %s\n", i+1, player.Name)

		status, err := ctrl.GetPlayerStatus(player.ID)
		if err != nil {
			fmt.Printf("   Status: Error - %v\n", err)
			continue
		}

		fmt.Printf("   State:  %s\n", status.State)
		if status.Title != "" {
			fmt.Printf("   Title:  %s\n", status.Title)
		}
		if status.Artist != "" {
			fmt.Printf("   Artist: %s\n", status.Artist)
		}
		if status.Album != "" {
			fmt.Printf("   Album:  %s\n", status.Album)
		}
		fmt.Println()
	}
}
