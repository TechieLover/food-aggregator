package main

import (
	"fmt"
	"net/http"

	"mbrdi/food-aggregator/internal/config"
	"mbrdi/food-aggregator/internal/handlers"
)

func main() {
	fmt.Println(":" + config.Port)
	http.ListenAndServe(":"+config.Port, handlers.GetRouter())
}
