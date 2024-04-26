package main

import (
	"fmt"
	"os"

	"github.com/IslamWalid/xkb-i3/internal/i3win"
)

func main() {
	err := i3win.WindowEventListner()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
