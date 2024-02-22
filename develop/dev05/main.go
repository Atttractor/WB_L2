package main

import (
	"grep/grep"
	"os"
)

func main() {
	os.Exit(grep.Run(os.Args[1:]))
}
