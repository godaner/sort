package sort

type BubbleSort struct {
}

func (b *BubbleSort) Sort(src Array) (dst Array) {
	for i := 0; i < src.Len()-1; i++ {
		for j := 0; j < src.Len()-1-i; j++ {
			if !src.Less(j, j+1) {
				src.Swap(j, j+1)
			}
		}
	}
	return src
}
