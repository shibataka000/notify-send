// Package main provides a CLI tool to send desktop notifications.
package main

import (
	"errors"
	"os"

	"github.com/gen2brain/beeep"
	"github.com/spf13/cobra"
)

func main() {
	var appName string

	rootCmd := &cobra.Command{
		Use:   "notify-send",
		Short: "Send a desktop notification",
		RunE: func(_ *cobra.Command, args []string) error {
			beeep.AppName = appName

			if len(args) == 0 {
				return errors.New("No summary specified.")
			}
			summary := args[0]

			body := ""
			if len(args) > 1 {
				body = args[1]
			}

			return beeep.Notify(summary, body, "")
		},
	}

	rootCmd.Flags().StringVarP(&appName, "app-name", "a", "notify-send", "Specifies the app name for the notification.")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
