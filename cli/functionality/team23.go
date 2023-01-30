/*
This file contains extra functionality for CLI commands. Make sure to 
capitalize first letter of functions to be used in external folders.
*/

package functionality

import (
	"fmt"
	"os"
)


func test(input string) (int){
	return 0
}


func READ_url_file(input []string) (int) {
	filename := os.Args[2]

	contents, err := os.ReadFile(filename)
    if err != nil {
        fmt.Println("File reading error", err)
        return 1
    }
    fmt.Println("Contents of file:", string(contents))
	
	return 0
}


func build() (int) {
	return 0
}