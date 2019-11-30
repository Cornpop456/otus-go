package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func now() string {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return fmt.Sprintf("Current time is %s\n", time.Format("15:04:05"))
}

func main() {
	res := now()
	fmt.Printf(res)
}
