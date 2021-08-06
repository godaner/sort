package sort

import (
	"sort"
)

type Select struct {
}

func (s *Select) Sort(data sort.Interface) {
	for i := 0; i < data.Len()-1; i++ {
		minIndex := i
		for j := i + 1; j < data.Len(); j++ {
			if data.Less(j, minIndex) {
				minIndex = j
			}
		}
		data.Swap(minIndex, i)
	}
}
