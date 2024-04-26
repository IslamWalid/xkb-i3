package notify

import (
	"fmt"
	"os/exec"

	"github.com/IslamWalid/xkb-i3/internal/options"
)

func Notify() error {
	opts := options.DefaultOpts

	if opts.I3Blocks {
		err := exec.Command("pkill", fmt.Sprint("-%s", opts.I3BlocksSignal), "i3blocks").Run()
		if err != nil {
			return err
		}
	}

	return nil
}
