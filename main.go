// Package main provides a CLI tool to send desktop notifications.
package main

import (
	"errors"
	"os"

	"github.com/gen2brain/beeep"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "notify-send <summary> [body]",
		Short: "a program to send desktop notifications",
		RunE: func(_ *cobra.Command, args []string) error {
			switch len(args) {
			case 0:
				return errors.New("No summary specified.")
			case 1:
				return beeep.Notify(args[0], "", "")
			default:
				return beeep.Notify(args[0], args[1], "")
			}
		},
		SilenceUsage: true,
	}

	rootCmd.Flags().StringVarP(&beeep.AppName, "app-name", "a", "notify-send", "Specifies the app name for the notification.")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
