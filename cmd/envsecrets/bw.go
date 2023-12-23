package envsecrets

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/luismayta/envsecrets/v1/internal/common/generator"
	"github.com/luismayta/envsecrets/v1/internal/config"
	"github.com/luismayta/envsecrets/v1/internal/errors"
	"github.com/luismayta/envsecrets/v1/third_party/bitwarden"
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
			conf := config.Initialize()
			bw := bitwarden.NewClient(conf)
			err := bw.SetFoldersIDs(args)
			errors.Must(err, errors.ErrorReadConfig, "Error in SetFoldersIDs")

			err = bw.FetchItems()
			errors.Must(err, errors.ErrorUnknown, "Error in FetchItems")

			generator.OutputEnv(bw)
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
