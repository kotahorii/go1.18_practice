package generics

func Solve(l, r int) int {
	if r-l == 1 {
		return l
	}

	m := (l + r) / 2
	s1 := Solve(l, m)
	s2 := Solve(m, r)
	return s1 + s2
}
