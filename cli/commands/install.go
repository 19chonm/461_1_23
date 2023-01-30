/*
Creation of install command. Details on what is executed is described in 'Long'
field of command variable.
*/

package commands

import (
	"fmt"
	"os/exec"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use: "install",
	Short: "installs dependencies to use app",
	Long: "as of now, only installs cobra package which is required to use our Go app",
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args[]string){
		
		app := "go"
		arg0 := "get"
		arg1 := "-u"
		arg2 := "github.com/spf13/cobra@latest"

		exec_output := exec.Command(app, arg0, arg1, arg2)
		stdout, err := exec_output.Output()
		
		if err != nil {
			fmt.Println(err.Error())
		} else{
			fmt.Println("Installation succesful", string(stdout))
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}