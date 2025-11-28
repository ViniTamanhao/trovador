package tui

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
)

// startBackgroundUpdates starts a goroutine that periodically refreshes data
func (a *App) startBackgroundUpdates() {
	go func() {
		for range a.updateTicker.C {
			a.app.QueueUpdateDraw(func() {
				// Only update if we have a selected player
				if a.selectedPlayer != nil {
					a.updateTrackInfo()
				}
			})
		}
	}()
}

// updateStatus updates the status bar
func (a *App) updateStatus(message string) {
	a.statusBar.SetText(fmt.Sprintf("[yellow]%s[-]", message))
	
	// Reset after 3 seconds
	go func() {
		time.Sleep(3 * time.Second)
		a.app.QueueUpdateDraw(func() {
			a.statusBar.SetText("[yellow]Press [white::b]?[yellow] for help | [white::b]q[yellow] to quit[-]")
		})
	}()
}

// showError shows an error in the status bar
func (a *App) showError(message string) {
	a.statusBar.SetText(fmt.Sprintf("[red]Error: %s[-]", message))
}

// showHelp displays a help modal
func (a *App) showHelp() {
	modal := tview.NewModal().
		SetText(`Trovador - Media Player Controller

Keyboard Shortcuts:
  Space     - Play/Pause selected player
  n         - Next track
  p         - Previous track
  s         - Stop playback
  r         - Refresh player list
  ↑/↓       - Navigate players
  Enter     - Select player
  ?         - Show this help
  q         - Quit

The display updates automatically every second.`).
		AddButtons([]string{"Close"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			// a.app.SetRoot(a, true)
		})

	a.app.SetRoot(modal, true)
}
