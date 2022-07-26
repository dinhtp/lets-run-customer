package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("grpc called")
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// grpcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// grpcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
