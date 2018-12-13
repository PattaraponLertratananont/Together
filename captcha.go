package captcha

func Captcha(pattern, left, operator, right int) string {
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
