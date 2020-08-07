package ReadClip

import (
	"fmt"
	"os/exec"
)

// returns command to paste info from clipboard
func getReadCmd() *exec.Cmd {
	return exec.Command("pbpaste")
}

//returns command to copy to clipboard
func getCopyCmd() *exec.Cmd {
	return exec.Command("pbcopy")
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

//takes in a string, writes string to clipboard, returns err
func WriteToClip(s string) error {
	writeCmd := getCopyCmd()
	writer, err := writeCmd.StdinPipe()
	if err != nil {
		return err
	}

	if err := writeCmd.Start(); err != nil {
		return err
	}

	if _, err := writer.Write([]byte(s)); err != nil {
		return err
	}

	if err := writer.Close(); err != nil {
		return err
	}
	return writeCmd.Wait()
}
