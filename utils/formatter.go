package utils

// форматирование счета для абонента
func FormatToBill(str string) string {
	for len(str) < 6 {
		str = "0" + str
	}
	return str
}
