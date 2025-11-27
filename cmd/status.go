package cmd

import (
	"fmt"
	"os"
	"trovador/interal/controller"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use: "status [player]",
	Short: "Show status of specific player",
	Long: "Display current playback status, track info, and metadata for player.",
	Run: runStatus,
}

func init() {
	rootCmd.AddCommand(statusCmd)
}


func runStatus(cmd *cobra.Command, args []string) {
	playerName := args[0]

	ctrl, err := controller.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	

	player := findPlayer(ctrl, playerName)
	if player == nil {
		fmt.Fprintf(os.Stderr, "Player '%s' not found. Use 'trovador list' to see available players.\n", playerName)
		os.Exit(1)
	}

	status, err := ctrl.GetPlayerStatus(player.ID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting status: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Player: %s\n", player.Name)
	fmt.Printf("	State:  %s\n", status.State)
	if status.Title != "" {
		fmt.Printf("	Title:  %s\n", status.Title)
	}
	if status.Artist != "" {
		fmt.Printf("	Artist: %s\n", status.Artist)
	}
	if status.Album != "" {
		fmt.Printf("	Album:  %s\n", status.Album)
	}

	if status.Duration > 0 {
		fmt.Printf("	Progress: %s\n", controller.FormatProgress(status.Position, status.Duration))
		fmt.Printf("        %s\n", controller.ProgressBar(status.Position, status.Duration, 40))
	} else if status.Position > 0 {
		fmt.Printf("	Position: %s\n", controller.FormatDuration(status.Position))
	}
}
