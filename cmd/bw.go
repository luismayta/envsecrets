package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/luismayta/envsecrets/v1/clients/bitwarden"
	"github.com/luismayta/envsecrets/v1/internal/app/common"
	"github.com/luismayta/envsecrets/v1/internal/errors"
)

// NewBW returns a new cobra command for BW
func NewBW() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bw",
		Short: "Bitwarden CLI",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			err := beforeChecksBW(cmd, args)
			errors.Must(err, errors.ErrorReadConfig, "Error in beforeChecks")
		},
		Run: func(cmd *cobra.Command, args []string) {
			bw := bitwarden.NewClient()
			err := bw.SetFoldersIDs(args)
			errors.Must(err, errors.ErrorReadConfig, "Error in SetFoldersIDs")

			err = bw.FetchItems()
			errors.Must(err, errors.ErrorUnknown, "Error in FetchItems")

			common.OutputEnv(bw)
		},
	}

	return cmd
}

func beforeChecksBW(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New(errors.ErrorReadConfig, "insufficient arguments")
	}

	if os.Getenv("BW_SESSION") == "" {
		return errors.New(errors.ErrorReadConfig, "BW_SESSION environment variable not set")
	}

	return nil
}
