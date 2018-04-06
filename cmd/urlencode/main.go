package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"strings"
)

func run() error {
	var reader *bufio.Reader
	if len(os.Args) >= 2 {
		reader = bufio.NewReader(strings.NewReader(strings.Join(os.Args[1:], " ")))
	} else {
		reader = bufio.NewReader(os.Stdin)
	}
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return err
		}

		if err == io.EOF && len(line) == 0 {
			return nil
		}

		encoded := url.QueryEscape(line)

		fmt.Print(encoded)
	}
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
