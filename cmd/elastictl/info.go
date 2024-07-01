package elastictl

import (
	elastictl "elastictl/internal"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:     "info",
	Aliases: []string{"in"},
	Short:   "Info elasticsearch",
	// Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		es := elastictl.CreateNewUser()

		res, err := es.Client.Info()

		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}

		defer res.Body.Close()

		r := fmt.Sprintln(res)
		fmt.Println(strings.Replace(r, "[200 OK] ", "", 1))

	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
