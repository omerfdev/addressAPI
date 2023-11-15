package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type City struct {
	Name       string `json:"name"`
	Alpha2Code string `json:"alpha_2_code"`
	Towns      []Town `json:"towns"`
}

type Town struct {
	Name      string     `json:"name"`
	Districts []District `json:"districts"`
}

type District struct {
	Name     string    `json:"name"`
	Quarters []Quarter `json:"quarters"`
}

type Quarter struct {
	Name string `json:"name"`
}

var data []City

func main() {
	loadData()

	router := gin.Default()

	// Endpoint to get all data
	router.GET("/allData", getAllData)

	// Endpoint to get data for a specific city
	router.GET("/cities/:cityName", getCityData)

	router.Run(":8080")
}

func loadData() {
	file, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error reading data file:", err)
		return
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println("Error unmarshalling data:", err)
		return
	}
}

func getAllData(c *gin.Context) {
	c.JSON(http.StatusOK, data)
}

func getCityData(c *gin.Context) {
	cityName := c.Param("cityName")

	for _, city := range data {
		if city.Name == cityName {
			c.JSON(http.StatusOK, city)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "City not found"})
}
