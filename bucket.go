package sort

type Bucket struct {
	OneBucketSize int
	Sorter        Sorter
}

func (b *Bucket) Sort(src Sortable) (dst Sortable) {
	datas := src.(*IntArray).Datas
	max, min := maxAndMin(datas)
	bucketCount := (max - min + 1) / b.OneBucketSize
	buckets := make([][]int, bucketCount+1, bucketCount+1)
	for i := 0; i < len(datas); i++ {
		bIndex := (datas[i] - min) / b.OneBucketSize
		buckets[bIndex] = append(buckets[bIndex], datas[i])
	}
	dst = &IntArray{}
	for _, bucket := range buckets {
		if bucket == nil {
			continue
		}
		dst.AppendSortable(b.Sorter.Sort(&IntArray{Datas: bucket}))
	}
	return dst
}

func maxAndMin(datas []int) (max, min int) {
	if len(datas) == 0 {
		return
	}
	min = datas[0]
	max = min
	for i := 0; i < len(datas); i++ {
		if datas[i] < min {
			min = datas[i]
		}
		if datas[i] > max {
			max = datas[i]
		}
	}
	return max, min
}
