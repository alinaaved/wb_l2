package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

const host = "0.ru.pool.ntp.org"

func getTime() (time.Time, error) {
	return ntp.Time(host)
}

func main() {
	currentTime, err := getTime()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(currentTime)

}
