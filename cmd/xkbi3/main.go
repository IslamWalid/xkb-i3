package main

import (
	"fmt"
	"os"

	"github.com/IslamWalid/xkb-i3/internal/i3win"
	"github.com/IslamWalid/xkb-i3/internal/i3ws"
	"github.com/IslamWalid/xkb-i3/internal/options"
)

func main() {
	var err error

	opts := options.DefaultOpts

	switch opts.Mode {
	case options.WorkspaceMode:
		err = i3ws.WorkspaceEventListner()

	case options.WindowMode:
		err = i3win.WindowEventListner()
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
