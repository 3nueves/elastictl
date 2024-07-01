package elastictl

import (
	elastictl "elastictl/internal"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var clusterstatusCmd = &cobra.Command{
	Use:     "clusterstatus",
	Aliases: []string{"clup"},
	Short:   "cluster status pretty",
	// Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		es := elastictl.CreateNewUser()

		res, err := es.Client.Cluster.Stats(
			es.Client.Cluster.Stats.WithHuman(),
			es.Client.Cluster.Stats.WithPretty(),
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
	rootCmd.AddCommand(clusterstatusCmd)
}
