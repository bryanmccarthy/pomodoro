package main

import (
	"fmt"
	"time"
  "os"
  "strconv"
)

// First arg is duration
// Ticks for duration seconds
func pomodoroTimer() {
  arg := os.Args[1]
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

func main() {
  pomodoroTimer()
}

