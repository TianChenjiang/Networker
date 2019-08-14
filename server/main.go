package main

import (
	"citicup-server/router"
)

func main() {
	router := router.InitRouter()
	router.Run(":8080")
}
