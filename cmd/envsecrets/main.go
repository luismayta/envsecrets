package envsecrets

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/luismayta/envsecrets/v1/internal/config"
	"github.com/luismayta/envsecrets/v1/internal/errors"
)

var conf = config.Initialize()

// Flags names
const (
	logLevelFlagName = "logLevel"
)

var rootCmd = &cobra.Command{
	Use:     conf.App.Name,
	Short:   conf.App.Description,
	Version: conf.App.Version,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		configureLogLevel(cmd)
	},
}

func configureLogLevel(cmd *cobra.Command) {
	logLevelStr := cmd.Flag(logLevelFlagName).Value.String()
	logLevel, err := log.ParseLevel(logLevelStr)
	if err != nil {
		errors.Must(err, errors.ErrorReadConfig, logLevelStr)
	}
	log.SetLevel(logLevel)
}

func init() {
	cobra.OnInitialize()
	persistentFlags := rootCmd.PersistentFlags()
	persistentFlags.String(
		logLevelFlagName,
		log.InfoLevel.String(),
		"Set log level (debug|info|warn|error)",
	)
	rootCmd.AddCommand(NewBW())
}

func Execute() error {
	return rootCmd.Execute()
}
