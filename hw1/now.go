package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func now() (string, error) {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Current time is %s\n", time.Format("15:04:05")), nil
}

func main() {
	res, err := now()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf(res)
}
