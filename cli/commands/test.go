/*
Creation of test command. Details on what is executed is described in 'Long'
field of command variable.
*/

package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use: "test",
	Short: "runs test suite",
	Long: `runs test suite located in -- exits 0 if everything is working. \n
	The minimum requirement for this test suite is that it contain at \n
	least 20 distinct test cases and achieve at least 80'%' code coverage \n
	as measured by line coverage. The output from this invocation should be \n
	a line written to stdout of the form: “X/Y test cases passed. Z% line \n
	coverage achieved.” \n
	Should exit 0 on success.`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args[]string){
		fmt.Println("test command recognized")
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}