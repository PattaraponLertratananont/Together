package fizzbuzz

import "testing"

func TestFizzBuzzOne(t *testing.T) {
	result := Fizzbuzz(1)

	if result != "1" {
		t.Errorf("Expect 1 but it gets: %q", result)
	}
}
func TestFizzBuzzTwo(t *testing.T) {
	result := Fizzbuzz(2)

	if result != "2" {
		t.Errorf("Expect 2 but it gets: %q", result)
	}
}
func TestFizzBuzzThree(t *testing.T) {
	result := Fizzbuzz(3)

	if result != "Fizz" {
		t.Errorf("Expect Fizz but it gets: %q", result)
	}
}

func TestFizzBuzzFour(t *testing.T) {
	result := Fizzbuzz(4)

	if result != "4" {
		t.Errorf("Expect 4 but it gets %q", result)
	}
}
