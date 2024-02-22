package unpacker

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func Run() {
	s, err := unpack("a4bc2d5e")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
}

func unpack(line string) (string, error) {
	var result strings.Builder
	runes := []rune(line)

	for i := 0; i < len(runes); i++ {
		if unicode.IsDigit(runes[i]) {
			if i == 0 {
				return "", nil
			}

			var num strings.Builder
			num.WriteRune(runes[i])
			letter := runes[i-1]

			for j := i + 1; j < len(runes)-1 && unicode.IsDigit(runes[j]); j++ {
				num.WriteRune(runes[j])
			}

			res, err := strconv.Atoi(num.String())
			if err != nil {
				return "", err
			}

			for j := 0; j < res-1; j++ {
				result.WriteRune(letter)
			}

			continue
		}

		_, err := result.WriteRune(runes[i])
		if err != nil {
			return "", err
		}
	}

	return result.String(), nil
}
