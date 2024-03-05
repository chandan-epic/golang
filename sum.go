package Sum

func Summ(a ...int) int {
	var s int = 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}
