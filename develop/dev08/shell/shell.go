package shell

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/mitchellh/go-ps"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func Run() int {
	var env environment
	env.run()
	return 0
}

type environment struct {
	output io.Writer
}

func (env *environment) run() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		command := scanner.Text()

		if strings.Contains(command, "|") {
			if err := env.execPipe(command); err != nil {
				fmt.Errorf("%v\n", err)
			}
		} else {
			env.output = os.Stdout
			if err := env.execCommand(command); err != nil {
				fmt.Errorf("%v\n", err)
			}
		}
		scanner.Scan()
	}
}

func (env *environment) execCommand(command string) error {
	args := strings.Split(command, " ")

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			dir, err := os.UserHomeDir()
			if err != nil {
				return nil
			}
			return os.Chdir(dir)
		} else {
			return os.Chdir(args[1])
		}
	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			return err
		}

		_, err = fmt.Fprintln(env.output, dir)
		if err != nil {
			return err
		}

		return nil
	case "echo":
		for _, elem := range args[1:] {
			_, err := fmt.Fprint(env.output, elem, " ")
			if err != nil {
				return err
			}
		}

		_, err := fmt.Fprintln(env.output)
		if err != nil {
			return err
		}

		return nil
	case "kill":
		if len(args) != 2 {
			return fmt.Errorf("invalid number of args")
		}

		pid, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}

		p, err := ps.Processes()
		if err != nil {
			return err
		}

		for _, elem := range p {
			if elem.Executable() == args[1] {
				pid = elem.Pid()
			}
		}

		if pid == 0 {
			return fmt.Errorf("can`t find %s process", args[1])
		}

		process, err := os.FindProcess(pid)
		if err != nil {
			return err
		}

		err = process.Kill()
		if err != nil {
			return err
		}

		return nil
	case "ps":
		if len(args) > 1 {
			return fmt.Errorf("too many args")
		}

		p, err := ps.Processes()
		if err != nil {
			return err
		}
		for _, elem := range p {
			_, err = fmt.Fprintf(env.output, "%d\t%s\n", elem.Pid(), elem.Executable())
			if err != nil {
				return err
			}
		}

		return nil
	default:
		return fmt.Errorf("command not found: %s", args[0])
	}
}

func (env *environment) execPipe(command string) error {
	commands := strings.Split(command, " | ")
	if len(commands) < 2 {
		return fmt.Errorf("too few args")
	}

	var buf bytes.Buffer
	for _, elem := range commands {
		com := exec.Command(elem)
		args := strings.Split(elem, " ")
		if len(args) > 2 {
			com = exec.Command(args[0], args[1:]...)
		}

		com.Stdin = bytes.NewReader(buf.Bytes())
		buf.Reset()
		com.Stdout = &buf

		if err := com.Start(); err != nil {
			return err
		}

		if err := com.Wait(); err != nil {
			return err
		}

		if _, err := fmt.Fprint(env.output, buf.String()); err != nil {
			return err
		}
	}

	return nil
}
