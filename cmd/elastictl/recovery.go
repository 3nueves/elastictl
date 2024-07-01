package elastictl

import (
	elastictl "elastictl/internal"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var indexName string

var recoveryCmd = &cobra.Command{
	Use:     "recovery",
	Aliases: []string{"rc"},
	Short:   "Cluster recovery info",
	// Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		es := elastictl.CreateNewUser()

		res, err := es.Client.Cat.Recovery(
			es.Client.Cat.Recovery.WithS(),
			es.Client.Cat.Recovery.WithH(),
			es.Client.Cat.Recovery.WithPretty(),
			es.Client.Cat.Recovery.WithHuman(),
			es.Client.Cat.Recovery.WithDetailed(true),
			es.Client.Cat.Recovery.WithIndex(indexName),
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

		fmt.Println()
		r := fmt.Sprintln(res)
		fmt.Println(strings.Replace(r, "[200 OK] ", "", 1))
	},
}

func init() {
	recoveryCmd.Flags().StringVarP(&indexName, "search", "s", "", "add index name for search")
	rootCmd.AddCommand(recoveryCmd)
}
