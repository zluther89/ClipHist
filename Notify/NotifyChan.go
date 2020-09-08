package Notify

import "fmt"

func Recieve(r chan bool) {
	for _ = range r {
		fmt.Println("Test Recieve func")
	}
}
