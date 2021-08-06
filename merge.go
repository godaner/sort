package sort

type Merge struct {
	NewInterface func(len int) Interface
}

func (m *Merge) Sort(src Interface) (dst Interface) {
	nd := m.NewInterface(src.Len())
	part(0, src.Len()-1, src, nd)
	return nd
}

func part(left, right int, od, nd Interface) {
	if right-left+1 <= 1 {
		return
	}
	left1 := left
	right1 := (right + left) / 2
	left2 := right1 + 1
	right2 := right
	part(left1, right1, od, nd)
	part(left2, right2, od, nd)
	merge(left1, right1, left2, right2, od, nd)
}

func merge(left1, right1, left2, right2 int, od, nd Interface) {
	nIndex := left1
	for left1 <= right1 && left2 <= right2 {
		if od.Less(left1, left2) {
			nd.Set(nIndex, od.Get(left1))
			left1++
		} else {
			nd.Set(nIndex, od.Get(left2))
			left2++
		}
		nIndex++
	}
	for left1 <= right1 {
		nd.Set(nIndex, od.Get(left1))
		nIndex++
		left1++
	}
	for left2 <= right2 {
		nd.Set(nIndex, od.Get(left2))
		nIndex++
		left2++
	}
}
