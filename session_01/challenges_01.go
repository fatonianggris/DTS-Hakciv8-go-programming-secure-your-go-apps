package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const url = "https://jsonplaceholder.typicode.com/posts"

type Payload struct {
	ValueWater  int    `json:"valuewater"`
	ValueWind   int    `json:"valuewind"`
	StatusWater string `json:"statuswater"`
	StatusWind  string `json:"statuswind"`
}

func getStatusWind(value int) string {
	if value < 6 {
		return "aman"
	} else if value >= 7 && value <= 15 {
		return "siaga"
	} else {
		return "bahaya"
	}
}

func getStatusWater(value int) string {
	if value < 5 {
		return "aman"
	} else if value >= 6 && value <= 8 {
		return "siaga"
	} else {
		return "bahaya"
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		valueWater := rand.Intn(100) + 1
		valueWind := rand.Intn(100) + 1
		statusWater := getStatusWater(valueWater)
		statusWind := getStatusWind(valueWind)

		payload := Payload{
			ValueWater:  valueWater,
			ValueWind:   valueWind,
			StatusWater: statusWater,
			StatusWind:  statusWind,
		}

		jsonData, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		resp, err := http.Post(url, "application/json", strings.NewReader(string(jsonData)))
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println("Response Status:", resp.Status)
		fmt.Println("Value Water:", valueWater, "m")
		fmt.Println("Value Wind:", valueWind, "m/s")
		fmt.Println("Status Water:", statusWater)
		fmt.Println("Status Wind:", statusWind)

		time.Sleep(15 * time.Second)
	}
}
