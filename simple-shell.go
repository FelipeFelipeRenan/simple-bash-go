package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/TwiN/go-color"
)

func execInput(input string) error {

	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New(color.Ize(color.Red, "Path required"))
		}
		return os.Chdir(args[1])
	case "go":
		if args[1] == "run" && args[2] == "simple-shell.go" {
			return errors.New(color.Ize(color.Red, "Already running this shell"))
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

func main() {
	reader := bufio.NewReader(os.Stdin)
	host, err := os.Hostname()
	if err != nil {
		print(color.Ize(color.Red, "Strange error"))
	}
	for {
		fmt.Print(color.Ize(color.Green, host), "> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

}
