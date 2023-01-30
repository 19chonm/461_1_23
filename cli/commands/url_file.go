/*
Creation of URL_FILE command. Details on what is executed is described in 'Long'
field of command variable.
*/

package commands

import (
	"os"
	"github.com/spf13/cobra"
	"github.com/19chonm/461_1_23/tree/cli-prototype/cli/functionality"
)

var urlfileCmd = &cobra.Command{
	Use: "URL_FILE",
	Short: "assign scores for packages found in URL_FILE",
	Long: `scans through the urls found within file passed as an argument \n
	and builds a score for the each package based on ramp-up time, correctness, \n
	bus factor, responsiveness from maintainer, and license compatibility. Will \n
	produce a net score and output in NDJSON format to stdout.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args[]string){
		functionality.READ_url_file(os.Args)
	},
}

func init() {
	rootCmd.AddCommand(urlfileCmd)
}