// Package tui has all the definitions and functionalities available through 'trovador tui'
package tui

import (
	"fmt"
	"time"
	"trovador/interal/controller"

	"github.com/rivo/tview"
)

type App struct {
	// tview application
	app          *tview.Application

	// UI components
	playerList   *tview.List
	trackInfo    *tview.TextView
	progressView *tview.TextView
	controlsHelp *tview.TextView
	statusBar    *tview.TextView
	
	// Application state
	controller   controller.MediaController
	players      []controller.Player
	selectedPlayer *controller.Player

	// Update ticker for live refresh
	updateTicker *time.Ticker
}

func New() (*App, error) {
	ctrl, err := controller.New()
	if err != nil {
		return nil, fmt.Errorf("failed to create controller: %w", err)
	}

	return &App{
		app: tview.NewApplication(),
		playerList: tview.NewList(),
		trackInfo: tview.NewTextView(),
		progressView: tview.NewTextView(),
		controlsHelp: tview.NewTextView(),
		statusBar: tview.NewTextView(),
		controller: ctrl,
		players: []controller.Player{},
		updateTicker: time.NewTicker(1 * time.Second),
	}, nil
}

func Run() error {
	app, err := New()
	if err != nil {
		return err
	}
	
	app.setupUI()
	app.setupKeyBindings()
	app.refreshPlayers()
	app.startBackgroundUpdates()

	defer app.updateTicker.Stop()

	return app.app.Run()
}

func (a *App) setupUI() {
	// playerList declarations
	a.playerList.SetBorder(true)

	a.playerList.ShowSecondaryText(true)
	
	a.playerList.SetTitle(" Players ").SetTitleAlign(tview.AlignLeft)

	// trackInfo declarations
	a.trackInfo.SetBorder(true)

	a.trackInfo.SetDynamicColors(true)

	a.trackInfo.SetTitle(" Player ")

	a.trackInfo.SetText("[gray]No player selected[-]").SetTextAlign(tview.AlignLeft)

	// progressView declarations
	a.progressView.SetBorder(true)

	a.progressView.SetDynamicColors(true)

	a.progressView.SetTitle(" Progress ")

	a.progressView.SetText("[gray]--:-- / --:--[-]").SetTextAlign(tview.AlignLeft)

	// controlsHelp declarations
	a.controlsHelp.SetBorder(true)

	a.controlsHelp.SetDynamicColors(true)

	a.controlsHelp.SetTitle(" Controls ")

	a.controlsHelp.SetText(getControlsText())

	// statusBar declarations
	a.statusBar.SetDynamicColors(true)

	a.statusBar.SetText("[yellow]Press [white::b]?[yellow] for help | [white::b]q[yellow] to quit[-]").SetTextAlign(tview.AlignCenter)

	// Layout: Left panel with players and controls
	leftPanel := tview.NewFlex().SetDirection(tview.FlexRow).AddItem(a.playerList, 0, 2, true).AddItem(a.controlsHelp, 8, 0, false)

	// Layout: Right panel with track info and progress
	rightPanel := tview.NewFlex().SetDirection(tview.FlexRow).AddItem(a.trackInfo, 0, 2, false).AddItem(a.progressView, 8, 0, false)

	// Main layout: Left and right panels side by side 
	mainFlex := tview.NewFlex().AddItem(leftPanel, 0, 1, true).AddItem(rightPanel, 0, 2, false)

	// Root layout: Main content + status bar
	rootFlex := tview.NewFlex().AddItem(mainFlex, 0, 9, true).AddItem(a.statusBar, 0, 1, false)

	a.app.SetRoot(rootFlex, true).SetFocus(a.playerList)
}

// getControlsText returns the help text for controls
func getControlsText() string {
	return `[yellow]Space[-]  Play/Pause
[yellow]n[-]      Next track
[yellow]p[-]      Previous track  
[yellow]s[-]      Stop
[yellow]r[-]      Refresh players
[yellow]?[-]      Help
[yellow]q[-]      Quit`
}

