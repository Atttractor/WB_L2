package main

import (
	"dooplik-dev01/ntp"
	"os"
)

// Вызов ntp.Run вернет код ошибки от 0 до 2
func main() {
	os.Exit(ntp.Run(os.Args[1:]))
}
