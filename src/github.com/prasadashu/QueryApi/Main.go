package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func getWeatherData(apiKey string, city string, weatherChannel chan string, wg *sync.WaitGroup) {

	// Define the API URL
	apiURL := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&APPID=%s", city, apiKey)

	// Defer wait group done to be called at the end of the function
	defer wg.Done()

	// Query Weather API
	response, err := http.Get(apiURL)

	// Check for any error while querying API
	if err != nil {
		fmt.Println("Error while fetching weather: ", err.Error())
	}

	// Defer closing of body at the end of the function
	defer response.Body.Close()

	// Get response body
	body, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Error while reading body: ", err.Error())
	}

	// Insert response body to weather channel
	weatherChannel <- string(body)

	return
}

func main() {
	start := time.Now()

	// Get the API Key from user
	apiKey := flag.String("k", "null", "api key for weather api")
	flag.Parse()

	citySlice := []string{"Chicago", "London", "Atlanta"}
	wg := sync.WaitGroup{}
	weatherChannel := make(chan string)

	for _, city := range citySlice {
		wg.Add(1)
		go getWeatherData(*apiKey, city, weatherChannel, &wg)
	}

	go func() {
		wg.Wait()
		close(weatherChannel)
	}()

	for data := range weatherChannel {
		fmt.Println(data)
	}

	fmt.Println("Time elapsed: ", time.Since(start))
}
