package service

type FizzBuzzService interface {
	GetFizzBuzz(number int) string
}

type FizzBuzz struct {

}

func (fizzbuzz FizzBuzz) GetFizzBuzz(number int) string {
	return "fizz"
}