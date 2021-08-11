package sort

import "strconv"

func unSortInt() []int {
	return unSortIntN(1 << 16)
}
func unSortIntN(n int64) []int {
	data := make([]int, n)
	for i := 0; i < len(data); i++ {
		data[i] = i ^ 0xcccc
	}
	return data
}
func unSortStringN(n int64) []string {
	unsorted := make([]string, n)
	for i := range unsorted {
		unsorted[i] = strconv.Itoa(i ^ 0x2cc)
	}
	return unsorted
}

func unSortString() []string {
	return unSortStringN(1 << 10)
}
