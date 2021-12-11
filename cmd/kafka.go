package cmd

import (
	"fmt"
	"github.com/avast/retry-go"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"
	"time"
)

func init() {
	rootCmd.AddCommand(kafkaCmd)
}

var kafkaCmd = &cobra.Command{
	Use:   "kafka KAFKA_URL",
	Short: "Waits for a successful connection to Kafka.",
	Long: `Waits for a successful connection to Kafka.
Example:
	waitfor kafka localhost:9090`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := retry.Do(
			func() error {
				var url = args[0]

				if Verbose {
					fmt.Printf("Connecting to Kafka at %s\n", url)
				}
				conn, err := kafka.Dial("tcp", url)
				if err != nil {
					return err
				}
				defer conn.Close()

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
