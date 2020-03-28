package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", u.Name)
	fmt.Println("Feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}
