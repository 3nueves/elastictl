package elastictl

import (
	elastictl "elastictl/internal"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:     "status",
	Aliases: []string{"sta"},
	Short:   "cluster status",
	// Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		es := elastictl.CreateNewUser()

		res, err := es.Client.Cluster.Health(
			es.Client.Cluster.Health.WithPretty(),
			es.Client.Cluster.Health.WithHuman(),
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
	rootCmd.AddCommand(statusCmd)
}
