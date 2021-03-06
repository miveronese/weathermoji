package main

import (
	"fmt"
	"strings"

	"github.com/peterhellberg/emojilib"
)

var ids = map[int]string{
	200: ":thunder_cloud_and_rain: :zap: :umbrella:",
	201: ":thunder_cloud_and_rain: :zap: :umbrella:",
	202: ":thunder_cloud_and_rain: :zap: :umbrella:",
	210: ":thunder_cloud_and_rain: :zap: :umbrella:",
	211: ":thunder_cloud_and_rain: :zap: :umbrella:",
	212: ":thunder_cloud_and_rain: :zap: :umbrella:",
	221: ":thunder_cloud_and_rain: :zap: :umbrella:",
	230: ":thunder_cloud_and_rain: :zap: :umbrella:",
	231: ":thunder_cloud_and_rain: :zap: :umbrella:",
	232: ":thunder_cloud_and_rain: :zap: :umbrella:",
	300: ":closed_umbrella:",
	301: ":closed_umbrella:",
	302: ":closed_umbrella:",
	311: ":closed_umbrella:",
	312: ":closed_umbrella:",
	313: ":closed_umbrella:",
	314: ":closed_umbrella:",
	321: ":closed_umbrella:",
	500: ":cloud_with_rain: :umbrella_without_rain:",
	501: ":cloud_with_rain: :umbrella_without_rain:",
	502: ":cloud_with_rain: :umbrella_without_rain:",
	503: ":cloud_with_rain: :umbrella_without_rain:",
	504: ":cloud_with_rain: :umbrella_without_rain:",
	511: ":cloud_with_rain: :umbrella_without_rain:",
	520: ":cloud_with_rain: :umbrella_without_rain:",
	521: ":cloud_with_rain: :umbrella_without_rain:",
	600: ":snowflake:",
	601: ":snowman:",
	602: ":snowman:",
	611: ":snowflake:",
	612: ":snowflake:",
	615: ":snowflake:",
	616: ":snowflake:",
	620: ":snowflake:",
	621: ":snowflake:",
	622: ":snowflake:",
	701: ":foggy:",
	711: ":foggy:",
	721: ":foggy:",
	731: ":foggy:",
	741: ":foggy:",
	751: ":foggy:",
	761: ":foggy:",
	771: ":foggy:",
	781: ":cloud_with_tornado:",
	800: ":rainbow:",
	801: ":cloud:",
	802: ":cloud:",
	803: ":cloud:",
	804: ":cloud:",
	900: ":cloud_with_tornado:",
	901: ":cloud_with_tornado:",
	902: ":cloud_with_tornado:",
	903: ":penguin:",
	904: ":sunny: :icecream:",
	905: ":leaves:",
	906: "ice_skate:",
	951: ":leaves:",
	952: ":leaves:",
	953: ":leaves:",
	954: ":leaves:",
	955: ":leaves:",
	956: ":leaves:",
	957: ":leaves:",
	958: ":cloud_with_tornado:",
	959: ":cloud_with_tornado:",
	960: ":cloud_with_lightning:",
	961: ":thunder_cloud_and_rain:",
	962: ":thunder_cloud_and_rain:",
}

func removeEmptySpaces(description string) string {
	result := strings.Replace(description, " ", "", -1)
	return result
}

func ConvertToString(forecast Forecast) string {
	forecastString := fmt.Sprint(
		ids[forecast.WeatherId],
		"  :clock12:", forecast.Temperature,
		"˚F  :arrow_up:", forecast.High,
		"˚F  :arrow_down:", forecast.Low, "˚F #",
		removeEmptySpaces(forecast.WeatherLevel),
		" #weather #sanfrancisco",
	)
	return convertToEmoji(forecastString)
}

func convertToEmoji(text string) string {
	return emojilib.ReplaceWithPadding(text)
}
