package options

import "flag"

const (
	persistDesc  = `Persist workspace keyboard layout in database (applicable with "workspace" mode only).`
	helpDesc     = `Print usage message.`
	i3blocksDesc = `Set signal to notify i3blocks after changing the keyboard layout.`
	modeDesc     = `Specify the mode: "workspace" or "window" (defaut is "window").`
)

type Options struct {
	Mode           Mode
	Persist        bool // not implemented yet
	Help           bool
	I3Blocks       bool
	I3BlocksSignal string
}

var Opts Options

func init() {
	var help bool
	var i3blocks string
	var mode Mode

	flag.BoolVar(&help, "help", false, helpDesc)
	flag.StringVar(&i3blocks, "i3blocks", "", i3blocksDesc)
	flag.Var(&mode, "mode", modeDesc)

	flag.Parse()

	Opts = Options{
		Mode:           mode,
		Persist:        false,
		Help:           help,
		I3Blocks:       len(i3blocks) > 0,
		I3BlocksSignal: i3blocks,
	}
}
