package grep

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func Run(args []string) int {
	var app environment

	if err := app.fArgs(args); err != nil {
		fmt.Println(err)
		return 2
	}

	if err := app.run(); err != nil {
		fmt.Println(err)
		return 1
	}

	return 0
}

type environment struct {
	after   int
	before  int
	context int
	count   int
	ignore  bool
	invert  bool
	fixed   bool
	lnum    bool
	pattern string
	reader  io.ReadCloser
	inp     []string
}

func (env *environment) fArgs(args []string) error {
	fl := flag.NewFlagSet("grep", flag.ContinueOnError)
	fl.IntVar(&env.after, "A", 0, "печатать +N строк после совпадения")
	fl.IntVar(&env.before, "B", 0, "печатать +N строк до совпадения")
	fl.IntVar(&env.context, "C", 0, "(A+B) печатать ±N строк вокруг совпадения")
	fl.IntVar(&env.count, "c", -1, "количество строк")
	fl.BoolVar(&env.ignore, "i", false, "игнорировать регистр")
	fl.BoolVar(&env.invert, "V", false, "вместо совпадения, исключать")
	fl.BoolVar(&env.fixed, "F", false, "точное совпадение со строкой, не паттерн")
	fl.BoolVar(&env.lnum, "n", false, "печатать номер строки")

	if err := fl.Parse(args); err != nil {
		fl.Usage()
		return err
	}

	if env.after == 0 {
		env.after = env.context
	}
	if env.before == 0 {
		env.before = env.context
	}

	if env.ignore {
		env.pattern = "(?i)"
	}

	env.pattern = fl.Arg(0)
	if env.fixed {
		env.pattern = fmt.Sprintf("^%s$", env.pattern)
	}

	file, err := os.Open(fl.Arg(1))
	if err != nil {
		fmt.Printf("Ошибка при открытии файла %v: %v", fl.Arg(1), err)
		return err
	}
	env.reader = file

	return nil
}

func (env *environment) run() error {
	defer env.reader.Close()

	scanner := bufio.NewScanner(env.reader)
	for scanner.Scan() {
		env.inp = append(env.inp, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	r, err := regexp.Compile(env.pattern)
	if err != nil {
		return err
	}

	matched := make([]int, 0, len(env.inp)/2)

	for i, v := range env.inp {
		if env.count == 0 {
			break
		}
		if r.MatchString(v) && !env.invert || !r.MatchString(v) && env.invert {
			matched = append(matched, i)
			env.count--
		}
	}

	env.printRes(matched)

	return nil
}

func (env *environment) printRes(matched []int) {
	printed := make(map[int]struct{})
	matchedM := make(map[int]struct{})
	for _, elem := range matched {
		matchedM[elem] = struct{}{}
	}

	for _, elem := range matched {
		if env.before > 0 || env.after > 0 {
			if _, ok := printed[elem]; ok {
				continue
			}

			start := elem - env.before
			if elem-env.before < 0 {
				start = 0
			}
			finish := elem + env.after
			if elem+env.after > len(env.inp)-1 {
				finish = len(env.inp) - 1
			}

			for ; start <= finish; start++ {
				if _, ok := printed[start]; ok {
					continue
				}
				if _, ok := matchedM[start]; ok && start != elem {
					break
				}
				if env.lnum {
					if _, ok := matchedM[start]; ok {
						fmt.Printf("%d:%s\n", start+1, env.inp[start])
					} else {
						fmt.Printf("%d-%s\n", start+1, env.inp[start])
					}
					printed[start] = struct{}{}
					continue
				}
				fmt.Println(env.inp[start])
				printed[start] = struct{}{}
			}
			continue
		}

		if env.lnum {
			fmt.Printf("%d:%s\n", elem+1, env.inp[elem])
			continue
		}
		fmt.Println(elem)
	}
}
