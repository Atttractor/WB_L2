package telnet

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type environment struct {
	timeout time.Duration
	address string
}

func (e *environment) run() error {
	dialer := net.Dialer{
		Timeout: e.timeout,
	}
	conn, err := dialer.Dial("tcp", e.address)
	if err != nil {
		return err
	}

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	ctx, cancel := context.WithCancel(context.Background())
	g := new(errgroup.Group)

	g.Go(func() error {
		reader := bufio.NewReader(os.Stdin)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("n1 stop ch")
				return nil
			default:
				fmt.Print("$: ")
				t, err := reader.ReadString('\n')
				if err != nil {
					return err
				}

				_, err = fmt.Fprint(conn, t)
				if err != nil {
					return err
				}
			}
		}
	})

	g.Go(func() error {
		reader := bufio.NewReader(conn)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("n2 got stop")
				return nil
			default:
				t, err := reader.ReadString('\n')
				if err != nil {
					return err
				}
				fmt.Printf("got from server: %s", t)
			}
		}
	})

	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
		<-stop
		fmt.Println("\nПолучен сигнал завершения")
		cancel()
	}()

	err = g.Wait()
	if err != io.EOF {
		return err
	}
	fmt.Println("Сервер закрыт, выход из программы")

	return nil
}

func (e *environment) fArgs(args []string) error {
	f := flag.NewFlagSet("telnet", flag.ContinueOnError)
	f.DurationVar(&e.timeout, "timeout", time.Second*5, "timeout")

	if err := f.Parse(args); err != nil {
		f.Usage()
		return err
	}

	addr := net.JoinHostPort(f.Arg(0), f.Arg(1))
	e.address = addr

	return nil
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
