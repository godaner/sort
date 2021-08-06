package sort

type Insert struct {
}

func (i *Insert) Sort(src Interface) (dst Interface) {
	for i := 0; i < src.Len()-1; i++ {
		for j := i + 1; j > 0 && !src.Less(j-1, j); j-- {
			src.Swap(j, j-1)
		}
	}
	return src
}
