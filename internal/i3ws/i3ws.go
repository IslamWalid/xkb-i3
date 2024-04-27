package i3ws

import (
	"github.com/IslamWalid/xkb-i3/internal/database"
	"github.com/IslamWalid/xkb-i3/internal/notify"
	"github.com/IslamWalid/xkb-i3/internal/xkeyboard"
	"go.i3wm.org/i3/v4"
)

func WorkspaceEventListner() error {
	var err error

	receiver := i3.Subscribe(i3.WorkspaceEventType)
	defer receiver.Close()

	for receiver.Next() {
		event := receiver.Event().(*i3.WorkspaceEvent)
		switch event.Change {
		case "focus":
			err = focusEvent(event.Old.Name, event.Current.Name)

		case "empty":
			emptyEvent(event.Current.Name)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func focusEvent(oldId, curId string) error {
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

func emptyEvent(id string) {
	database.DeleteLayoutIndex(id)
}
