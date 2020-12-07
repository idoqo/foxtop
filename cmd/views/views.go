package views

import (
	"fmt"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"gitlab.com/idoko/foxtop/db"
)

func Run(db db.Database) error {
	if err := ui.Init(); err != nil {
		return err
	}
	defer ui.Close()

	store, err := db.AllHosts()
	if err != nil {
		return err
	}

	hosts := []string{}
	for _, host := range store.Hosts() {
		h := fmt.Sprintf("%-10d %-15s", host.VisitCount, host.HostName())
		hosts = append(hosts, h)
	}

	list := widgets.NewList()
	list.Title = "Hosts"
	list.Rows = hosts
	list.TextStyle = ui.NewStyle(ui.ColorYellow)
	list.WrapText = false
	list.SetRect(0, 0, 100, 100)
	ui.Render(list)

	events := ui.PollEvents()
	for {
		e := <-events
		switch e.ID {
		case "q", "<C-c":
			return nil
		case "j", "<Down>":
			list.ScrollDown()
		case "k", "<Up>":
			list.ScrollUp()
		}
		ui.Render(list)
	}
}
