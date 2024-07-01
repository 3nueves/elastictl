package elastictl

import (
	elastictl "elastictl/internal"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var shardsnodesCmd = &cobra.Command{
	Use:     "shardsnodes",
	Aliases: []string{"sn"},
	Short:   "Distribution of shards by nodes",
	// Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		es := elastictl.CreateNewUser()

		res, err := es.Client.Cat.Allocation(
			es.Client.Cat.Allocation.WithPretty(),
			es.Client.Cat.Allocation.WithHuman(),
			es.Client.Cat.Allocation.WithV(true),
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
	rootCmd.AddCommand(shardsnodesCmd)
}
