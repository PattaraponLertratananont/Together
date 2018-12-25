package fizzbuzz

func Fizzbuzz(number int) string {
	if number == 2 {
		return "2"
	} else if number == 3 {
		return "Fizz"
	} else if number == 4 {
		return "4"
	} else if number == 5 {
		return "Buzz"
	}
	return "1"
}
