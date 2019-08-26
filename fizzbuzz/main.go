package main

import (
	"fizzbuzz/api"
	"flag"
	"github.com/gin-gonic/gin"
)

func main(){
	host := flag.String("host", "3000", "start api host port")
	flag.Parse()
	route := gin.Default()
	fizzBuzzAPI := api.FizzBuzzAPI{}
	route.POST("/api/v1/fizzbuzz", fizzBuzzAPI.FizzBuzzHandler)
	route.Run(":" + *host)
}
