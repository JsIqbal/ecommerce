package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd of the binary
var (
	RootCmd = &cobra.Command{
		Use:   "backend-search",
		Short: "backend search server binary",
	}
)

func init() {
	RootCmd.AddCommand(serveRestCmd)
	RootCmd.AddCommand(seederCmd)
}

// Execute executes the root command
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
