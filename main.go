package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type WDef struct {
	Main string `json:"main"`
	Desc string `json:"description"`
}

type Current struct {
	Sunrise     int64   `json:"sunrise"`
	Sunset      int64   `json:"sunset"`
	TempCurrent float64 `json:"temp"`
	Humidity    int     `json:"humidity"`
	Feels       float64 `json:"feels_like"`
	Description []WDef  `json:"weather"`
}

type Temp struct {
	Day     int     `json:"day"`
	Night   int     `json:"night"`
	Evening int     `json:"eve"`
	Morning int     `json:"morn"`
	Min     float64 `json:"min"`
	Max     float64 `json:"max"`
}
type DailyForecasts struct {
	Sunrise int64 `json:"sunrise"`
	Sunset  int64 `json:"sunset"`
	Temp    Temp  `json:"temp"`
}

type WeatherResponse struct {
	Description []DailyForecasts `json:"daily"`
	Current     Current          `json:"current"`
}

func formatDayResponse(payload WeatherResponse) {
	log.Println("day response")
	log.Print(payload)
}

func formatWeekResponse(payload WeatherResponse) {
	log.Println("week response")
	log.Print(payload)
}

func getWeather(weatherInterval string) {
	weatherAPI := os.Getenv("WEATHER_API")
	// we need to get dynamic lat long somehow
	lat := "33.441792"  //TODO: Dynamic somehow
	lon := "-94.037689" //TODO: Dynamic somehow
	weatherURL := fmt.Sprintf("http://api.openweathermap.org/data/2.5/onecall?exclude=hourly,minutely,alerts&units=metric&lat=%s&lon=%s&appid=%s", lat, lon, weatherAPI)
	log.Print(weatherURL)

	//Actual GET Req here
	res, err := http.Get(weatherURL)
	//Handle Err
	if err != nil {
		log.Fatal("Error retrieving weather from the endpoint: ", err)
	}
	//Handle non 200 response
	if res.StatusCode != 200 {
		log.Fatal("StatusCode response ", res.StatusCode)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	var weatherResponse WeatherResponse
	json.Unmarshal([]byte(body), &weatherResponse)

	//Format and output of the response handler
	switch weatherInterval {
	case "day":
		formatDayResponse(weatherResponse)
	case "week":
		formatWeekResponse(weatherResponse)
		log.Print(weatherInterval)
	default:
		log.Print("Unsupported weather command, use help command to find out what's supported")
	}
}

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	log.Println("----------------------program start---------------------")
	//Load env file
	initEnv()

	//Load up the cli
	app := &cli.App{}
	app.Name = "Weather forecast in your CLI"
	app.Usage = "Use of the the commands below to get started"
	app.Commands = []*cli.Command{
		{
			Name:    "day",
			Aliases: []string{"d"},
			Usage:   "weather for the current day",
			Action: func(c *cli.Context) error {
				getWeather("day")
				return nil
			},
		},
		{
			Name:    "week",
			Aliases: []string{"w"},
			Usage:   "weather for the current week",
			Action: func(c *cli.Context) error {
				getWeather("week")
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
