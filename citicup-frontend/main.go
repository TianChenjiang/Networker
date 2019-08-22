package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.StaticFS("/networker/user", http.Dir("static/networker/user"))
	router.Run(":8000")
}
