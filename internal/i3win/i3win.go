package i3win

import (
	"github.com/IslamWalid/xkb-i3/internal/db"
	"github.com/IslamWalid/xkb-i3/internal/xkeyboard"
	"go.i3wm.org/i3/v4"
)

func WindowEventHandler() error {
	var curFocusID i3.NodeID
	var err error

	xkb, err := xkeyboard.New()
	if err != nil {
		return err
	}

	recv := i3.Subscribe(i3.WindowEventType)
	defer recv.Close()

	curFocusID, err = getFocusedWindowID()
	if err != nil {
		return err
	}

	for recv.Next() {
		event := recv.Event().(*i3.WindowEvent)

		switch event.Change {
		case "focus":
			err = focusEventHandler(xkb, curFocusID, event.Container.ID)
			curFocusID = event.Container.ID

		case "close":
			err = closeEventHandler(event.Container.ID)
			if event.Container.Focused {
				curFocusID = 0
			}

		default:
			continue
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func getFocusedWindowID() (id i3.NodeID, err error) {
	tree, err := i3.GetTree()
	if err != nil {
		return id, err
	}

	focusedWin := tree.Root.FindFocused(func(node *i3.Node) bool {
		return node.Focused && node.Type == i3.Con
	})

	if focusedWin != nil {
		id = focusedWin.ID
	}

	return id, err
}

func focusEventHandler(xkb xkeyboard.XKeyboard, oldID, curID i3.NodeID) (err error) {
	var index int
	var ok bool

	index = xkb.GetLayoutIndex()

	if oldID != 0 {
		db.SetWindowLayoutIndex(int64(oldID), index)
	}

	if index, ok = db.GetWindowLayoutIndex(int64(curID)); ok {
		xkb.SetLayoutIndex(index)
	}

	return nil
}

func closeEventHandler(curID i3.NodeID) (err error) {
	db.DeleteWindowLayoutIndex(int64(curID))

	return nil
}
