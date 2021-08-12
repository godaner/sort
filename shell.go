package sort

type Shell struct {
	Gap  int
	cGap int
}

func (s *Shell) Sort(src Sortable) (dst Sortable) {
	s.cGap = s.Gap
	for ; s.cGap >= 1; s.cGap = s.cGap / 2 { // cGap/2
		for first := 0; first < s.cGap; first++ { // every group first element
			// simple insert sort
			for i := first; i < src.Len()-s.cGap; i = i + s.cGap {
				for j := s.cGap + i; j > first && !src.Less(j-s.cGap, j); j = j - s.cGap {
					src.Swap(j-s.cGap, j)
				}
			}
		}
	}
	return src
}
