package database

var layoutIndex map[string]int

func init() {
	layoutIndex = make(map[string]int)
}

func GetLayoutIndex(id string) (index int, ok bool) {
	index, ok = layoutIndex[id]

	return index, ok
}

func SetLayoutIndex(id string, index int) {
	layoutIndex[id] = index
}

func DeleteLayoutIndex(id string) {
	delete(layoutIndex, id)
}
