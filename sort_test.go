package sort

import "strconv"

func unSortInt() []int {
	data := make([]int, 1<<16)
	for i := 0; i < len(data); i++ {
		data[i] = i ^ 0xcccc
	}
	return data
}

func unSortString() []string {
	unsorted := make([]string, 1<<10)
	for i := range unsorted {
		unsorted[i] = strconv.Itoa(i ^ 0x2cc)
	}
	return unsorted
}
