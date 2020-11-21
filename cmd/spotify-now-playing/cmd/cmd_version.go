package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version string

func runVersion(cmd *cobra.Command, args []string) error {
	fmt.Printf("v%s\n", Version)
	return nil
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Long:  "Show version.",
	RunE:  runVersion,
}

func init() { //nolint:gochecknoinits
	rootCmd.AddCommand(versionCmd)
}
