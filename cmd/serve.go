package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
    Use:   "serve",
    Short: "Run customer service serve command",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Run customer service serve command called")
    },
}

func init() {
    rootCmd.AddCommand(serveCmd)
}
