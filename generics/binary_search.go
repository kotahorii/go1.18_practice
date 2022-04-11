package generics

import "golang.org/x/exp/constraints"


func Binary_search[T constraints.Ordered](a_list []T, n T) bool {
	first := 0
	last := len(a_list) - 1

	for last >= first {
		mid := (first + last) / 2

		if n == a_list[mid] {
			return true
		}

		if n > a_list[mid] {
			first = mid + 1
		} else {
			last = mid - 1
		}
	}
	return false
}
