package main

import (
	"log"

	"github.com/xmn-services/buckets-network/bundles/gui"
	"github.com/xmn-services/buckets-network/domain/memory/windows"
)

func main() {
	title := "My Window"
	width := uint(800)
	height := uint(600)

	window, err := windows.NewBuilder().Create().WithTitle(title).WithWidth(width).WithHeight(height).Now()
	app := gui.NewOpenglApplication()
	err = app.Execute(window, nil)
	if err != nil {
		log.Println(err.Error())
	}
}
