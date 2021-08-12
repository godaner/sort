package sort

import (
	"testing"
)

func TestCount_Sort(t *testing.T) {
	t.Log((&Count{}).Sort(&IntArray{Datas: []int{-1, 546, 6, 64, 46, 46, 64, 6, 68989, 79, 313, 0}}))
}
func BenchmarkCount_Sort(b *testing.B) {
	b.StopTimer()
	sorter := &Count{}
	for i := 0; i < b.N; i++ {
		arr := &IntArray{Datas: unSortIntN(5000)}
		b.StartTimer()
		dst := sorter.Sort(arr)
		b.StopTimer()
		if !dst.IsAsc() {
			b.Error("array is not asc")
		}
	}
}
