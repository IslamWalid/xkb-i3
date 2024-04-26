package db

type DB struct {
	layoutIndex map[string]int
}

func New() DB {
	return DB{layoutIndex: make(map[string]int)}
}

func (db DB) GetLayoutIndex(id string) (index int, ok bool) {
	index, ok = db.layoutIndex[id]

	return index, ok
}

func (db DB) SetLayoutIndex(id string, index int) {
	db.layoutIndex[id] = index
}

func (db DB) DeleteLayoutIndex(id string) {
	delete(db.layoutIndex, id)
}
