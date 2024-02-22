package cut

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type environment struct {
	fields    []int
	delimiter string
	separated bool
	reader    io.ReadCloser
}

func Run(args []string) int {
	var env environment

	if err := env.fArgs(args); err != nil {
		fmt.Println(err)
		return 2
	}

	if err := env.run(); err != nil {
		fmt.Println(err)
		return 1
	}

	return 0
}

func (env *environment) fArgs(args []string) error {
	fl := flag.NewFlagSet("cut", flag.ContinueOnError)
	fl.Func("f", "выбрать поля (колонки)", env.parseFields)
	fl.StringVar(&env.delimiter, "d", "\t", "использовать другой разделитель")
	fl.BoolVar(&env.separated, "s", false, "только строки с разделителем")

	if err := fl.Parse(args); err != nil {
		fl.Usage()
		return err
	}

	env.reader = os.Stdin

	return nil
}

func (env *environment) parseFields(s string) error {
	if strings.Contains(s, ",") {
		m := strings.Split(s, ",")
		for _, elem := range m {
			if num, err := strconv.Atoi(elem); err != nil {
				return err
			} else {
				env.fields = append(env.fields, num)
			}
		}
	}
	return nil
}

func (env *environment) run() error {
	scanner := bufio.NewScanner(env.reader)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), env.delimiter)

		if len(line) == 1 {
			if env.separated {
				continue
			} else {
				fmt.Println(line[0])
				continue
			}
		}

		for _, elem := range env.fields {
			fmt.Printf("%s%s", line[elem], env.delimiter)
		}
		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
