package divide

import "testing"

func TestDivideGet2Num(t *testing.T) {
	result := Divide(14688, 11232)

	if result != 24 {
		t.Errorf("Wrong %v is not correct answer", result)
	}
}
