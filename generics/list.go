package generics

func F[T StrInt](xs []T) (result []T) {
	result = append(result, xs...)
	return
}

type StrInt interface{}

type MyInt int
type MyString string

func (i MyInt) String()    {}
func (s MyString) String() {}
