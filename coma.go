package main

//this package executes bash commands, shows their output and waits for a
//key press to exeute the next command or does it automatically with a time
//delay. It is aimed at providing instructional content for teachers and
//hackers who wish to demo their script in front of an audience
import (
	"bufio"
	//"io"
	"os"
	//"os/Exec"
	"fmt"
	"time"
)

var EnabledRealtype bool
var EnabledWaitForConfirm bool

func init() {
	EnabledRealtype = true       //simulates real typing on terminal
	EnabledWaitForConfirm = true // waits for confirm after each line is typed
}
func main() {
	filepath := os.Args[1]
	f, err := os.Open(filepath) // f contains the pointer to file
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var textlines []string
	for scanner.Scan() {
		textlines = append(textlines, scanner.Text())
	}
	defer f.Close() // file will be closed when main function returns
	var count = 0
	var lines = len(textlines)
	for _, line := range textlines {
		if EnabledRealtype {
			for _, char := range line {
				fmt.Printf("%c", char)
				time.Sleep(100 * time.Millisecond)
			}
		}
		if EnabledWaitForConfirm && count < lines-1 {
			var confirm = ""
			fmt.Printf("\n continue ? type n to break")
			fmt.Scanf("%s", confirm)
			if confirm == "n" {
				return
			}
		} else {
			if count == lines-1 {
				fmt.Printf("\n")
				return
			}
		}
		fmt.Printf("\n")
		count++ // keeps track of line count
	}
}
