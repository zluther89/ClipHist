package Clip

import (
	"ClipHist/ClipDB"
	"fmt"
	"os/exec"
	"time"
)

type ClipEntry struct {
	Content, Timestamp string
}

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

//Saves ClipEntry to clipboard, returns err
func (c *ClipEntry) Save() error {
	writeCmd := getCopyCmd()
	writer, err := writeCmd.StdinPipe()
	if err != nil {
		return err
	}

	if err := writeCmd.Start(); err != nil {
		return err
	}

	if _, err := writer.Write([]byte(c.Content)); err != nil {
		return err
	}

	if err := writer.Close(); err != nil {
		return err
	}
	return writeCmd.Wait()
}

var lastClip string

// Taked in a notification channel, starts a channel to poll Clipboard for changes, write changes to db and notifies alert channel
func ChanStart(notify chan<- bool, db ClipDB.DB) {
	tick := time.NewTicker(time.Second)
	defer tick.Stop()
	for _ = range tick.C {
		if clip := ReadClip(); clip != lastClip && clip != "" {
			lastClip = clip
			db.Write(clip)
			notify <- true
		}
	}
}
