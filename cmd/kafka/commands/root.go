package commands

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use:   "myclub",
		Short: "test",
	}
)

func Run() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(cmdAddPlayer)
}
