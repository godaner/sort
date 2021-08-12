package sort

type IntArray struct {
	Datas []int
}

func (mi *IntArray) Cover(i, j int) {
	mi.Datas[j] = mi.Datas[i]
}

func (mi *IntArray) LessObj(i int, obj interface{}) bool {
	if mi.Get(i).(int) < obj.(int) {
		return false
	}
	return true
}

func (mi *IntArray) IsDesc() bool {
	if mi.Len() <= 1 {
		return true
	}
	for i := 0; i < mi.Len()-1; i++ {
		if mi.Less(i, i+1) {
			return false
		}
	}
	return true
}

func (mi *IntArray) IsAsc() bool {
	if mi.Len() <= 1 {
		return true
	}
	for i := 0; i < mi.Len()-1; i++ {
		if !mi.Less(i, i+1) {
			return false
		}
	}
	return true
}

func (mi *IntArray) Append(i interface{}) {
	mi.Datas = append(mi.Datas, i.(int))
}

func (mi *IntArray) LessSortableJ(i int, arr Sortable, j int) bool {
	if mi.Datas[i] < (arr.Get(j).(int)) {
		return true
	}
	return false
}

func (mi *IntArray) Slice(s, e int) (r Sortable) {
	return &IntArray{
		Datas: mi.Datas[s:e],
	}
}

func (mi *IntArray) Set(i int, data interface{}) {
	mi.Datas[i] = data.(int)
}

func (mi *IntArray) Len() int {
	return len(mi.Datas)
}

func (mi *IntArray) Less(i, j int) bool {
	return mi.Datas[i] < mi.Datas[j]
}

func (mi *IntArray) Swap(i, j int) {
	t := mi.Datas[i]
	mi.Datas[i] = mi.Datas[j]
	mi.Datas[j] = t
}
func (mi *IntArray) Get(i int) (data interface{}) {
	return mi.Datas[i]
}

type StringArray struct {
	Datas []string
}

func (mi *StringArray) Get(i int) (data interface{}) {
	return mi.Datas[i]
}

func (mi *StringArray) Set(i int, data interface{}) {
	mi.Datas[i] = data.(string)
}

func (mi *StringArray) Len() int {
	return len(mi.Datas)
}

func (mi *StringArray) Less(i, j int) bool {
	return mi.Datas[i] < mi.Datas[j]
}

func (mi *StringArray) Swap(i, j int) {
	t := mi.Datas[i]
	mi.Datas[i] = mi.Datas[j]
	mi.Datas[j] = t
}

func (mi *StringArray) LessSortableJ(i int, arr Sortable, j int) bool {
	if mi.Datas[i] < (arr.Get(j).(string)) {
		return true
	}
	return false
}

func (mi *StringArray) Slice(s, e int) (r Sortable) {
	return &StringArray{
		Datas: mi.Datas[s:e],
	}
}
func (mi *StringArray) Append(i interface{}) {
	mi.Datas = append(mi.Datas, i.(string))
}
func (mi *StringArray) IsDesc() bool {
	if mi.Len() <= 1 {
		return true
	}
	for i := 0; i < mi.Len()-1; i++ {
		if mi.Less(i, i+1) {
			return false
		}
	}
	return true
}

func (mi *StringArray) IsAsc() bool {
	if mi.Len() <= 1 {
		return true
	}
	for i := 0; i < mi.Len()-1; i++ {
		if !mi.Less(i, i+1) {
			return false
		}
	}
	return true
}
func (mi *StringArray) LessObj(i int, obj interface{}) bool {
	if mi.Get(i).(string) < obj.(string) {
		return false
	}
	return true
}
func (mi *StringArray) Cover(i, j int) {
	mi.Datas[j] = mi.Datas[i]
}
