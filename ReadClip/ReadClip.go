package ReadClip

import (
	"fmt"
	"os/exec"
)

// returns command to paste info from clipboard
func getReadCmd() *exec.Cmd {
	return exec.Command("pbpaste")
}

// returns string results of paste command
func ReadClip() string {
	pasteCmd := getReadCmd()
	res, err := pasteCmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	return string(res)
}
