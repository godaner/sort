package sort

import "sort"

type Sortable interface {
	Sort(data sort.Interface)
}
