package main

import (
	"fmt"
	"time"
)

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

	// fmt.Println("server starts", port)

	// err := server.ListenAndServe()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	fmt.Println("server starts")
	time.Sleep(time.Second * 5)
	fmt.Println("server stops")
}
