/*
Creation of test command. Details on what is executed is described in 'Long'
field of command variable.
*/

package commands

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	//"github.com/19chonm/461_1_23/api"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "runs test suite",
	Long: `Runs test suite located in -- exits 0 if everything is working.
	The minimum requirement for this test suite is that it contain at 
	least 20 distinct test cases and achieve at least 80'%' code coverage 
	as measured by line coverage. The output from this invocation should be 
	a line written to stdout of the form: “X/Y test cases passed. Z% line 
	coverage achieved.” 
	Should exit 0 on success.`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		//packagesToTest := []string{"api/", "worker/"}

		app := "go"
		testArgs := []string{"test", "./...", "-cover", "-v"}

		exec_output := exec.Command(app, testArgs...)
		stdout, err := exec_output.CombinedOutput()

		// distinction between "PASS" and "--- PASS" because if all tests cases pass,
		// "PASS" is printed out again
		output := string(stdout)
		testsPassed := strings.Count(output, "--- PASS")
		testsRan := strings.Count(output, "=== RUN")

		r := regexp.MustCompile(`coverage:\s*\d+\.\d+%\sof\sstatements`)
		matches := r.FindStringSubmatch(output)
		coverage := matches[0]

		// covIdx := strings.Index(output, "coverage: ")
		// statIdx := strings.Index(output, " of statements")
		// subStr := output[covIdx:statIdx]
		// fmt.Println(subStr)
		// begIdx := api.GetNthOccurance(subStr, 20, 1)
		// endIdx:= api.GetNthOccurance(subStr, 20, 2)
		// cove := subStr[begIdx+1:endIdx]
		// //coverage := output[numsIdx+1:endIdx]

		if err != nil || testsPassed > testsRan {
			fmt.Println("CLI: ", err.Error())
			fmt.Println("SHIT ahppeingds")
			os.Exit(1)
		}

		fmt.Printf("%d/%d test cases passed. %s line coverage achieved\n", testsPassed, testsRan, coverage)
		os.Exit(0)

	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
