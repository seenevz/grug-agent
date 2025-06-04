package main

import (
	_ "embed"
	"fmt"
)

//go:embed .key
var ANTHROPIC_AGENT string

func main() {
	fmt.Println(ANTHROPIC_AGENT)
}
