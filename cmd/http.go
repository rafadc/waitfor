package cmd

import (
	"fmt"
	"github.com/avast/retry-go"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"time"
)

func init() {
	rootCmd.AddCommand(httpCmd)
}

var httpCmd = &cobra.Command{
	Use:   "http URL",
	Short: "Waits for an HTTP call to be successful",
	Long:  `Waits for an HTTP call to be successful. It can wait for a given status or just any response.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		err := retry.Do(
			func() error {
				var url = args[0]

				if Verbose {
					fmt.Printf("Making HTTP request to %s\n", url)
				}
				resp, err := http.Get(url)
				if err != nil {
					return err
				}
				defer resp.Body.Close()
				_, err = ioutil.ReadAll(resp.Body)
				if err != nil {
					return err
				}

				return nil
			},
			retry.Attempts(Retries),
			retry.Delay(Delay*time.Second),
		)
		if err != nil {
			fmt.Printf("Error %v\n", err)
		}

	},
}
