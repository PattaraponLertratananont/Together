package fizzbuzz

import "testing"

func TestFizzBuzzOne(t *testing.T) {
	result := Fizzbuzz(1)

	if result != "1" {
		t.Errorf("Expect 1 but it gets: %q", result)
	}
}
