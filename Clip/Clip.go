package Clip

import (
	"ClipHist/ClipDB"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os/exec"
	"time"

	"golang.org/x/net/websocket"
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

type ClipChannels struct {
	tick   *time.Ticker
	notify chan bool
}

func Init() ClipChannels {
	tick := time.NewTicker(time.Second)
	notifyTest := make(chan bool, 1)
	return ClipChannels{tick, notifyTest}

}

type message struct {
	Message string `json:"message"`
}

func (c ClipChannels) Listen(ws *websocket.Conn, m string) {
	mes := message{m}
	for _ = range c.notify {
		fmt.Println("read from notify channel")
		if err := websocket.JSON.Send(ws, mes); err != nil {
			log.Println(err)
			break
		}
	}
}

// Takes in a notification channel, starts a channel to poll Clipboard for changes, write changes to db and notifies alert channel
func (c ClipChannels) Start(db ClipDB.DB, lastClip *string) {
	for _ = range c.tick.C {
		clip, err := ReadClip()
		if err != nil {
			log.Fatal(err)
			return
		}
		if clip != *lastClip && clip != "" {
			lastClip = &clip
			db.Write(clip)
			//write to the notification, unless its full
			select {
			case c.notify <- true:
				fmt.Println("Wrote to channel")
			default:
				fmt.Println("Channel is full")

			}

		}
	}

}
