package Clip

import (
	"ClipHist/ClipDB"
	"encoding/json"
	"fmt"
	"io"
	"log"
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

// returns command to copy to clipboard
func getCopyCmd() *exec.Cmd {
	return exec.Command("pbcopy")
}

// returns a Clipentry from an json io.Reader
func DecodeClip(r io.Reader) (ClipEntry, error) {
	dec := json.NewDecoder(r)
	var rB ClipEntry
	if e := dec.Decode(&rB); e != nil {
		return ClipEntry{}, e
	}

	return rB, nil
}

// returns string results of paste command
func ReadClip() (string, error) {
	pasteCmd := getReadCmd()
	res, err := pasteCmd.Output()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(res), nil
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

var lastClip string = ""

// Takes in a notification channel, starts a channel to poll Clipboard for changes, write changes to db and notifies alert channel
func ChanStart(notify chan<- bool, db ClipDB.DB) {
	tick := time.NewTicker(time.Second)
	defer tick.Stop()
	for _ = range tick.C {
		clip, err := ReadClip()
		if err != nil {
			log.Fatal("ChanStart error:", err)
			return
		}
		if clip != lastClip && clip != "" {
			lastClip = clip
			db.Write(clip)
			notify <- true
		}
	}
	return
}
