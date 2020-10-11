//TODO: Print mock forecast in case of -t flag
//TODO: Non empty -t flag checks for day/week/month string matching
//TODO: Valid -t flag makes a json request parses the response and prints out

package main

import (
	// "encoding/json"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func getWeather(t string) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) //body is json
	// fmt.Printf(": %+v", len(body))
	var todos []Todo
	json.Unmarshal([]byte(body), &todos)
	fmt.Printf("Todos : %+v", todos)
	// var data []string
	// _ = json.Unmarshal([]byte(body), &data)
	// log.Printf("Unmarshaled: %v", data)
	// fmt.Printf("JSON: %s", body)
}

func main() {
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
