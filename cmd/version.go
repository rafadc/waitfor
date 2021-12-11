package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `Probably this won't' change a lot but...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("waitfor v0.1")
	},
}
