package generics

func F[T StrInt](xs []T) (result []T) {
	result = append(result, xs...)
	return
}

type StrInt interface {
	StrInt()
}

type MyInt int
type MyString string

func (i MyInt) StrInt()    {}
func (s MyString) StrInt() {}
