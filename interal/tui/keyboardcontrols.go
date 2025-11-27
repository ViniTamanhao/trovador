package tui

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

// setupKeyBindings configures keyboard shortcuts
func (a *App) setupKeyBindings() {
	a.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// Global shortcuts
		switch event.Rune() {
		case 'q':
			a.app.Stop()
			return nil
		case 'r':
			a.refreshPlayers()
			a.updateStatus("Refreshed!")
			return nil
		case '?':
			a.showHelp()
			return nil
		case ' ':
			a.playPause()
			return nil
		case 'n':
			a.next()
			return nil
		case 'p':
			a.previous()
			return nil
		case 's':
			a.stop()
			return nil
		}

		return event
	})
}

// Control actions
func (a *App) playPause() {
	if a.selectedPlayer == nil {
		a.updateStatus("No player selected")
		return
	}

	err := a.controller.PlayPause(a.selectedPlayer.ID)
	if err != nil {
		a.showError("Play/Pause failed: " + err.Error())
		return
	}

	a.updateStatus("⏯  Toggled play/pause")
	
	// Update immediately
	go func() {
		time.Sleep(100 * time.Millisecond)
		a.app.QueueUpdateDraw(func() {
			a.updateTrackInfo()
			a.updatePlayerList()
		})
	}()
}

func (a *App) next() {
	if a.selectedPlayer == nil {
		a.updateStatus("No player selected")
		return
	}

	err := a.controller.Next(a.selectedPlayer.ID)
	if err != nil {
		a.showError("Next failed: " + err.Error())
		return
	}

	a.updateStatus("⏭  Next track")
	
	go func() {
		time.Sleep(300 * time.Millisecond)
		a.app.QueueUpdateDraw(func() {
			a.updateTrackInfo()
		})
	}()
}

func (a *App) previous() {
	if a.selectedPlayer == nil {
		a.updateStatus("No player selected")
		return
	}

	err := a.controller.Previous(a.selectedPlayer.ID)
	if err != nil {
		a.showError("Previous failed: " + err.Error())
		return
	}

	a.updateStatus("⏮  Previous track")
	
	go func() {
		time.Sleep(300 * time.Millisecond)
		a.app.QueueUpdateDraw(func() {
			a.updateTrackInfo()
		})
	}()
}

func (a *App) stop() {
	if a.selectedPlayer == nil {
		a.updateStatus("No player selected")
		return
	}

	err := a.controller.Stop(a.selectedPlayer.ID)
	if err != nil {
		a.showError("Stop failed: " + err.Error())
		return
	}

	a.updateStatus("⏹  Stopped")
	a.updateTrackInfo()
}
