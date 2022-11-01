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

  ticker := time.NewTicker(1 * time.Second)
  for duration > 0 {
      fmt.Println("time left:", duration, "minute(s)")
      time.Sleep(1 * time.Minute)
      duration-- 
    }
  ticker.Stop()
  fmt.Println("Time up!")
}

func main() {
  pomodoroTimer()
}

