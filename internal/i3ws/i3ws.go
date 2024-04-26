package i3ws

import (
	"github.com/IslamWalid/xkb-i3/internal/evhand"
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
			err = evhand.FocusEventHandler(event.Old.Name, event.Current.Name)
		}

		if err != nil {
			return err
		}
	}

	return nil
}
