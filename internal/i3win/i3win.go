package i3win

import (
	"strconv"

	"github.com/IslamWalid/xkb-i3/internal/evhand"
	"go.i3wm.org/i3/v4"
)

func WindowEventListner() error {
	var curId, oldId int64

	recv := i3.Subscribe(i3.WindowEventType)
	defer recv.Close()

	curId, err := getFocusedWindowId()
	if err != nil {
		return err
	}

	for recv.Next() {
		event := recv.Event().(*i3.WindowEvent)

		switch event.Change {
		case "focus":
			oldId = int64(curId)
			curId = int64(event.Container.ID)
			evhand.FocusEventHandler(strconv.FormatInt(curId, 10), strconv.FormatInt(oldId, 10))

		case "close":
			evhand.CloseEventHandler(strconv.FormatInt(int64(event.Container.ID), 10))
			if event.Container.Focused {
				curId = 0
			}
		}
	}

	return nil
}

func getFocusedWindowId() (id int64, err error) {
	tree, err := i3.GetTree()
	if err != nil {
		return id, err
	}

	focusedWin := tree.Root.FindFocused(func(node *i3.Node) bool {
		return node.Focused && node.Type == i3.Con
	})

	if focusedWin != nil {
		id = int64(focusedWin.ID)
	}

	return id, err
}
