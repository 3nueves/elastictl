package elastictl

import (
	elastictl "elastictl/internal"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var nameFile string
var templateCmd = &cobra.Command{
	Use:     "template",
	Aliases: []string{"tem"},
	Short:   "list templates",
	// Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		es := elastictl.CreateNewUser()

		res, err := es.Client.Cat.Templates(
			es.Client.Cat.Templates.WithHuman(),
			es.Client.Cat.Templates.WithV(true),
			es.Client.Cat.Templates.WithName(nameFile),
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
	templateCmd.Flags().StringVarP(&nameFile, "search", "s", "", "Search template")
	rootCmd.AddCommand(templateCmd)
}
