package elastictl

import (
	elastictl "elastictl/internal"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var nameQueue string
var threadpoolCmd = &cobra.Command{
	Use:     "threadpool",
	Aliases: []string{"thpool"},
	Short:   "cluster thread pool",
	// Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		es := elastictl.CreateNewUser()

		res, err := es.Client.Cat.ThreadPool(
			es.Client.Cat.ThreadPool.WithPretty(),
			es.Client.Cat.ThreadPool.WithHuman(),
			es.Client.Cat.ThreadPool.WithThreadPoolPatterns(nameQueue),
		)
		if err != nil {
			log.Printf("Error getting response from Elasticsearch: %s", err)
			return
		}

		defer res.Body.Close()

		if res.IsError() {
			log.Printf("HTTP response error: %s", res.Status())
			return
		}

		r := fmt.Sprintln(res)
		fmt.Println(strings.Replace(r, "[200 OK] ", "", 1))

	},
}

func init() {
	threadpoolCmd.Flags().StringVarP(&nameQueue, "search", "s", "", "Search Queues")
	rootCmd.AddCommand(threadpoolCmd)
}
