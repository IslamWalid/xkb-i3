package notify

import (
	"fmt"
	"os/exec"

	"github.com/IslamWalid/xkb-i3/internal/options"
)

func Notify() error {
	opts := options.DefaultOpts

	if opts.I3Blocks {
		signal := fmt.Sprintf("-%s", opts.I3BlocksSignal)

		err := exec.Command("pkill", signal, "i3blocks").Run()
		if err != nil {
			return err
		}
	}

	return nil
}
