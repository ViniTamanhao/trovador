// Package cmd defines all the possible commands and functionalities for trovador
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.1.0"

var rootCmd = &cobra.Command{
	Use:   "trovador",
	Short: "Media player controller/interface",
	Long: `Trovador - A easy to use utility for media player control. 
		Supports Spotify, browsers, and any MPRIS-compatible player.`,
	Version: version,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Global flags can go here
	// TODO rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}
