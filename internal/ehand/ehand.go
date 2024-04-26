package ehand

import (
	"github.com/IslamWalid/xkb-i3/internal/database"
	"github.com/IslamWalid/xkb-i3/internal/notify"
	"github.com/IslamWalid/xkb-i3/internal/xkeyboard"
)

func FocusEventHandler(xkb xkeyboard.XKeyboard, db database.DB, oldID, curID string) error {
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
