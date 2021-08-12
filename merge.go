package sort

type Merge struct {
	NewSortable func(len int) Sortable
}

func (m *Merge) Sort(src Sortable) (dst Sortable) {
	return m.part(src)
}

func (m *Merge) part(src Sortable) (r Sortable) {
	if src.Len() <= 1 {
		return src
	}
	mid := src.Len() / 2
	r1 := m.part(src.Slice(0, mid))
	r2 := m.part(src.Slice(mid, src.Len()))
	return m.merge(r1, r2)
}

func (m *Merge) merge(r1, r2 Sortable) (r Sortable) {
	r = m.NewSortable(0)
	i, j := 0, 0
	for ; i < r1.Len() && j < r2.Len(); {
		if r1.LessSortableJ(i, r2, j) {
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
