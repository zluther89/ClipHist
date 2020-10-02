package clipboard

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"
)

type Entry struct {
	Content, Timestamp string
}

// returns a Clipentry from an json io.Reader
func Decode(r io.Reader) (Entry, error) {
	dec := json.NewDecoder(r)
	var rB Entry
	if err := dec.Decode(&rB); err != nil {
		return Entry{}, err
	}

	return rB, nil
}

type ClipChannels struct {
	tick         *time.Ticker
	notify, quit chan bool
	last         string
}

func Init(t time.Duration) ClipChannels {
	tick := time.NewTicker(t)
	notifyTest := make(chan bool, 1)
	quit := make(chan bool)
	return ClipChannels{tick, notifyTest, quit, ""}

}

type Writer interface {
	Write(s string) (sql.Result, error)
}

// Polls the cliboard for changes, writers changes to a writer and notifies notify chan
func (c ClipChannels) StartListener(w Writer) error {
	for _ = range c.tick.C {
		clip, err := ReadClip()
		if err != nil {
			return err
		}
		tc := strings.TrimSpace(clip)
		if clip != c.last && tc != "" {
			c.last = clip
			_, err := w.Write(clip)
			if err != nil {
				return err
			}
			select {
			case c.notify <- true:
			default:
				fmt.Println("Channel is full")
			}

		}
	}
	return nil

}
