package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	phrase := os.Args[1:]

	start := time.Now()
	EchoSlow(phrase)
	fmt.Println("Echo slow took: ", time.Since(start))

	start = time.Now()
	EchoFast(phrase)
	fmt.Println("Echo fast took: ", time.Since(start))
}

func EchoSlow(phrase []string) {
	s, sep := "", ""
	for _, arg := range phrase {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func EchoFast(phrase []string) {
	fmt.Println(strings.Join(phrase, " "))
}

// WTF IS THIS REALLY GOING TO MAKE A PERFORMANCE DIFFERENCE I CANT THINK OF AN ORDER OF MAGNITUDE BEING BECUASE OF THIS LONG STRING
// Echo slow took:  67.79µs
// WTF IS THIS REALLY GOING TO MAKE A PERFORMANCE DIFFERENCE I CANT THINK OF AN ORDER OF MAGNITUDE BEING BECUASE OF THIS LONG STRING
// Echo fast took:  5.12µs
