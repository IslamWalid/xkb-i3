package xkb

import "os/exec"

func CurrentKbLayout() (layout string, err error) {
	data, err := exec.Command("xkb-switch").Output()
	if err != nil {
		return layout, err
	}

	layout = string(data[:len(data)-1])

	return layout, err
}

func SetKbLayout(layout string) (err error) {
	err = exec.Command("xkb-switch", "-s", layout).Run()
	if err != nil {
		return err
	}

	return nil
}
