package sort

import (
	"testing"
)

func TestShell_Sort(t *testing.T) {
	t.Log((&Shell{
		Gap: 5,
	}).Sort(&IntArray{Datas: []int{-1, 546, 6, 64, 46, 46, 64, 6, 68989, 79, 313, 0}}))
}

func BenchmarkShell_Sort(b *testing.B) {
	b.StopTimer()
	sorter := &Shell{
		Gap: 10,
	}
	for i := 0; i < b.N; i++ {
		arr := &IntArray{Datas: unSortIntN(15)}
		b.StartTimer()
		dst := sorter.Sort(arr)
		b.StopTimer()
		if !dst.IsAsc() {
			b.Error("array is not asc")
		}
	}
}
