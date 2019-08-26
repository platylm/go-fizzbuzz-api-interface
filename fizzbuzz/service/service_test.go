package service_test

import (
	"fizzbuzz/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetFizzBuzzNumber_Input_3_Should_Be_Fizz(t *testing.T){
	expected := "fizz"
	number := 3
	fizzBuzzService := service.FizzBuzz{}

	actual := fizzBuzzService.GetFizzBuzz(number)

	assert.Equal(t, expected, string(actual))
}