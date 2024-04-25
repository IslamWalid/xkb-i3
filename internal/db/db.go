package db

var langMap map[int64]string

func init() {
	langMap = make(map[int64]string)
}

func GetWindowLang(id int64) (lang string, ok bool) {
	lang, ok = langMap[id]

	return lang, ok
}

func SetWindowLang(id int64, lang string) {
	langMap[id] = lang
}

func DeleteWindow(id int64) {
	delete(langMap, id)
}
