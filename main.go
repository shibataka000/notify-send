package main

import (
	"os"
	"strings"

	"github.com/gen2brain/beeep"
	"github.com/spf13/cobra"
)

func main() {
	var appName string

	rootCmd := &cobra.Command{
		Use:   "notify-send",
		Short: "Send a desktop notification",
		RunE: func(cmd *cobra.Command, args []string) error {
			return beeep.Notify(appName, strings.Join(args, " "), "")
		},
	}

	rootCmd.Flags().StringVarP(&appName, "app-name", "a", "", "Specifies the app name for the notification.")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
