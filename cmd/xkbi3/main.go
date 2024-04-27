package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/IslamWalid/xkb-i3/internal/i3win"
	"github.com/IslamWalid/xkb-i3/internal/i3ws"
	"github.com/IslamWalid/xkb-i3/internal/options"
)

func main() {
	var err error

	if options.Opts.Help {
		flag.Usage()
		os.Exit(0)
	}

	switch options.Opts.Mode {
	case options.Workspace:
		err = i3ws.WorkspaceEventListner()

	case options.Window:
		err = i3win.WindowEventListner()
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
