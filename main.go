package main

import (
	"gioui.org/unit"
	"log"
	"os"

	"gioui.org/app"
)

func main() {
	go func() {
		w := app.NewWindow(func(metric unit.Metric, config *app.Config) {
			config.Size.X = 320
			config.Size.Y = 600
		})
		err := run(w)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
