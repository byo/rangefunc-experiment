package main

import (
	"bufio"
	"fmt"
	"io"
	"iter"
	"strings"
)

const text = `Some
long
multiline
text
`

func splitLines(r io.Reader) iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		scanner := bufio.NewScanner(r)
		for lineNo := 1; scanner.Scan(); lineNo++ {
			if !yield(lineNo, scanner.Text()) {
				break
			}
		}
	}
}

func main() {
	r := strings.NewReader(text)
	for lineNo, line := range splitLines(r) {
		fmt.Printf("%d: %s\n", lineNo, line)
	}
}
