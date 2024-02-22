package main

import (
	"fmt"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func toLowercase(data []string) []string {
	result := make([]string, 0, len(data))

	for _, elem := range data {
		result = append(result, strings.ToLower(elem))
	}

	return result
}

func equal(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	letters := make(map[rune]struct{})

	for _, elem := range str1 {
		letters[elem] = struct{}{}
	}

	for _, elem := range str1 {
		if _, ok := letters[elem]; !ok {
			return false
		}
	}

	return true
}

func notExists(mas []string, str string) bool {
	for _, elem := range mas {
		if str == elem {
			return false
		}
	}
	return true
}

func searchAnagram(data []string) map[string][]string {
	result := make(map[string][]string)
	data = toLowercase(data)

	for _, elem := range data {
		flag := true
		for key := range result {
			if equal(elem, key) {
				if notExists(result[key], elem) && elem != key {
					result[key] = append(result[key], elem)
				}
				flag = false
				continue
			}
		}

		if flag {
			result[elem] = []string{}
		}
	}

	for key, value := range result {
		if len(value) <= 1 {
			delete(result, key)
		}
	}

	return result
}

func main() {
	data := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "СтоЛИК", "ПЯТКА", "листок", "asdf", "fdsa"}
	fmt.Println(searchAnagram(data))

}
