package main

import (
	"radio"

	"github.com/ddkwork/app"
	"github.com/richardwilkes/unison"
)

func main() {
	app.Run("radio", func(w *unison.Window) {
		radio.New().Layout(w.Content())
	})
}
