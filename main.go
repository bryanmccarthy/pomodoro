package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	arg := os.Args[1] // arg is num minutes
	duration, err := strconv.Atoi(arg)
	if err != nil {
		os.Exit(3)
	}

	for duration > 0 {
		fmt.Println("time left:", duration, "minute(s)")
		time.Sleep(1 * time.Minute)
		duration--
	}
	fmt.Println("Time up!")
}
