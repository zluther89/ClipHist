package Clip

import (
	"ClipHist/ClipDB"
	"fmt"
	"testing"
	"time"
)

func TestSave(t *testing.T) {
	testString := "testing Content"
	c := ClipEntry{testString, time.Now().String()}

	c.Save()
	readVal, err := ReadClip()
	if err != nil {
		t.Error("Error from Readlclip", err)
	}

	if testString != readVal {
		t.Error("Clip does not save to clipboard")
	}

	c.Content = "New test contnent"
	c.Save()
	readVal, err = ReadClip()
	if readVal == testString {
		t.Error("Clip does not save to clipboard")
	}
	//reset blank cb
	c = ClipEntry{"testString", time.Now().String()}
	c.Save()
}

func TestStart(t *testing.T) {
	db, _ := ClipDB.Init("../sqliteDB/ClipHist.db")

	ch := Init()
	go ch.Start(db)

	//it should pull from a full channel
	go func() {
		c := ClipEntry{"testing Content", time.Now().String()}
		c.Save()
		tick := time.NewTicker(time.Second * 2)
		saveCount := 1
		for saveCount > 0 {
			select {
			case <-ch.notify:
				saveCount -= 1
				fmt.Println("routine one", saveCount)
			case <-tick.C:
				t.Error("channel did not save the correct number of times")
				return
			}
		}
		tick.Stop()

	}()

	time.Sleep(5 * time.Second)
}

func TestStartTwo(t *testing.T) {
	db, _ := ClipDB.Init("../sqliteDB/ClipHist.db")

	ch := Init()
	go ch.Start(db)

	func() {
		tick := time.NewTicker(time.Second * 10)
		saveCount := 4
		go func() {
			for saveCount > 1 {
				select {
				case <-ch.notify:
					saveCount -= 1
					fmt.Println("routine 3", saveCount)
				case <-tick.C:
					fmt.Print("tick")
					t.Error("channel did not save the correct number of times", saveCount)
					return
				}
			}
		}()
		c := ClipEntry{"A new string", time.Now().String()}
		if err := c.Save(); err != nil {
			fmt.Println(err)
		}

		time.Sleep(2 * time.Second)

		c = ClipEntry{"Another new string", time.Now().String()}
		if err := c.Save(); err != nil {
			fmt.Println(err)
		}

		time.Sleep(2 * time.Second)

		c = ClipEntry{"Another nsdfsew string", time.Now().String()}
		if err := c.Save(); err != nil {
			fmt.Println(err)
		}

		time.Sleep(2 * time.Second)

	}()
	time.Sleep(15 * time.Second)
}
