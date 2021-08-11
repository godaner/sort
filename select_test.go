package sort

import (
	"testing"
)

func TestSelect_Sort(t *testing.T) {
	t.Log((&Select{}).Sort(&IntArray{Datas: []int{-1, 546, 6, 64, 46, 46, 64, 6, 68989, 79, 313, 0}}))
}

func BenchmarkSelect_Sort(b *testing.B) {
	b.StopTimer()
	sorter := &Select{
	}
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		arr := &IntArray{Datas: unSortIntN(5000)}
		dst := sorter.Sort(arr)
		b.StopTimer()
		if !dst.IsAsc() {
			b.Error("array is not asc")
		}
	}
}
