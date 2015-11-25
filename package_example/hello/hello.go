package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/johnkelly/go_samples/package_example/string"
)

func longRunning(i int) {
	duration := time.Duration(rand.Int31n(100)) * time.Millisecond
	time.Sleep(duration)
	fmt.Println(i)
}

func main() {
	fmt.Println("Hello, new gopher!")
	fmt.Println(string.Reverse("Hello, new gopher!"))
	fmt.Println(string.VowelsOnly("Hello, new gopher!"))
	for i := 0; i < 1000; i++ {
		go longRunning(i)
	}
	time.Sleep(5000 * time.Millisecond)
}
