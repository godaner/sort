package sort

type Shell struct {
	Gap int
}

func (s *Shell) Sort(src Interface) (dst Interface) {
	for ; s.Gap >= 1; s.Gap = s.Gap / 2 { // gap/2
		for g := 0; g < s.Gap; g++ { // every group first element
			// simple insert sort
			for i := g; i < src.Len()-s.Gap; i = i + s.Gap {
				for j := s.Gap + 1; j < src.Len(); j = j + s.Gap {
					if !src.Less(j-s.Gap, j) {
						src.Swap(j-s.Gap, j)
					}
				}
			}
		}
	}
	return src
}
