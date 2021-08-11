package sort

import (
	"testing"
)

func TestMerge_Sort(t *testing.T) {
	t.Log((&Merge{
		NewArray: func(len int) Array {
			return &IntArray{Datas: make([]int, len, len)}
		},
	}).Sort(&IntArray{Datas: []int{-1, 546, 6, 64, 46, 46, 64, 6, 68989, 79, 313, 0}}))
}
