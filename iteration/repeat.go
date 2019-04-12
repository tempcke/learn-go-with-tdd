package iteration

func Repeat (str string) string {
	var result string
	for i := 0; i < 5; i++ {
		result = result + str
	}
	return result
}