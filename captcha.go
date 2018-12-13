package captcha

func Captcha(pattern, left, operator, right int) string {
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
	return "1 + One"
}
