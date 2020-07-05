package main

import (
	"fmt"
	"os"
	"os/user"
	"monkey-interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Welcome to the Monkey REPL, %s!\n", user.Username)
	fmt.Printf("You can type commands.\n")
	repl.Start(os.Stdin, os.Stdout)
}
