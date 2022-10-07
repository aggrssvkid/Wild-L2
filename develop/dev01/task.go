package main

import (
	"fmt"
	"os"

	ntp "github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		_, err := fmt.Fprint(os.Stderr, "Не удалось подключиться к хосту")
		if err != nil {
			os.Exit(2)
		}
		os.Exit(1)
	}
	fmt.Println(time)
}
