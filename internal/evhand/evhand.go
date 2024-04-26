package evhand

import (
	"github.com/IslamWalid/xkb-i3/internal/database"
	"github.com/IslamWalid/xkb-i3/internal/notify"
	"github.com/IslamWalid/xkb-i3/internal/xkb"
)

func FocusEventHandler(oldId, curId string) error {
	var index int
	var ok bool

	index = xkb.GetLayoutIndex()

	if oldId != "0" {
		database.SetLayoutIndex(oldId, index)
	}

	if index, ok = database.GetLayoutIndex(curId); ok {
		xkb.SetLayoutIndex(index)

		return notify.Notify()
	}

	return nil
}

func CloseEventHandler(id string) {
	database.DeleteLayoutIndex(id)
}
