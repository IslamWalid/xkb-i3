package xkb

/*
#cgo pkg-config: x11
#include <X11/XKBlib.h>
*/
import "C"

import (
	"fmt"
	"os"
)

var (
	display *C.Display
	xkbDesc *C.XkbDescRec
)

func init() {
	var returnReason C.int

	display = C.XkbOpenDisplay(nil, nil, nil, nil, nil, &returnReason)
	switch returnReason {
	case C.XkbOD_BadLibraryVersion:
		fmt.Fprintln(os.Stderr, "bad XKB library version")
		os.Exit(1)

	case C.XkbOD_ConnectionRefused:
		fmt.Fprintln(os.Stderr, "connection refused to X server")
		os.Exit(1)

	case C.XkbOD_BadServerVersion:
		fmt.Fprintln(os.Stderr, "bad X server version")
		os.Exit(1)

	case C.XkbOD_NonXkbServer:
		fmt.Fprintln(os.Stderr, "no XKB server")
		os.Exit(1)
	}

	xkbDesc = C.XkbAllocKeyboard()
	if xkbDesc == nil {
		fmt.Fprintln(os.Stderr, "faild to allocate XKB description")
	}
	xkbDesc.dpy = display
}

func GetLayoutIndex() (index int) {
	var state C.XkbStateRec

	C.XkbGetState(display, C.XkbUseCoreKbd, &state)

	index = int(state.group)

	return index
}

func SetLayoutIndex(index int) {
	C.XkbLockGroup(display, C.XkbUseCoreKbd, C.uint(index))
	C.XFlush(display)
}

func Close() {
	C.XCloseDisplay(display)
	C.XkbFreeKeyboard(xkbDesc, C.XkbAllComponentsMask, 1)
}
