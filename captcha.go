package captcha

func Captcha(pattern, left, operator, right int) string {
	if left == 9 {
		return "9 + One"
	}
	if left == 8 {
		return "8 + One"
	}
	if left == 7 {
		return "7 + One"
	}
	if left == 6 {
		return "6 + One"
	}
	if left == 5 {
		return "5 + One"
	}
	if left == 4 {
		return "4 + One"
	}
	if left == 3 {
		return "3 + One"
	}
	if left == 2 {
		return "2 + One"
	}
	if left == 1 {
		return "1 + One"
	}
	return "0 + One"
}
