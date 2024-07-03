package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"os"
	"time"
)

func main() {
	t, err := getTime()
	if err != nil {
		log.Println("Can't take current time: ", err)
		os.Exit(1)
	}
	fmt.Println("Current time:", t)
}

func getTime() (time.Time, error) {
	t, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		return time.Time{}, err
	}
	return t, err
}
