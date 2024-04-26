package i3win

import (
	"github.com/IslamWalid/xkb-i3/internal/db"
	"github.com/IslamWalid/xkb-i3/internal/xkb"
	"go.i3wm.org/i3/v4"
)

func WindowEventHandler() error {
	var curFocusID i3.NodeID
	var err error

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
			err = focusEventHandler(curFocusID, event.Container.ID)
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

func focusEventHandler(oldID, curID i3.NodeID) (err error) {
	var lang string
	var ok bool

	lang, err = xkb.CurrentKbLayout()
	if err != nil {
		return err
	}

	if oldID != 0 {
		db.SetWindowLang(int64(oldID), lang)
	}

	if lang, ok = db.GetWindowLang(int64(curID)); ok {
		err = xkb.SetKbLayout(lang)
		if err != nil {
			return err
		}
	}

	return nil
}

func closeEventHandler(curID i3.NodeID) (err error) {
	db.DeleteWindowLang(int64(curID))

	return nil
}
