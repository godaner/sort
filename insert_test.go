package sort

import (
	"sort"
	"testing"
)

func TestInsert_Sort(t *testing.T) {
	var data sort.Interface = &IntsInterface{Ints: []int{-1, 546, 6, 64, 46, 46, 64, 6, 68989, 79, 313, 0}}
	(&Insert{}).Sort(data)
	t.Log(data)
}
