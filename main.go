package main

import (
  "fmt"
  "os"
)

func main() {
  
  api := NewWeatherApi(os.Getenv("OPENWEATHER_API_KEY"))

  fmt.Print(api.Now())

  // TODO run twitter bot
  // TODO post every 24 hours (once a day)
}
