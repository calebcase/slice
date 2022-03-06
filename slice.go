// Package slice provides generic implementations of common slice operations
// (e.g. https://github.com/golang/go/wiki/SliceTricks).
package slice

import "math/rand"

type Of[T any] []T

func (s *Of[T]) Append(vs ...T) *Of[T] {
	*s = append(*s, vs...)

	return s
}

func (s Of[T]) Copy() Of[T] {
	ns := make(Of[T], len(s))

	copy(ns, s)

	return ns
}

func (s *Of[T]) Cut(i, j int) *Of[T] {
	*s = append((*s)[:i], (*s)[j:]...)

	return s
}

func (s *Of[T]) Delete(i int) *Of[T] {
	*s = append((*s)[:i], (*s)[i+1:]...)

	return s
}

func (s *Of[T]) Expand(i, n int) *Of[T] {
	*s = append((*s)[:i], append(make(Of[T], n), (*s)[i:]...)...)

	return s
}

func (s *Of[T]) Extend(i, n int) *Of[T] {
	*s = append(*s, make(Of[T], n)...)

	return s
}

func (s *Of[T]) Filter(keep func(T) bool) *Of[T] {
	n := 0

	for _, v := range *s {
		if keep(v) {
			(*s)[n] = v
			n++
		}
	}

	var zero T
	for i := n; i < len(*s); i++ {
		(*s)[i] = zero
	}

	*s = (*s)[:n]

	return s
}

func (s *Of[T]) Insert(i int, vs ...T) *Of[T] {
	*s = append((*s)[:i], append(vs, (*s)[i:]...)...)

	return s
}

func (s *Of[T]) Push(vs ...T) *Of[T] {
	return s.Append(vs...)
}

func (s *Of[T]) Pop() (v T) {
	v, *s = (*s)[len(*s)-1], (*s)[:len(*s)-1]

	return v
}

func (s *Of[T]) Unshift(vs ...T) *Of[T] {
	*s = append(vs, (*s)...)

	return s
}

func (s *Of[T]) Shift() (v T) {
	v, *s = (*s)[0], (*s)[1:]

	return v
}

func (s *Of[T]) Reverse() *Of[T] {
	for i := len((*s))/2 - 1; i >= 0; i-- {
		opp := len((*s)) - 1 - i
		(*s)[i], (*s)[opp] = (*s)[opp], (*s)[i]
	}

	return s
}

func (s *Of[T]) Shuffle() *Of[T] {
	for i := len(*s) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}

	return s
}

func (s Of[T]) Batch(n int) []Of[T] {
	bs := make([]Of[T], 0, (len(s)+n-1)/n)

	for n < len(s) {
		s, bs = s[n:], append(bs, s[0:n:n])
	}

	return append(bs, s)
}

func (s Of[T]) SlidingWindow(n int) []Of[T] {
	if len(s) <= n {
		return []Of[T]{s}
	}

	w := make([]Of[T], 0, len(s)-n+1)

	for i, j := 0, n; j <= len(s); i, j = i+1, j+1 {
		w = append(w, s[i:j])
	}

	return w
}
