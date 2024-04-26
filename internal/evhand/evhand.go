package evhand

import (
	"github.com/IslamWalid/xkb-i3/internal/db"
	"github.com/IslamWalid/xkb-i3/internal/notify"
	"github.com/IslamWalid/xkb-i3/internal/xkb"
)

func FocusEventHandler(oldId, curId string) error {
	var index int
	var ok bool

	index = xkb.GetLayoutIndex()

	if oldId != "0" {
		db.SetLayoutIndex(oldId, index)
	}

	if index, ok = db.GetLayoutIndex(curId); ok {
		xkb.SetLayoutIndex(index)

		return notify.Notify()
	}

	return nil
}

func CloseEventHandler(id string) {
	db.DeleteLayoutIndex(id)
}
