package sort

type Quick struct {
}

func (q *Quick) Sort(src Sortable) (dst Sortable) {
	return part(src)
}

func part(src Sortable) Sortable {
	if src.Len() <= 1 {
		return src
	}
	tmp := src.Get(0)
	left, right := 0, src.Len()-1
	for left < right {
		for left < right && !src.LessObj(right, tmp) {
			right--
		}
		src.Cover(right, left)
		left++
		for left < right && src.LessObj(left, tmp) {
			left++
		}
		src.Cover(left, right)
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
