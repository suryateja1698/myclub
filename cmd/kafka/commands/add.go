package commands

import "github.com/spf13/cobra"

var (
	cmdAddPlayer = &cobra.Command{
		Use:   "add",
		Short: "add players",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
)
