package sort

import (
	"testing"
)

func TestBucket_Sort(t *testing.T) {
	t.Log((&Bucket{}).Sort(&IntArray{Datas: []int{-1, 546, 6, 64, 46, 46, 64, 6, 68989, 79, 313, 0}}))
}
func BenchmarkBucket_Sort(b *testing.B) {
	b.StopTimer()
	sorter := &Bucket{
		OneBucketSize: 100,
		Sorter:        &Quick{},
	}
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
