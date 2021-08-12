package sort

type Count struct {
}

func (c *Count) Sort(src Sortable) (dst Sortable) {
	ia := src.(*IntArray)
	max := maxInt(ia.Datas)
	bucket := make([]int, max+1, max+1)
	for i := 0; i < len(ia.Datas); i++ {
		bucket[ia.Datas[i]]++
	}
	for i, j := 0, 0; i < len(bucket); i++ {
		for k := bucket[i]; k > 0; k-- {
			src.Set(j, i)
			j++
		}
	}
	return src
}

func maxInt(datas []int) (max int) {
	if len(datas) == 0 {
		return 0
	}
	max = datas[0]
	for _, d := range datas {
		if d > max {
			max = d
		}
	}
	return max
}
