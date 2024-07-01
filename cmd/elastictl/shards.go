package elastictl

import (
	elastictl "elastictl/internal"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var shardsCmd = &cobra.Command{
	Use:     "shards",
	Aliases: []string{"sh"},
	Short:   "list shards and states",
	// Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		es := elastictl.CreateNewUser()

		res, err := es.Client.Cat.Shards(
			es.Client.Cat.Shards.WithPretty(),
			es.Client.Cat.Shards.WithV(true),
			es.Client.Cat.Shards.WithHuman(),
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
	rootCmd.AddCommand(shardsCmd)
}
