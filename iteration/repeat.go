package iteration

func Repeat (str string, count int) string {
	var result string
	for i := 0; i < count; i++ {
		result += str
	}
	return result
}