package elastictl

import (
	"bufio"
	elastictl "elastictl/internal"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// var (
// 	nameIndexDelete []string
// )

var indexDeleteCmd = &cobra.Command{
	Use:     "indexDelete [prefix*]",
	Aliases: []string{"del"},
	Short:   "delete indexes starting by specified prefix",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		es := elastictl.CreateNewUser()

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Do you want remove index?\n\n", args, "\n\n[No/Yes]\n")
		data, err := reader.ReadString('\n')

		if err != nil {
			log.Printf("Error to read variable: %s", err)
			return
		}

		if data == "Yes\n" {
			res, err := es.Client.Indices.Delete(
				args,
				// nameIndexDelete,
				es.Client.Indices.Delete.WithPretty(),
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
		}
	},
}

func init() {
	rootCmd.AddCommand(indexDeleteCmd)

	// indexDeleteCmd.Flags().StringArrayVarP(&nameIndexDelete, "delete", "d", nameIndexDelete, "Delete indices")

	// indexDeleteCmd.Flags().StringVarP(&uuidIndexDelete, "uuid", "i", "", "uuid indices")
}
