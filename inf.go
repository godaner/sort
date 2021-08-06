package sort

type IntsInterface struct {
	Datas []int
}

func (mi *IntsInterface) Set(i int, data interface{}) {
	mi.Datas[i] = data.(int)
}

func (mi *IntsInterface) Len() int {
	return len(mi.Datas)
}

func (mi *IntsInterface) Less(i, j int) bool {
	return mi.Datas[i] < mi.Datas[j]
}

func (mi *IntsInterface) Swap(i, j int) {
	t := mi.Datas[i]
	mi.Datas[i] = mi.Datas[j]
	mi.Datas[j] = t
}
func (mi *IntsInterface) Get(i int) (data interface{}) {
	return mi.Datas[i]
}

type StrsInterface struct {
	Datas []string
}

func (mi *StrsInterface) Get(i int) (data interface{}) {
	return mi.Datas[i]
}

func (mi *StrsInterface) Set(i int, data interface{}) {
	mi.Datas[i] = data.(string)
}

func (mi *StrsInterface) Len() int {
	return len(mi.Datas)
}

func (mi *StrsInterface) Less(i, j int) bool {
	return mi.Datas[i] < mi.Datas[j]
}

func (mi *StrsInterface) Swap(i, j int) {
	t := mi.Datas[i]
	mi.Datas[i] = mi.Datas[j]
	mi.Datas[j] = t
}
