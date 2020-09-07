package Notify

import "fmt"

func Recieve(r <-chan bool) {
	for {
		select {
		case <-r:
			fmt.Println("Test Recieve func")

		}
	}
}
