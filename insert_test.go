package sort

import (
	"testing"
)

func TestInsert_Sort(t *testing.T) {
	t.Log((&Insert{}).Sort(&IntsInterface{Datas: []int{-1, 546, 6, 64, 46, 46, 64, 6, 68989, 79, 313, 0}}))
}
