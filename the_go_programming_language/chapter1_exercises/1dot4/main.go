package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	seenFiles := make(map[string][]string)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, seenFiles, "")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, seenFiles, arg)
			f.Close()
		}
	}
	for line, files := range seenFiles {
		if len(files) > 1 {
			fmt.Printf("%s\t%s\n", files, line)
		}
	}
}

func countLines(f *os.File, seenFiles map[string][]string, fileName string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if _, ok := seenFiles[input.Text()]; ok {
			seenFiles[input.Text()] = append(seenFiles[input.Text()], fileName)
		} else {
			seenFiles[input.Text()] = []string{fileName}
		}
	}
}
