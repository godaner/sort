package sort

import (
	"sort"
)

type BubbleSort struct {
}

func (b *BubbleSort) Sort(data sort.Interface) {
	for i := 0; i < data.Len()-1; i++ {
		for j := 0; j < data.Len()-1-i; j++ {
			if !data.Less(j, j+1) {
				data.Swap(j, j+1)
			}
		}
	}
}
