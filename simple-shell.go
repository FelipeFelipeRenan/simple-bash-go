package main

import(
	"fmt"
	"bufio"
	"os"
	"os/exec"
	"strings"
	"errors"
	
)

func execInput(input string) error {

	input = strings.TrimSuffix(input, "\n")
	

	args := strings.Split(input, " ")
	switch args[0] {
		case "cd":
			if len(args) < 2{
				return errors.New("path required")
			}
			return os.Chdir(args[1])
		case "go":
			if args[1] == "run" && args[2] == "main.go"{
				return errors.New("Running this shell")
			}
		case "exit":
			fmt.Print("Bye\n")
			os.Exit(0)
		
	}
	
	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}


func main(){
	reader:= bufio.NewReader(os.Stdin)
	for{
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil{
			fmt.Fprintln(os.Stderr, err)
		}

		if err = execInput(input); err != nil{
			fmt.Fprintln(os.Stderr, err)
		}
	}

}

