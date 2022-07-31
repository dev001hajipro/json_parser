package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/dev001hajipro/json_parser/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hi %s!", user.Username)

	repl.Start(os.Stdin, os.Stdout)
}
