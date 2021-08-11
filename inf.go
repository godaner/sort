package sort

type IntArray struct {
	Datas []int
}

func (mi *IntArray) Append(i interface{}) {
	mi.Datas = append(mi.Datas, i.(int))
}

func (mi *IntArray) LessArray(i int, arr Array, j int) bool {
	if mi.Datas[i] < (arr.Get(j).(int)) {
		return true
	}
	return false
}

func (mi *IntArray) Slice(s, e int) (r Array) {
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

func (mi *StringArray) LessArray(i int, arr Array, j int) bool {
	if mi.Datas[i] < (arr.Get(j).(string)) {
		return true
	}
	return false
}

func (mi *StringArray) Slice(s, e int) (r Array) {
	return &StringArray{
		Datas: mi.Datas[s:e],
	}
}
func (mi *StringArray) Append(i interface{}) {
	mi.Datas = append(mi.Datas, i.(string))
}
