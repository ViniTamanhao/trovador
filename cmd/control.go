package cmd

import (
	"fmt"
	"os"

	"trovador/interal/controller"

	"github.com/spf13/cobra"
)

var playPauseCmd = &cobra.Command{
	Use:     "play-pause [player]",
	Aliases: []string{"pp", "toggle"},
	Short:   "Toggle play/pause for a player",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		controlPlayer(args[0], "play-pause")
	},
}

var nextCmd = &cobra.Command{
	Use:     "next [player]",
	Aliases: []string{"n"},
	Short:   "Skip to next track",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		controlPlayer(args[0], "next")
	},
}

var previousCmd = &cobra.Command{
	Use:     "previous [player]",
	Aliases: []string{"prev", "p"},
	Short:   "Go to previous track",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		controlPlayer(args[0], "previous")
	},
}

var stopCmd = &cobra.Command{
	Use:   "stop [player]",
	Short: "Stop playback",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		controlPlayer(args[0], "stop")
	},
}

func init() {
	rootCmd.AddCommand(playPauseCmd)
	rootCmd.AddCommand(nextCmd)
	rootCmd.AddCommand(previousCmd)
	rootCmd.AddCommand(stopCmd)
}

func controlPlayer(playerName, action string) {
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

	switch action {
	case "play-pause":
		err = ctrl.PlayPause(player.ID)
	case "next":
		err = ctrl.Next(player.ID)
	case "previous":
		err = ctrl.Previous(player.ID)
	case "stop":
		err = ctrl.Stop(player.ID)
	default:
		fmt.Fprintf(os.Stderr, "Unknown action: %s\n", action)
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing %s: %v\n", action, err)
		os.Exit(1)
	}

	fmt.Printf("Action %s on player %s concluded.\n", action, player.Name)
}
