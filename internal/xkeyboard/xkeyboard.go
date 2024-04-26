package xkeyboard

/*
#cgo pkg-config: x11
#include <X11/XKBlib.h>
*/
import "C"

import "errors"

type XKeyboard struct {
	display *C.Display
	xkbDesc *C.XkbDescRec
}

func New() (xkb XKeyboard, err error) {
	var returnReason C.int

	xkb.display = C.XkbOpenDisplay(nil, nil, nil, nil, nil, &returnReason)
	switch returnReason {
	case C.XkbOD_BadLibraryVersion:
		return xkb, errors.New("bad XKB library version")

	case C.XkbOD_ConnectionRefused:
		return xkb, errors.New("connection refused to X server")

	case C.XkbOD_BadServerVersion:
		return xkb, errors.New("bad X server version")

	case C.XkbOD_NonXkbServer:
		return xkb, errors.New("no XKB server")
	}

	xkb.xkbDesc = C.XkbAllocKeyboard()
	if xkb.xkbDesc == nil {
		return xkb, errors.New("faild to allocate XKB description")
	}
	xkb.xkbDesc.dpy = xkb.display

	return xkb, err
}

func (xkb XKeyboard) GetLayoutIndex() (index int) {
	var state C.XkbStateRec

	C.XkbGetState(xkb.display, C.XkbUseCoreKbd, &state)

	index = int(state.group)

	return index
}

func (xkb XKeyboard) SetLayoutIndex(index int) {
	C.XkbLockGroup(xkb.display, C.XkbUseCoreKbd, C.uint(index))
	C.XFlush(xkb.display)
}

func (xkb XKeyboard) Close() {
	C.XCloseDisplay(xkb.display)
	C.XkbFreeKeyboard(xkb.xkbDesc, C.XkbAllComponentsMask, 1)
}
