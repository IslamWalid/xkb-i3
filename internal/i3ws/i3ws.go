package i3ws

import (
	"strconv"

	"github.com/IslamWalid/xkb-i3/internal/evhand"
	"go.i3wm.org/i3/v4"
)

func WorkspaceEventListner() error {
	var oldId, curId string
	var err error

	receiver := i3.Subscribe(i3.WorkspaceEventType)
	defer receiver.Close()

	for receiver.Next() {
		event := receiver.Event().(*i3.WorkspaceEvent)
		switch event.Change {
		case "focus":
			oldId = strconv.FormatInt(int64(event.Old.ID), 10)
			curId = strconv.FormatInt(int64(event.Current.ID), 10)
			err = evhand.FocusEventHandler(oldId, curId)
		}

		if err != nil {
			return err
		}
	}

	return nil
}
