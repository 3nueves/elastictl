package elastictl

import (
	elastictl "elastictl/internal"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var errorsCmd = &cobra.Command{
	Use:     "errors",
	Aliases: []string{"er"},
	Short:   "show detailed errors in cluster",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		es := elastictl.CreateNewUser()

		res, err := es.Client.Cluster.AllocationExplain(
			es.Client.Cluster.AllocationExplain.WithPretty(),
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
	rootCmd.AddCommand(errorsCmd)
}
