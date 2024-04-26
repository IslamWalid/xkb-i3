package i3ws

import (
	"github.com/IslamWalid/xkb-i3/internal/db"
	"github.com/IslamWalid/xkb-i3/internal/xkb"
	"go.i3wm.org/i3/v4"
)

func WorkspaceEventHandler() error {
	var err error

	receiver := i3.Subscribe(i3.WorkspaceEventType)
	defer receiver.Close()

	for receiver.Next() {
		event := receiver.Event().(*i3.WorkspaceEvent)
		switch event.Change {
		case "focus":
			err = focusEventHandler(event)
		default:
			continue
		}

		if err != nil {
			return err
		}
	}

    return nil
}

func focusEventHandler(event *i3.WorkspaceEvent) (err error) {
	var lang string
	var ok bool

	lang, err = xkb.CurrentKbLayout()
	if err != nil {
		return err
	}

	db.SetWorkspaceLang(event.Old.Name, lang)

	if lang, ok = db.GetWorkspaceLang(event.Current.Name); ok {
		err = xkb.SetKbLayout(lang)
		if err != nil {
			return err
		}
	}

	return nil
}
