package db

var (
	winLayoutIndex map[int64]int
	wsLayoutIndex  map[string]int
)

func init() {
	winLayoutIndex = make(map[int64]int)
	wsLayoutIndex = make(map[string]int)
}

func GetWindowLayoutIndex(id int64) (index int, ok bool) {
	index, ok = winLayoutIndex[id]

	return index, ok
}

func SetWindowLayoutIndex(id int64, index int) {
	winLayoutIndex[id] = index
}

func DeleteWindowLayoutIndex(id int64) {
	delete(winLayoutIndex, id)
}

func GetWorkspaceLayoutIndex(name string) (index int, ok bool) {
	index, ok = wsLayoutIndex[name]

	return index, ok
}

func SetWorkspaceLayoutIndex(name string, index int) {
	wsLayoutIndex[name] = index
}
