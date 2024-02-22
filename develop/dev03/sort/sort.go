package mySort

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func RunApp(args []string) int {
	var app appEnv
	if err := app.fArgs(args); err != nil {
		return 2
	}

	if err := app.run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Run error: %v\n", err)
		return 1
	}

	return 0
}

type appEnv struct {
	col        int
	num        bool
	reverse    bool
	noCopy     bool
	readCloser io.ReadCloser
}

func (app *appEnv) fArgs(args []string) error {
	fl := flag.NewFlagSet("sort", flag.ContinueOnError)
	fl.IntVar(&app.col, "k", 1, "Указание колонки для сортировки")
	fl.BoolVar(&app.num, "n", false, "Сортировать по числовому значению")
	fl.BoolVar(&app.reverse, "r", false, "Сортировать в обратном порядке")
	fl.BoolVar(&app.noCopy, "u", false, "Не выводить повторяющиеся строки")

	if err := fl.Parse(args); err != nil {
		fl.Usage()
		return err
	}

	file, err := os.Open(fl.Arg(0))
	if err != nil {
		return err
	}
	app.readCloser = file

	return nil

}

func (app *appEnv) run() error {
	defer app.readCloser.Close()
	data := make([]string, 0)

	scanner := bufio.NewScanner(app.readCloser)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	data = app.sort(data)

	for _, v := range data {
		_, _ = fmt.Fprintf(os.Stdout, "%s\n", v)
	}

	return nil
}

func (app *appEnv) sort(data []string) []string {
	if app.reverse {
		sort.Sort(sort.Reverse(sort.StringSlice(data)))
	} else {
		sort.Strings(data)
	}

	if app.noCopy {
		data = noCopy(data)
	}

	return data
}

func (app *appEnv) sortColumns(data []string) []string {
	t := stringTable{
		data:      make([][]string, 0, len(data)),
		column:    app.col - 1,
		isNumeric: app.num,
	}

	for _, v := range data {
		t.data = append(t.data, strings.Fields(v))
	}

	if app.reverse {
		sort.Sort(sort.Reverse(t))
	} else {
		sort.Sort(t)
	}

	for i, v := range t.data {
		data[i] = strings.Join(v, " ")
	}

	if app.noCopy {
		data = noCopy(data)
	}

	return data
}

type stringTable struct {
	data      [][]string
	column    int
	isNumeric bool
}

func (t stringTable) Len() int {
	return len(t.data)
}

func (t stringTable) Less(i, j int) bool {
	col := t.column
	if col > len(t.data[i])-1 || col > len(t.data[j]) {
		col = 0
	}

	if t.isNumeric {
		n1 := trimNonNumber(t.data[i][col])
		n2 := trimNonNumber(t.data[j][col])

		i1, err := strconv.Atoi(n1)
		if err != nil {
			return t.data[i][col] < t.data[j][col]
		}
		j1, err := strconv.Atoi(n2)
		if err != nil {
			return t.data[i][col] < t.data[j][col]
		}

		return i1 < j1
	}
	return t.data[i][col] < t.data[j][col]
}

func (t stringTable) Swap(i, j int) {
	t.data[i], t.data[j] = t.data[j], t.data[i]
}
