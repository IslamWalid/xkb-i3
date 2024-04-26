package i3ws

import (
	"github.com/IslamWalid/xkb-i3/internal/db"
	"github.com/IslamWalid/xkb-i3/internal/xkeyboard"
	"go.i3wm.org/i3/v4"
)

func WorkspaceEventHandler() error {
	var err error

	xkb, err := xkeyboard.New()
	if err != nil {
		return err
	}

	receiver := i3.Subscribe(i3.WorkspaceEventType)
	defer receiver.Close()

	for receiver.Next() {
		event := receiver.Event().(*i3.WorkspaceEvent)
		switch event.Change {
		case "focus":
			err = focusEventHandler(xkb, event)
		default:
			continue
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func focusEventHandler(xkb xkeyboard.XKeyboard, event *i3.WorkspaceEvent) (err error) {
	var index int
	var ok bool

	index = xkb.GetLayoutIndex()
	if err != nil {
		return err
	}

	db.SetWorkspaceLayoutIndex(event.Old.Name, index)

	if index, ok = db.GetWorkspaceLayoutIndex(event.Current.Name); ok {
		xkb.SetLayoutIndex(index)
	}

	return nil
}
