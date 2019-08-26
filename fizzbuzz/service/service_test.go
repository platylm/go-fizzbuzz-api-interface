package service_test

import (
	"fizzbuzz/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetFizzBuzzNumber_Input_3_Should_Be_Fizz(t *testing.T) {
	expected := "fizz"
	number := 3
	fizzBuzzService := service.FizzBuzz{}

	actual := fizzBuzzService.GetFizzBuzz(number)

	assert.Equal(t, expected, string(actual))
}

func Test_GetFizzBuzzNumber_Input_5_Should_Be_Fizz(t *testing.T) {
	expected := "buzz"
	number := 5
	fizzBuzzService := service.FizzBuzz{}

	actual := fizzBuzzService.GetFizzBuzz(number)

	assert.Equal(t, expected, string(actual))
}

func Test_GetFizzBuzzNumber_Input_15_Should_Be_FizzBuzz(t *testing.T) {
	expected := "fizzbuzz"
	number := 15
	fizzBuzzService := service.FizzBuzz{}

	actual := fizzBuzzService.GetFizzBuzz(number)

	assert.Equal(t, expected, string(actual))
}

func Test_GetFizzBuzzNumber_Input_8_Should_Be_8(t *testing.T) {
	expected := "8"
	number := 8
	fizzBuzzService := service.FizzBuzz{}

	actual := fizzBuzzService.GetFizzBuzz(number)

	assert.Equal(t, expected, string(actual))
}
