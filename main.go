package main

import (
	"fmt"
	"golang_repl/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Welcome, %s \n\n", user.Name)

	r := repl.NewRepl(">>> ", []string{"q", "quit"})
	r.Start(os.Stdin, os.Stdout)
}
