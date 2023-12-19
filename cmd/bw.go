package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewBW returns a new cobra command for BW
func NewBW() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bw",
		Short: "Bitwarden CLI",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running Bitwarden CLI")
		},
	}

	return cmd
}
