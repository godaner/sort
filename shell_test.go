package sort

import (
	"testing"
)

func TestShell_Sort(t *testing.T) {
	t.Log((&Shell{
		Gap: 5,
	}).Sort(&IntArray{Datas: []int{-1, 546, 6, 64, 46, 46, 64, 6, 68989, 79, 313, 0}}))
}
