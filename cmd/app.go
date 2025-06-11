package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func Execute() error {
    rootCmd := &cobra.Command{
        Use:   "app",
        Short: "Main application CLI",
        Long:  "Main application CLI",
    }

    rootCmd.AddCommand(service)

    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
    return nil
}
