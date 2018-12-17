package divide

import "testing"

func TestDivideGet2Num(t *testing.T) {
<<<<<<< HEAD
	result := Divide(5, 11232)
=======
	result := Divide2(14688, 11232)
>>>>>>> ef82087ee983a2362675e5ec768f8debda2cc1b4

	if result != 24 {
		t.Errorf("Wrong %v is not correct answer", result)
	}
}

func TestDivideGet3Num(t *testing.T) {
	result := Divide3(250047000, 398223000, 240786000)

	if result != 256 {
		t.Errorf("Wrong %v is not correct answer", result)
	}
}
