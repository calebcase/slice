package slice

func Ptr[T any](v T) *T {
	return &v
}

type Slice[T any] []T

func (s *Slice[T]) Append(vs ...T) *Slice[T] {
	*s = append(*s, vs...)

	return s
}

func (s *Slice[T]) Copy() *Slice[T] {
	ns := make(Slice[T], len(*s))

	copy(ns, *s)

	return &ns
}

func (s *Slice[T]) Cut(i, j int) *Slice[T] {
	*s = append((*s)[:i], (*s)[j:]...)

	return s
}

func (s *Slice[T]) Delete(i int) *Slice[T] {
	*s = append((*s)[:i], (*s)[i+1:]...)

	return s
}

func (s *Slice[T]) Expand(i, n int) *Slice[T] {
	*s = append((*s)[:i], append(make(Slice[T], n), (*s)[i:]...)...)

	return s
}

func (s *Slice[T]) Extend(i, n int) *Slice[T] {
	*s = append(*s, make(Slice[T], n)...)

	return s
}

func (s *Slice[T]) Filter(keep func(T) bool) *Slice[T] {
	n := 0

	for _, v := range (*s) {
		if keep(v) {
			(*s)[n] = v
			n++
		}
	}

	*s = (*s)[:n]

	return s
}

func (s *Slice[T]) Insert(i int, vs ...T) *Slice[T] {
	*s = append((*s)[:i], append(vs, (*s)[i:]...)...)

	return s
}

func (s *Slice[T]) Push(vs ...T) *Slice[T] {
	return s.Append(vs...)
}

func (s *Slice[T]) Pop() (v T) {
	v, *s = (*s)[len(*s)-1], (*s)[:len(*s)-1]

	return v
}

func (s *Slice[T]) Unshift(vs ...T) *Slice[T] {
	*s = append(vs, (*s)...)

	return s
}

func (s *Slice[T]) Shift() (v T) {
	v, *s = (*s)[0], (*s)[1:]

	return v
}

func (s *Slice[T]) Reverse() *Slice[T] {
	for i := len((*s))/2-1; i >= 0; i-- {
		opp := len((*s))-1-i
		(*s)[i], (*s)[opp] = (*s)[opp], (*s)[i]
	}

	return s
}

func (s *Slice[T]) Value() []T {
	return *s
}
