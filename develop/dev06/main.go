package main

import (
	"cut/cut"
	"os"
)

func main() {
	os.Exit(cut.Run(os.Args[1:]))
}
