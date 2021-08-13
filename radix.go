package sort

import (
	"fmt"
	"math"
)

type Radix struct {
}

func (r *Radix) Sort(src Sortable) (dst Sortable) {
	datas := src.(*IntArray).Datas
	// 当前要排序的位数，从后往前是[0,+)
	digitPos := 0
	maxLen := len(fmt.Sprint(maxInt(datas)))
	for digitPos <= maxLen {
		buckets := make([][]int, 10, 10)
		for i := 0; i < len(datas); i++ {
			// 获取每个数的digitPos位置的数字d，并且放入相应的bucket中
			pos := int(math.Pow10(digitPos))
			d := datas[i] / pos % 10
			buckets[d] = append(buckets[d], datas[i])
		}
		datas = make([]int, 0, 0)
		for i := 0; i < len(buckets); i++ {
			datas = append(datas, buckets[i]...)
		}
		digitPos++
	}
	return &IntArray{Datas: datas}
}
