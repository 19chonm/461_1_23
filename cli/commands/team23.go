/*
Root command is created as team23. Holds no functionality but all other
commands are built on top of this command. Creation of new commands 
requires an init function per command with rootCmd.AddCommand(<newCmd>)
*/

package commands

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use: "team23",
	Short: "team23 - root command for app",
	Long: "team23 is the root command to navigate through Team 23's CLI",
	Run: func(cmd *cobra.Command, args[]string){

	},
}


func Execute() {
	fmt.Println("Start execution")
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error using CLI '%s'", err)
		os.Exit(1)
	}
	os.Exit(0)
}