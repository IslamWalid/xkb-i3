package main

import (
	"log"

	"github.com/IslamWalid/xkb-i3/internal/i3win"
	"github.com/IslamWalid/xkb-i3/internal/xkeyboard"
)

func main() {
	xkb, err := xkeyboard.New()
	if err != nil {
		log.Fatal(err)
	}
	defer xkb.Close()

	err = i3win.WindowEventHandler(xkb)
	if err != nil {
		log.Fatal(err)
	}
}
