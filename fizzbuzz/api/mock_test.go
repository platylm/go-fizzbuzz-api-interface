package api_test

type mockFizzBuzzService struct {

}

func (mock mockFizzBuzzService) GetFizzBuzz(number int) string {
	return "fizz"
}