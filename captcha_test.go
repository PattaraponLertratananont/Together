package captcha

import "testing"

func TestCaptchaGetOneZeroOneOne(t *testing.T) {
	result := Captcha(1, 0, 1, 1)

	if result != "0 + One" {
		t.Errorf("Its should return 0 + One but get %q", result)
	}
}

func TestCaptchaGetAllOne(t *testing.T) {
	result := Captcha(1, 1, 1, 1)

	if result != "1 + One" {
		t.Errorf("Its should return 1 + One but get %q", result)
	}
}

func TestCaptchaGetOneTwoOneOne(t *testing.T) {
	result := Captcha(1, 2, 1, 1)

	if result != "2 + One" {
		t.Errorf("Its should return 2 + One but get %q", result)
	}
}

func TestCaptchaGetOneThreeOneOne(t *testing.T) {
	result := Captcha(1, 3, 1, 1)

	if result != "3 + One" {
		t.Errorf("Its should return 3 + One but get %q", result)
	}
}

func TestCaptchaGetOneFourOneOne(t *testing.T) {
	result := Captcha(1, 4, 1, 1)

	if result != "4 + One" {
		t.Errorf("Its should return 4 + One but get %q", result)
	}
}

func TestCaptchaGetOneFiveOneOne(t *testing.T) {
	result := Captcha(1, 5, 1, 1)

	if result != "5 + One" {
		t.Errorf("It should return 5 + One but get %q", result)
	}
}

func TestCaptchaGetOneSixOneOne(t *testing.T) {
	result := Captcha(1, 6, 1, 1)

	if result != "6 + One" {
		t.Errorf("Its should return 6 + One but get %q", result)
	}
}

func TestCaptchaGetOneSevenOneOne(t *testing.T) {
	result := Captcha(1, 7, 1, 1)

	if result != "7 + One" {
		t.Errorf("Its should return 7 + One but get %q", result)
	}
}

func TestCaptchaGetOneEightOneOne(t *testing.T) {
	result := Captcha(1, 8, 1, 1)

	if result != "8 + One" {
		t.Errorf("Its should return 8 + One but get %q", result)
	}
}

func TestCaptchaGetOneNineOneOne(t *testing.T) {
	result := Captcha(1, 9, 1, 1)

	if result != "9 + One" {
		t.Errorf("It should return 9 + One but get %q", result)
	}
}

func TestCaptchaGetOneZeroOneZero(t *testing.T) {
	result := Captcha(1, 0, 1, 0)

	if result != "0 + Zero" {
		t.Errorf("It should return 0 + Zero but get %q", result)
	}
}

func TestCaptchaGetOneZeroOneTwo(t *testing.T) {
	result := Captcha(1, 0, 1, 2)

	if result != "0 + Two" {
		t.Errorf("It should return 0 + Two but get %q", result)
	}
}

func TestCaptchaGetOneZeroOneThree(t *testing.T) {
	result := Captcha(1, 0, 1, 3)

	if result != "0 + Three" {
		t.Errorf("It should return 0 + Three but get %q", result)
	}
}
func TestCaptchaGetOneZeroOneFour(t *testing.T) {
	result := Captcha(1, 0, 1, 4)

	if result != "0 + Four" {
		t.Errorf("It should return 0 + Four but get %q", result)
	}
}

func TestCaptchaGetOneZeroOneFive(t *testing.T) {
	result := Captcha(1, 0, 1, 5)

	if result != "0 + Five" {
		t.Errorf("It should return 0 + Five but get %q", result)
	}
}

func TestCaptchaGetOneZeroOneSix(t *testing.T) {
	result := Captcha(1, 0, 1, 6)

	if result != "0 + Six" {
		t.Errorf("It should return 0 + Six but get %q", result)
	}
}

func TestCaptchaGetOneZeroOneSeven(t *testing.T) {
	result := Captcha(1, 0, 1, 7)

	if result != "0 + Seven" {
		t.Errorf("It should return 0 + Seven but get %q", result)
	}
}
func TestCaptchaGetOneZeroOneEight(t *testing.T) {
	result := Captcha(1, 0, 1, 8)

	if result != "0 + Eight" {
		t.Errorf("It should return 0 + Eight but get %q", result)
	}
}

func TestCaptchaGetOneZeroOneNine(t *testing.T) {
	result := Captcha(1, 0, 1, 9)

	if result != "0 + Nine" {
		t.Errorf("It should return 0 + Nine but get %q", result)
	}
}
func TestCaptchaGetOneZeroTwoZero(t *testing.T) {
	result := Captcha(1, 0, 2, 0)

	if result != "0 - Zero" {
		t.Errorf("It should return 0 - Zero but get %q", result)
	}
}
