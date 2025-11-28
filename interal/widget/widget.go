// Package widget defines everything related to the widget command
package widget

import (
	"fmt"
	"time"
	"trovador/interal/controller"

	"github.com/rivo/tview"
)

type Widget struct {
	app 				   *tview.Application
	mainView       *tview.TextView
	controller     controller.MediaController
	selectedPlayer *controller.Player
	updateTicker   *time.Ticker
}

func New() (*Widget, error) {
	ctrl, err := controller.New()
	if err != nil {
		return nil, fmt.Errorf("failed to create controller: %w", err)
	}

	return &Widget{
		app:          tview.NewApplication(),
		mainView:     tview.NewTextView(),
		controller:   ctrl,
		updateTicker: time.NewTicker(1 * time.Second),
	}, nil
}

func Run() error {
	w, err := New()
	if err != nil {
		return err
	}

	return w.app.Run()
}

func (w *Widget) setupUI() {
	// mainView declarations
	w.mainView.SetBorder(true)

	w.mainView.SetDynamicColors(true)

	w.mainView.SetTitle(" Trovador ").SetTitleAlign(tview.AlignCenter)

	w.mainView.SetText("[gray]Loading...[-]").SetTextAlign(tview.AlignCenter)

	// Main flex definition
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().AddItem(nil, 0, 1, false).AddItem(w.mainView, 45, 1, true).AddItem(nil, 0, 1, false), 10, 1, true).
		AddItem(nil, 0, 1, false)

	w.app.SetRoot(flex, true).SetFocus(w.mainView)
}
