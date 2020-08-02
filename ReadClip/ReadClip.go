package ReadClip

import (
	"fmt"
	"os/exec"
)

func getReadCmd() *exec.Cmd {
	return exec.Command("pbpaste")
}

func ReadClip() string {
	pasteCmd := getReadCmd()
	res, err := pasteCmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	return string(res)
}
