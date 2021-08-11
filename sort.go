package sort

type Array interface {
	Len() int
	Less(i, j int) bool
	LessObj(i int, obj interface{}) bool
	LessArray(i int, arr Array, j int) bool
	Swap(i, j int)
	Set(i int, data interface{})
	Get(i int) (data interface{})
	Slice(s, e int) (r Array)
	Append(i interface{})
	IsDesc() bool
	IsAsc() bool
}

type Sortable interface {
	Sort(src Array) (dst Array)
}
