package main

import (
	mySort "mySort/sort"
	"os"
)

func main() {
	os.Exit(mySort.RunApp(os.Args[1:]))
}
