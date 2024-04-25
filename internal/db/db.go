package db

var (
	winLang map[int64]string
	wsLang  map[string]string
)

func init() {
	winLang = make(map[int64]string)
	wsLang = make(map[string]string)
}

func GetWindowLang(id int64) (lang string, ok bool) {
	lang, ok = winLang[id]

	return lang, ok
}

func SetWindowLang(id int64, lang string) {
	winLang[id] = lang
}

func DeleteWindowLang(id int64) {
	delete(winLang, id)
}

func GetWorkspaceLang(name string) (lang string, ok bool) {
	lang, ok = wsLang[name]

	return lang, ok
}

func SetWorkspaceLang(name, lang string) {
	wsLang[name] = lang
}
