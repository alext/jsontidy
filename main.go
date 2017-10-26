package main

import (
	"fmt"
	"os"
)

func main() {
	err := jsontidy(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
