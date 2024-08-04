package main

import (
	"radio"

	"github.com/ddkwork/unison"
	"github.com/ddkwork/unison/app"
)

func main() {
	app.Run("radio", func(w *unison.Window) {
		w.Content().AddChild(radio.New().Layout())
	})
}
