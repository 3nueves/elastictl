package elastictl

import (
	elastictl "elastictl/internal"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var nameIndice string
var indicesCmd = &cobra.Command{
	Use:     "indices",
	Aliases: []string{"idx"},
	Short:   "list indices",
	// Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		es := elastictl.CreateNewUser()

		res, err := es.Client.Cat.Indices(
			es.Client.Cat.Indices.WithHuman(),
			es.Client.Cat.Indices.WithV(true),
			es.Client.Cat.Indices.WithPretty(),
			es.Client.Cat.Indices.WithH(),
			es.Client.Cat.Indices.WithIndex(nameIndice),
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
	indicesCmd.Flags().StringVarP(&nameIndice, "search", "s", "", "Search indices")
	rootCmd.AddCommand(indicesCmd)
}
