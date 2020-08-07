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
	//var lines = len(textlines)

	for _, line := range textlines {
		if EnabledRealtype {
			for _, char := range line {
				fmt.Printf("%c", char)
				time.Sleep(100 * time.Millisecond)
			}
		}
		comastr := strings.Fields(line)                     // extract string from line using space as delimiter
		command := exec.Command(comastr[0], comastr[1:]...) // form command from strings in line
		var confirm = ""
		fmt.Printf("\n continue ? type n to break")
		fmt.Scanf("%s", &confirm)
		if confirm == "n" {
			return
		}
		fmt.Printf("\n")
		timer := time.Now()
		command.Stdout = os.Stdout //TODO: take input and output from file
		command.Stdin = os.Stdin
		if err := command.Start(); err != nil {
			fmt.Println(err)
		}
		command.Wait()       // wait for command to execute
		timer2 := time.Now() //finish time
		elapsed := timer2.Sub(timer)
		fmt.Printf("\n")
		fmt.Printf("time to execute:%v\n", elapsed)
		count++ // keeps track of line count
	}
}
