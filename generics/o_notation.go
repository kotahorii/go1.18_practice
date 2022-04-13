package generics

func Multiple(num, x, y int) (count int) {
	for i := 1; i < num+1; i++ {
		if i%x == 0 || i%y == 0 {
			count += 1
		}
	}
	return
}

func DoubleCard(n, s int) (count int) {
	for i := 1; i < n+1; i++ {
		for j := 1; j < n+1; j++ {
			if i+j > s {
				break
			}
			count += 1
		}
	}
	return
}

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}
