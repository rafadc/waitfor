package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var Verbose bool
var Retries uint
var Delay time.Duration

var rootCmd = &cobra.Command{
	Use:   "waitfor",
	Short: "waitfor is a command line that makes waiting for services easier",
	Long: `waitfor is a command line that makes waiting for services easier
			for more info visit http://github.com/rafadc/waitfor`,
}

func Execute() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().UintVarP(&Retries, "retries", "r", 5, "number of retries")
	rootCmd.PersistentFlags().DurationVarP(&Delay, "delay", "d", 5, "delay in seconds between retries")
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
