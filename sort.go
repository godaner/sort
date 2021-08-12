package sort

// Sortable can be sorted
type Sortable interface {
	// Len get Sortable length
	Len() int
	// Less index j less index j?
	Less(i, j int) bool
	// LessObj index i less obj?
	LessObj(i int, obj interface{}) bool
	// LessSortableJ index i less arr index j?
	LessSortableJ(i int, arr Sortable, j int) bool
	// Swap exchange index i and index j?
	Swap(i, j int)
	Set(i int, data interface{})
	// Cover cover index j with  index j
	Cover(i, j int)
	// Get get element by index
	Get(i int) (data interface{})
	// Slice base origin Sortable
	Slice(s, e int) (r Sortable)
	// Append append element
	Append(e interface{})
	// IsDesc Sortable is desc?
	IsDesc() bool
	// IsAsc Sortable is asc?
	IsAsc() bool
}

// Sorter sort the Sortable
type Sorter interface {
	// Sort do sort
	Sort(src Sortable) (dst Sortable)
}
