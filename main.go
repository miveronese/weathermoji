package main

import (
  "fmt"
  "os"
)

func main() {

  api := NewWeatherApi(os.Getenv("OPENWEATHER_API_KEY"))



  forecast := api.Now()

  // if forecast.Humidity > 70 { 
  //   fmt.Println(PrintEmoji(":sunny: :sunglasses: :sweat_smile: :icecream:"))
  // } else {
  //   fmt.Println("meh")
  // }

  fmt.Println(forecast.WeatherId)

  
 
  //para criar uma array
  // x := [5]float64{ 98, 93, 77, 82, 83 }

  /// todo: arrays contendo ids de acordo com o tipo de tempo
  ///

  // TODO run twitter bot
  // TODO post every 24 hours (once a day)


}
