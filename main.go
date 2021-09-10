package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type coins struct {
	id                          string  `json:"id"`
	symbol                      string  `json:"symbol`
	name                        string  `json:"name"`
	current_price               int     `json:"current_price"`
	market_cap_rank             int     `json:"market_cap_rank"`
	fully_diluted_valuation     int     `json:"fully_diluted_valuation"`
	total_volume                int     `json:"total_volume`
	price_change_24h            float64 `json:"price_change_24h`
	price_change_percentage_24h float64 `json:"price_change_percentage_24h"`
	market_cap_change_24h       float64 `json:"market_cap_change_24h"`
}

var coinsData = []coins{
	{
		id:                          "bitcoin",
		symbol:                      "btc",
		name:                        "Bitcoin",
		current_price:               10000,
		market_cap_rank:             1,
		fully_diluted_valuation:     97000,
		total_volume:                34449404895,
		price_change_24h:            31.08,
		price_change_percentage_24h: 0.06679,
		market_cap_change_24h:       0.096,
	},
}

func getCoinsStatus(c *gin.Context) { // RETURN HTTP status to the console
	c.IndentedJSON(http.StatusOK, coinsData)
}
func postCoinsStatus(c *gin.Context) { //bind ibto the json file
	var newCoins coins

	// Call BindJSON to bind the received JSON to

	if err := c.BindJSON(&newCoins); err != nil {
		return
	}

	// Add the new album to the slice.
	coinsData = append(coinsData, newCoins)
	c.IndentedJSON(http.StatusCreated, newCoins)
}
func getCoinsByID(c *gin.Context) { //get elements by the id no
	id := c.Param("id")

	for _, a := range coinsData {
		if a.id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "data not found"})
}
func main() {
	//fmt.Printf("Start")
	jsonFile, debug := os.Open("API.json")

	if debug != nil {
		fmt.Println(debug)
	}
	fmt.Println("Successfully Opened the API.json file")
	router := gin.Default()
	router.GET("/coins", getCoinsStatus)
	router.GET("/coins/:id", getCoinsByID)
	router.POST("/coins", postCoinsStatus)

	router.Run("localhost:8080")
}
