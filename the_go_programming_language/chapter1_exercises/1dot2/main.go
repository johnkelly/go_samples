package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println("Index: ", i, " Arg: ", arg)
	}
}
