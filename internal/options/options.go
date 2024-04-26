package options

const (
	WindowMode Mode = iota
	WorkspaceMode
)

type Mode int

type Options struct {
	Mode           Mode
	I3Blocks       bool
	I3BlocksSignal string
}

var DefaultOpts Options

func init() {
	DefaultOpts = Options{
		Mode:           WindowMode,
		I3Blocks:       true,
		I3BlocksSignal: "SIGRTMIN+12",
	}
}
