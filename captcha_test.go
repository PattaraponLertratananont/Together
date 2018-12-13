package captcha

import "testing"

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

func TestCaptchaGetOneOneOneTwo(t *testing.T) {
	result := Captcha(1, 1, 1, 2)

	if result != "1 + Two" {
		t.Errorf("Its should return 1 + Two but get %q", result)
	}
}

func TestCaptchaGetTwoOneOneTwo(t *testing.T) {
	result := Captcha(2, 1, 1, 2)

	if result != "One + 2" {
		t.Errorf("Its should return One + 2 but get %q", result)
	}
}
