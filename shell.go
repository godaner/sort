package sort

import (
	"sort"
)

type Shell struct {
	Gap int
}

func (s *Shell) Sort(data sort.Interface) {
	for ; s.Gap >= 1; s.Gap = s.Gap / 2 { // gap/2
		for g := 0; g < s.Gap; g++ { // every group first element
			// simple insert sort
			for i := g; i < data.Len()-s.Gap; i = i + s.Gap {
				for j := s.Gap + 1; j < data.Len(); j = j + s.Gap {
					if !data.Less(j-s.Gap, j) {
						data.Swap(j-s.Gap, j)
					}
				}
			}
		}
	}

}
