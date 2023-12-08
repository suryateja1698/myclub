package commands

import (
	"os"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/spf13/cobra"
)

var (
	Logger        log.Logger
	Topic         string
	BrokerAddress string
	rootCmd       = &cobra.Command{
		Use:   "myclub",
		Short: "test",
	}
)

func Run() error {
	return rootCmd.Execute()
}

func init() {
	logger := log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "timestamp", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)
	Logger = logger
	Topic = "real-madrid"
	BrokerAddress = "localhost:9093"

	rootCmd.AddCommand(cmdAddPlayer)

	level.Info(logger).Log("message", "worker is running")
}
