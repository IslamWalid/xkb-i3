package notify

import (
	"fmt"
	"os/exec"

	"github.com/IslamWalid/xkb-i3/internal/options"
)

func Notify() error {
	if options.Opts.I3Blocks {
		signal := fmt.Sprintf("-%s", options.Opts.I3BlocksSignal)

        // TODO: use library to signal i3blocks with its name
		err := exec.Command("pkill", signal, "i3blocks").Run()
		if err != nil {
			return err
		}
	}

	return nil
}
