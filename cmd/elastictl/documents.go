package elastictl

import (
	elastictl "elastictl/internal"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var nameIndex string
var documentsCmd = &cobra.Command{
	Use:     "documents",
	Aliases: []string{"doc"},
	Short:   "list documents",
	// Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		es := elastictl.CreateNewUser()

		res, err := es.Client.Search(
			es.Client.Search.WithHuman(),
			es.Client.Search.WithPretty(),
			es.Client.Search.WithBatchedReduceSize(5),
			es.Client.Search.WithQuery(nameIndex),
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
	documentsCmd.Flags().StringVarP(&nameIndex, "search", "s", "", "Search documents")
	rootCmd.AddCommand(documentsCmd)
}
