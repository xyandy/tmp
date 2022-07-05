package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"flights"
)

var (
	port int
)

type FlightJson struct {
	Flights [][]string `json:"flights"`
}

// gin handler
func sortFlightsHandler(c *gin.Context) {
	var flightJson FlightJson
	if err := c.ShouldBindJSON(&flightJson); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	ret := flights.GetSrcDestFromFlights(flightJson.Flights)
	c.JSON(http.StatusOK, ret)
}

func main() {
	// flag.IntVar(&port, "port", 9000, "http server port")

	// gin.SetMode(gin.ReleaseMode)
	// router := gin.Default()
	// router.HandleMethodNotAllowed = true
	// router.POST("/sort_flights", sortFlightsHandler)

	// server := &http.Server{
	// 	Addr:         fmt.Sprintf(":%d", port),
	// 	Handler:      router,
	// 	ReadTimeout:  10 * time.Second,
	// 	WriteTimeout: 10 * time.Second,
	// }

	// fmt.Println("server starts at port", port)

	// err := server.ListenAndServe()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	fmt.Println("server starts at port", port)
	time.Sleep(time.Second * 20)
	fmt.Println("server stops at port", port)
}
