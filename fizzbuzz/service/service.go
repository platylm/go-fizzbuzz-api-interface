package service

import "strconv"

type FizzBuzzService interface {
	GetFizzBuzz(number int) string
}

type FizzBuzz struct {

}

func (fizzbuzz FizzBuzz) GetFizzBuzz(number int) string {
	if number%15 == 0 {
		return "fizzbuzz"
	}
	if number%5 == 0 {
		return "buzz"
	}
	if number%3 == 0 {
		return "fizz"
	}
	return strconv.Itoa(number)
}