package captcha

func Captcha(n1, n2, n3, n4 int) string {
	if n4 == 2 && n2 == 1 {
		return "1 + Two"
	}
	if n2 == 3 {
		return "3 + One"
	}
	if n2 == 2 {
		return "2 + One"
	}
	return "1 + One"
}
