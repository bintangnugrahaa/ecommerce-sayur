package cmd

import (
	"fmt"
	"order-service/internal/adapter/message"

	"github.com/spf13/cobra"
)

var workerProductSnapCmd = &cobra.Command{
	Use:   "worker-product-snap",
	Short: "Menjalankan worker untuk consume RabbitMQ dan index ke Elasticsearch",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Worker untuk product snapshot indexing sedang berjalan...")
		message.ConsumeFromProduct()
	},
}
