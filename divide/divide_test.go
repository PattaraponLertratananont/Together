package divide

import "testing"

func TestDivideGet2Num(t *testing.T) {
	result := Divide2(14688, 11232)

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
