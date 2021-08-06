package sort

import (
	"testing"
)

func TestMerge_Sort(t *testing.T) {
	t.Log((&Merge{
		NewInterface: func(len int) Interface {
			return &IntsInterface{Datas: make([]int, len, len)}
		},
	}).Sort(&IntsInterface{Datas: []int{-1, 546, 6, 64, 46, 46, 64, 6, 68989, 79, 313, 0}}))
}
