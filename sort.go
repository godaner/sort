package sort

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
	Set(i int, data interface{})
	Get(i int, ) (data interface{})
}

type Sortable interface {
	Sort(src Interface) (dst Interface)
}
