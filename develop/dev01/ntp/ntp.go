package ntp

import (
	"flag"
	"fmt"
	"time"

	"github.com/beevik/ntp"
)

type agruments struct {
	host string
}

func Run(args []string) int {
	var argVar agruments

	fl := flag.NewFlagSet("ntp", flag.ContinueOnError)
	fl.StringVar(&argVar.host, "host", "0.beevik-ntp.pool.ntp.org", "ntp host")

	fmt.Println(args)

	if err := fl.Parse(args); err != nil {
		return 2
	}

	if err := getTime(argVar); err != nil {
		return 1
	}

	return 0
}

func getTime(args agruments) error {
	t, err := ntp.Time(args.host)
	if err != nil {
		return err
	}

	fmt.Println(t.UTC().Format(time.UnixDate))
	return nil
}