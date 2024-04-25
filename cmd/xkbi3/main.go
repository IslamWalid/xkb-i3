package main

import (
	"log"

	"github.com/IslamWalid/xkb-i3/internal/i3win"
	"go.i3wm.org/i3/v4"
)

func main() {
	var curFocusID i3.NodeID
	var err error

	recv := i3.Subscribe(i3.WindowEventType)
	defer recv.Close()


	curFocusID, err = i3win.GetFocusedWindowID()
	if err != nil {
        log.Fatal(err)
    }

	for recv.Next() {
		event := recv.Event().(*i3.WindowEvent)

		switch event.Change {
		case "focus":
			err = i3win.WindowFocusEventHandler(curFocusID, event.Container.ID)
			curFocusID = event.Container.ID

		case "close":
			err = i3win.WindowCloseEventHandler(event.Container.ID)
			if event.Container.Focused {
				curFocusID = 0
			}

		default:
			continue
		}

		if err != nil {
            log.Fatal(err)
		}
	}
}
