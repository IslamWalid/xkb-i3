package options

import "fmt"

const (
	Window Mode = iota
	Workspace
)

type Mode int

func (m *Mode) String() string {
	return string(*m)
}

func (m *Mode) Set(value string) error {
	switch value {
	case "workspace":
		*m = Workspace
	case "window":
		*m = Window
	default:
		return fmt.Errorf("invalid mode value: %s", value)
	}

	return nil
}
