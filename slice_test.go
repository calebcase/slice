package slice

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSlice(t *testing.T) {
	xs := Of[int]{}

	xs.Append(1, 2, 3)
	require.EqualValues(t, []int{1, 2, 3}, xs)

	xs.Append(4).Append(5).Append(6)
	require.EqualValues(t, []int{1, 2, 3, 4, 5, 6}, xs)

	xs0 := xs.Copy()
	require.EqualValues(t, []int{1, 2, 3, 4, 5, 6}, xs0)

	xs0.Append(7, 8, 9)
	require.EqualValues(t, []int{1, 2, 3, 4, 5, 6}, xs)
	require.EqualValues(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, xs0)

	xs0.Cut(3, 6)
	require.EqualValues(t, []int{1, 2, 3, 7, 8, 9}, xs0)

	xs.Expand(3, 3)
	require.EqualValues(t, []int{1, 2, 3, 0, 0, 0, 4, 5, 6}, xs)

	xs.Cut(3, 6)
	xs.Extend(3, 3)
	require.EqualValues(t, []int{1, 2, 3, 4, 5, 6, 0, 0, 0}, xs)

	xs.Filter(func(v int) bool {
		return v%2 == 1
	})
	require.EqualValues(t, []int{1, 3, 5}, xs)

	xs.Insert(1, 2, 2)
	require.EqualValues(t, []int{1, 2, 2, 3, 5}, xs)

	xs.Delete(1)
	require.EqualValues(t, []int{1, 2, 3, 5}, xs)

	xs.Push(6, 6)
	require.EqualValues(t, []int{1, 2, 3, 5, 6, 6}, xs)

	v := xs.Pop()
	require.EqualValues(t, 6, v)
	require.EqualValues(t, []int{1, 2, 3, 5, 6}, xs)

	xs.Unshift(-1, -0).Unshift(-2)
	require.EqualValues(t, []int{-2, -1, 0, 1, 2, 3, 5, 6}, xs)

	v = xs.Shift()
	require.EqualValues(t, -2, v)

	xs.Reverse()
	require.EqualValues(t, []int{6, 5, 3, 2, 1, 0, -1}, xs)

	xs.Shift()
	xs.Shift()
	xs.Reverse()
	require.EqualValues(t, []int{-1, 0, 1, 2, 3}, xs)

	rand.Seed(42)
	xs.Shuffle()
	require.EqualValues(t, []int{0, 3, 1, 2, -1}, xs)

	bs := xs.Batch(3)
	// FIXME: This seems like it should work, but it fails...
	//require.EqualValues(t, [][]int{{0, 3, 1}, {2, -1}}, bs)
	require.EqualValues(t, []int{0, 3, 1}, bs[0])
	require.EqualValues(t, []int{2, -1}, bs[1])

	sw := xs.SlidingWindow(3)
	require.EqualValues(t, []int{0, 3, 1}, sw[0])
	require.EqualValues(t, []int{3, 1, 2}, sw[1])
	require.EqualValues(t, []int{1, 2, -1}, sw[2])
}
