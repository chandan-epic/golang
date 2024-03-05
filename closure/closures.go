package Closures

func Cosure() func() int {
	return func() int {
		return 1
	}
}
