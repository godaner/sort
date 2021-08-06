package sort

type Select struct {
}

func (s *Select) Sort(src Interface) (dst Interface) {
	for i := 0; i < src.Len()-1; i++ {
		minIndex := i
		for j := i + 1; j < src.Len(); j++ {
			if src.Less(j, minIndex) {
				minIndex = j
			}
		}
		src.Swap(minIndex, i)
	}
	return src
}
