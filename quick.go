package sort

type Quick struct {
}

func (q *Quick) Sort(src Array) (dst Array) {
	return part(src)
}

func part(src Array) Array {
	if src.Len() <= 1 {
		return src
	}
	tmp := src.Get(0)
	left, right := 0, src.Len()-1
	for left < right {
		for left < right && !src.LessObj(right, tmp) {
			right--
		}
		src.Set(left, src.Get(right))
		left++
		for left < right && src.LessObj(left, tmp) {
			left++
		}
		src.Set(right, src.Get(left))
		right--
	}
	src.Set(right, tmp)
	if 0 <= right-1 {
		src = part(src.Slice(0, right-1))
	}
	if right+1 <= src.Len()-1 {
		src = part(src.Slice(right+1, src.Len()-1))
	}
	return src
}
