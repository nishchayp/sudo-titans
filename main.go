package main

import (
	"github.com/nishchayp/sudo-titans/app"
	"github.com/nishchayp/sudo-titans/view"
)

func main() {
	finish := make(chan bool)

	go func() {
		app.Run()
	}()

	go func() {
		view.Run()
	}()

	<-finish
}
