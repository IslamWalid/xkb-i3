package i3ws

import (
	"github.com/IslamWalid/xkb-i3/internal/db"
	"github.com/IslamWalid/xkb-i3/internal/xkeyboard"
	"go.i3wm.org/i3/v4"
)

func WorkspaceEventHandler(xkb xkeyboard.XKeyboard) {
	receiver := i3.Subscribe(i3.WorkspaceEventType)
	defer receiver.Close()

	for receiver.Next() {
		event := receiver.Event().(*i3.WorkspaceEvent)
		switch event.Change {
		case "focus":
			focusEventHandler(xkb, event)
		}

	}
}

func focusEventHandler(xkb xkeyboard.XKeyboard, event *i3.WorkspaceEvent) {
	var index int
	var ok bool

	index = xkb.GetLayoutIndex()

	db.SetWorkspaceLayoutIndex(event.Old.Name, index)

	if index, ok = db.GetWorkspaceLayoutIndex(event.Current.Name); ok {
		xkb.SetLayoutIndex(index)
	}
}
