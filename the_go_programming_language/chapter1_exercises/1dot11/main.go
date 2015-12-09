package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	urls := make([]string, 0, 1000000)

	f, _ := os.Open("top-1m.csv")

	r := csv.NewReader(bufio.NewReader(f))

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		urls = append(urls, "http://"+record[1])
	}

	for _, url := range urls[1000:1500] {
		go fetch(url, ch)
	}
	for range urls[:1000] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

// Get http://www.indiamart.com/: dial tcp: lookup www.indiamart.com: no such host
// Get http://axisbank.co.in: dial tcp: lookup axisbank.co.in: no such host
// Get http://himasoku.com: dial tcp: lookup himasoku.com: no such host
// Get http://51sole.com: dial tcp: lookup 51sole.com: no such host
// Get http://feng.com: dial tcp: lookup feng.com: no such host
// Get http://yhd.com: dial tcp: lookup yhd.com: no such host
// while reading http://ddanzi.com: unexpected EOF
// Get http://52pk.com: dial tcp 115.182.153.16:80: i/o timeout

//It hangs waiting for a response if the site doesn't respond
