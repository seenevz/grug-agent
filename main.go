package main

import (
	_ "embed"
	"fmt"
	"log"
)

//go:embed .key
var ANTHROPIC_AGENT string

func init() {
	if len(ANTHROPIC_AGENT) == 0 {
		log.Fatal("Anthropic key is missing")
	}
}

func main() {
	fmt.Println(ANTHROPIC_AGENT)
}
