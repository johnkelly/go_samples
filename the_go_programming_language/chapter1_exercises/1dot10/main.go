package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	for _, url := range os.Args[1:] {
		fetch(url)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	random := strconv.Itoa(rand.Intn(1000))

	fileName := random + "_results.html"
	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println("File messed up!")
	}

	defer file.Close()

	nbytes, err := io.Copy(file, resp.Body)

	resp.Body.Close()
	if err != nil {
		fmt.Printf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	fmt.Printf("%.2fs  %7d  %s\n", secs, nbytes, url)
}
