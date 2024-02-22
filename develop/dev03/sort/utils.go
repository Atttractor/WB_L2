package mySort

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func noCopy(data []string) []string {
	exists := make(map[string]struct{}, len(data))
	res := make([]string, 0, len(data))
	for _, elem := range data {
		if _, ok := exists[elem]; ok {
			continue
		}
		res = append(res, elem)
		exists[elem] = struct{}{}
	}

	return res
}

func writeToOutput(data []string) {
	for _, v := range data {
		fmt.Fprintf(os.Stdout, "%s\n", v)
	}
}

// trimNonNumber deletes non number runes from the end of the string
func trimNonNumber(str string) string {
	return strings.TrimRightFunc(str, func(r rune) bool {
		return !unicode.IsNumber(r)
	})
}
