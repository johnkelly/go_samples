package main

import (
	"strings"
	"testing"
)

var testString []string = strings.Split("WTF IS THIS REALLY GOING TO MAKE A PERFORMANCE DIFFERENCE I CANT THINK OF AN ORDER OF MAGNITUDE BEING BECUASE OF THIS LONG STRING", " ")

func BenchmarkEchoSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EchoSlow(testString)
	}
}

func BenchmarkEchoFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EchoFast(testString)
	}
}

// PASS
// BenchmarkEchoSlow-4	  500000	      2613 ns/op
// BenchmarkEchoFast-4	 3000000	       538 ns/op
// ok  	github.com/johnkelly/go_samples/the_go_programming_language/chapter1_exercises/1dot3	3.520s
