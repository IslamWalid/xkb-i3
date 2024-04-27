package i3win

import (
	"strconv"

	"github.com/IslamWalid/xkb-i3/internal/database"
	"github.com/IslamWalid/xkb-i3/internal/notify"
	"github.com/IslamWalid/xkb-i3/internal/xkeyboard"
	"go.i3wm.org/i3/v4"
)

func WindowEventListner() error {
	var curId, oldId string

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
			oldId = curId
			curId = strconv.FormatInt(int64(event.Container.ID), 10)
			focusEvent(oldId, curId)

		case "close":
			closeEvent(strconv.FormatInt(int64(event.Container.ID), 10))
			if event.Container.Focused {
				curId = ""
			}
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

func closeEvent(id string) {
	database.DeleteLayoutIndex(id)
}

func getFocusedWindowId() (id string, err error) {
	tree, err := i3.GetTree()
	if err != nil {
		return id, err
	}

	focusedWin := tree.Root.FindFocused(func(node *i3.Node) bool {
		return node.Focused && node.Type == i3.Con
	})

	if focusedWin != nil {
		id = strconv.FormatInt(int64(focusedWin.ID), 10)
	}

	return id, err
}
