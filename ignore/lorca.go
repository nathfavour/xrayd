package main

import (
	"github.com/zserge/lorca"
	"log"
)

func main() {
	// Create a new UI with a specified load URL and size
	ui, err := lorca.New("", "", 800, 600)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	// Load HTML to the UI
	ui.Load("data:text/html," + `<h1>Hello, World!</h1>`)

	// Wait until the UI window is closed
	<-ui.Done()
}
