package sort

type Merge struct {
	NewArray func(len int) Array
}

func (m *Merge) Sort(src Array) (dst Array) {
	return m.part(src)
}

func (m *Merge) part(src Array) (r Array) {
	if src.Len() <= 1 {
		return src
	}
	mid := src.Len() / 2
	r1 := m.part(src.Slice(0, mid))
	r2 := m.part(src.Slice(mid, src.Len()))
	return m.merge(r1, r2)
}

func (m *Merge) merge(r1, r2 Array) (r Array) {
	r = m.NewArray(0)
	i, j := 0, 0
	for ; i < r1.Len() && j < r2.Len(); {
		if r1.LessArray(i, r2, j) {
			r.Append(r1.Get(i))
			i++
		} else {
			r.Append(r2.Get(j))
			j++
		}
	}
	for i < r1.Len() {
		r.Append(r1.Get(i))
		i++
	}
	for j < r2.Len() {
		r.Append(r2.Get(j))
		j++
	}
	return r
}
