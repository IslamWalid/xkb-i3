package evhand

import (
	"github.com/IslamWalid/xkb-i3/internal/database"
	"github.com/IslamWalid/xkb-i3/internal/notify"
	"github.com/IslamWalid/xkb-i3/internal/xkeyboard"
)

func FocusEventHandler(oldId, curId string) error {
	var index int
	var ok bool

	index = xkeyboard.GetLayoutIndex()

	if len(oldId) > 0 {
		database.SetLayoutIndex(oldId, index)
	}

	if index, ok = database.GetLayoutIndex(curId); ok {
		xkeyboard.SetLayoutIndex(index)

		return notify.Notify()
	}

	return nil
}

func CloseEventHandler(id string) {
	database.DeleteLayoutIndex(id)
}
