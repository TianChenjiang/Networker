package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.StaticFS("/networker", http.Dir("networker"))
	router.Run(":8000")
}
