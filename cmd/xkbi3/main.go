package main

import (
	"log"

	"github.com/IslamWalid/xkb-i3/internal/i3win"
)

func main() {
	err := i3win.WindowEventHandler()
	if err != nil {
		log.Fatal(err)
	}
}
