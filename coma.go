package main

//this package executes bash commands, shows their output and waits for a
//key press to exeute the next command or does it automatically with a time
//delay. It is aimed at providing instructional content for teachers and
//hackers who wish to demo their script in front of an audience
import (
	"bufio"
	"fmt"
	//"io"
	"os"
	"os/exec"
	"strings"
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
		comastr := strings.Fields(line) // extract string from line using space as delimiter
		//comastrlen := len(comastr)
		////TODO: switch if number of command args is less than 2
		//var commandargs = comastr[1]
		//for i := 2; i < comastrlen; i++ {
		//commandargs = fmt.Sprintf("%s %s", commandargs, comastr[i])
		//}

		command := exec.Command(comastr[0], comastr[1:]...) // form command from strings in line

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
		command.Stdout = os.Stdout //TODO: take input and output from file
		command.Stdin = os.Stdin
		if err := command.Start(); err != nil {
			fmt.Println(err)
		}
		command.Wait() // wait for command to execute
		count++        // keeps track of line count
	}
}
