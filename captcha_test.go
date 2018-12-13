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
