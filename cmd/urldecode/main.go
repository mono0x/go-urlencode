package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
)

func run() error {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return err
		}

		if err == io.EOF && len(line) == 0 {
			return nil
		}

		decoded, err := url.QueryUnescape(line)
		if err != nil {
			return err
		}

		fmt.Print(decoded)
	}
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
