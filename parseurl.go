package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: parseurl 'http:/...'\n")
		return
	}

	it, err := url.Parse(os.Args[1])
	if err != nil {
		log.Fatalf("%s", err)
	}

	fmt.Printf("Scheme:      %s\n", it.Scheme)
	fmt.Printf("Hostname:    %s\n", it.Hostname())
	fmt.Printf("Port:        %s\n", it.Port())
	fmt.Printf("EscapedPath: %s\n", it.EscapedPath())

	for k, v := range it.Query() {
		fmt.Printf("Query %s: %s\n", k, v[0])
		for i := 1; i < len(v); i++ {
			fmt.Printf("    %s\n", v[i])
		}
	}
}
