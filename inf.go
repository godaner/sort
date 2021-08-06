package sort

type IntsInterface struct {
	Ints []int
}

func (ii *IntsInterface) Len() int {
	return len(ii.Ints)
}

func (ii *IntsInterface) Less(i, j int) bool {
	return ii.Ints[i] < ii.Ints[j]
}

func (ii *IntsInterface) Swap(i, j int) {
	t := ii.Ints[i]
	ii.Ints[i] = ii.Ints[j]
	ii.Ints[j] = t
}
