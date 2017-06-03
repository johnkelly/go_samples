package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

var ipCalls map[string]int
var start time.Time

// Implement a wrapper method rateLimitCall that calls the method call for the given IP if there has not been 100 calls from IP in the last second, else throws an exception.

func main() {
	start = time.Now()
	ipCalls = map[string]int{}
	for i := 0; i < 10000; i++ {
		fn := func() int {
			return square(i)
		}
		time.Sleep(5 * time.Millisecond)

		result, err := rateLimitCall(fn, "192.168.0.1")

		if err != nil {
			log.Printf("Hit rate limit for i: %d", i)
			continue
		}
		fmt.Printf("Result: %d - %d\n", i, result)
	}
}

func square(num int) int {
	return num * num
}

func rateLimitCall(f func() int, ip string) (int, error) {
	if time.Since(start).Seconds() > 1 {
		ipCalls = map[string]int{}
		start = time.Now()
	}
	if ipCalls[ip] <= 100 {
		ipCalls[ip] += 1
		return f(), nil
	} else {
		ipCalls[ip] += 1
		return -1, errors.New("Hit rate limit for ip")
	}
}
