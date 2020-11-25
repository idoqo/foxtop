package main

import (
	"fmt"
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"gitlab.com/idoko/foxtop/db"
)

func main() {

	dbFile := "tmp.sqlite"
	db, err := db.Connect(dbFile)
	if err != nil {
		log.Fatal(err)
	}
	store, err := db.AllHosts()
	if err != nil {
		log.Fatal(err)
	}

	hosts := []string{}
	for _, host := range store.Hosts() {
		host := fmt.Sprintf("%-10d %-15s", host.VisitCount(), host.HostName())
		hosts = append(hosts, host)
	}

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to set up UI: %v", err)
	}
	defer ui.Close()
	list := widgets.NewList()
	list.Title = "Hosts"
	list.Rows = hosts
	list.TextStyle = ui.NewStyle(ui.ColorYellow)
	list.WrapText = false
	list.SetRect(0, 0, 100, 100)
	ui.Render(list)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c":
			return
		case "j", "<Down>":
			list.ScrollDown()
		case "k", "<Up>":
			list.ScrollUp()
		}
		ui.Render(list)
	}
}
