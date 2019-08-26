package api

import (
	"fizzbuzz/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FizzBuzzRequest struct {
	Number int `json:"number"`
}

type FizzBuzzResponse struct {
	Message string `json:"message"`
}

type FizzBuzzAPI struct {
	FizzBuzzService service.FizzBuzzService
}

func (api FizzBuzzAPI) FizzBuzzHandler (context *gin.Context) {
	var fizzBuzzRequest FizzBuzzRequest
	err := context.Bind(&fizzBuzzRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}
	number := fizzBuzzRequest.Number
	message := api.FizzBuzzService.GetFizzBuzz(number)
	response := FizzBuzzResponse{Message:message}
	context.JSON(http.StatusOK,response)
}