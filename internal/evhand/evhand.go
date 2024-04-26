package evhand

import (
	"github.com/IslamWalid/xkb-i3/internal/db"
	"github.com/IslamWalid/xkb-i3/internal/notify"
	"github.com/IslamWalid/xkb-i3/internal/xkb"
)

func FocusEventHandler(oldID, curID string) error {
	var index int
	var ok bool

	index = xkb.GetLayoutIndex()

	if len(oldID) > 0 {
		db.SetLayoutIndex(oldID, index)
	}

	if index, ok = db.GetLayoutIndex(curID); ok {
		xkb.SetLayoutIndex(index)

		return notify.Notify()
	}

	return nil
}

func CloseEventHandler(curID string) {
	db.DeleteLayoutIndex(curID)
}
