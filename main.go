//TODO: Print mock forecast in case of -t flag
//TODO: Non empty -t flag checks for day/week/month string matching
//TODO: Valid -t flag makes a json request parses the response and prints out

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type WDef struct {
	Main string `json:"main"`
	Desc string `json: "description"`
}

type Temp struct {
	TempCurrent int `json:"temp"`
	Feels       int `json:"feels_like"`
	TempMin     int `json:"temp_min"`
	TempMax     int `json:"temp_max"`
	Humidity    int `json:"humidity"`
}

type WeatherResponse struct {
	WeatherDescription []WDef `json:"weather"`
	Temperature        Temp   `json:"main"`
}

func getWeather(t string) {
	weatherAPI := os.Getenv("WEATHER_API")
	city := "London" //TODO: dynamic
	country := "GB"  //TODO: dynamic
	weatherURL := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?units=metric&q=%s,%s&appid=%s", city, country, weatherAPI)
	// weatherURL1 := "http://www.api.openweathermap.org/data/2.5/weather"
	log.Print(weatherURL)
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

	re := string(body)
	log.Print((re))
	var weatherResponse WeatherResponse
	json.Unmarshal([]byte(body), &weatherResponse)
	log.Print(weatherResponse.WeatherDescription)
	log.Print(weatherResponse.Temperature)
	// var todos []Todo
	// json.Unmarshal([]byte(body), &todos)
	// log.Print(todos[0])
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("Todo : %+v\n", todos[i])

	// }
}

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	initEnv()
	//TODO: This can be a separate func? Set the flags
	time := flag.String("t", "", "Time (day|week|month) to get the forecast")
	flag.Parse()

	//TODO: Print Defaults if no arguments were passed
	if *time == "" {
		fmt.Println("Provide required time for forecast (day|week|month)")
		flag.PrintDefaults()
		os.Exit(1)
	} else if *time == "day" {
		fmt.Println("Printing DAY forecast")
		getWeather(*time)
		os.Exit(1)
	} else if *time == "week" {
		fmt.Println("Printing WEEK forecast")
		os.Exit(1)
	} else if *time == "month" {
		fmt.Println("Printing MONTH forecast")
		os.Exit(1)
	} else {
		fmt.Println("Please choose a supported forecast param (day|week|month)")
		os.Exit(1)
	}

	//Just print everything out for now
	fmt.Printf("weather: %s\n", *time)
}
