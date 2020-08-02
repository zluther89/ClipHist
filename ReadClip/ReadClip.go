package ReadClip

import "os/exec"

func getReadCmd() *exec.Cmd {
	return exec.Command("pbpaste")
}

func ReadClip() (string, error) {
	pasteCmd := getReadCmd()
	res, err := pasteCmd.Output()
	if err != nil {
		return "", err
	}
	return string(res), nil
}
