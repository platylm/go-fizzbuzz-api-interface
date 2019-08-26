package api_test

import (
	"fizzbuzz/api"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_FizzBuzz_Input_3_Should_Be_Fizz(t *testing.T){
	expected := `{"message":"fizz"}`
	requestBody := `{"number":3}`
	request := httptest.NewRequest("POST", "/api/v1/fizzbuzz", strings.NewReader(requestBody))
	writer := httptest.NewRecorder()
	mockService := mockFizzBuzzService{}
	fizzBuzzApi := api.FizzBuzzAPI{
		FizzBuzzService: &mockService,
	}

	testRoute := gin.Default()
	testRoute.POST("/api/v1/fizzbuzz", fizzBuzzApi.FizzBuzzHandler)
	testRoute.ServeHTTP(writer, request)
	response := writer.Result()
	actual, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, expected, string(actual))
}