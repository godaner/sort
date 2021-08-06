package sort

import "sort"

type Insert struct {
}

func (i *Insert) Sort(data sort.Interface) {
	for i := 0; i < data.Len()-1; i++ {
		for j := i + 1; j > 0 && !data.Less(j-1, j); j-- {
			data.Swap(j, j-1)
		}
	}
}
