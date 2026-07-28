// Command aac2m4b decrypts aax and aaxc files.
package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	if err := NewRootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}

// NewRootCommand builds the cobra command tree.
func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:          "aac2m4b",
		Short:        "decrypts aax and aaxc files",
		SilenceUsage: true,
		Version:      Version,
	}

	rootCmd.AddCommand(NewVersionCommand())

	return rootCmd
}
