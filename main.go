package main

import (
	"fmt"
	"os"
)

func main() {

	api := NewWeatherApi(os.Getenv("OPENWEATHER_API_KEY"))
	twitterApi := NewTwiiterService(
		os.Getenv("TWITTER_CONSUMER_KEY"),
		os.Getenv("TWITTER_CONSUMER_SECRET"),
		os.Getenv("TWITTER_API_TOKEN"),
		os.Getenv("TWITTER_API_SECRET"),
	)

	forecast := api.Now()

	fmt.Println(twitterApi)

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
