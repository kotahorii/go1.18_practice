package generics

func F(n int) int {
	if n <= 2 {
		return 1
	}
	return F(n-1) + F(n-2)
}
