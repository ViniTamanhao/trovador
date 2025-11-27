package tui

import (
	"fmt"
	"trovador/interal/controller"
)

// refreshPlayers fetches and displays all available players
func (a *App) refreshPlayers() {
	players, err := a.controller.GetPlayers()
	if err != nil {
		a.showError("Failed to get players: " + err.Error())
		return
	}

	a.players = players
	a.updatePlayerList()
}

func (a *App) updatePlayerList() {
	a.playerList.Clear()

	if len(a.players) == 0 {
		a.playerList.AddItem("No players found", "Start a media player", 0, nil)
		return
	}

	for i, player := range a.players {
		index := i

		status, err := a.controller.GetPlayerStatus(player.ID)
		icon := "⏸"
		secondaryText := "Unknown"

		if err == nil {
			switch status.State {
			case "Playing":
				icon = "▶"
				secondaryText = "Playing"
			case "Paused":
				icon = "⏸"
				secondaryText = "Paused"
			case "Stopped":
				icon = "⏹"
				secondaryText = "Stopped"
			}

			if status.Title != "" {
				secondaryText += " - " + status.Title
			}
		}

		displayText := fmt.Sprintf("%s  %s", icon, player.Name)

		a.playerList.AddItem(displayText, secondaryText, 0, func() {
			a.selectPlayer(index)
		})
	}

	if a.selectedPlayer == nil && len(a.players) > 0 {
		a.selectPlayer(0)
	}
}

func (a *App) selectPlayer(index int) {
	if index < 0 || index >= len(a.players) {
		return
	}

	a.selectedPlayer = &a.players[index]
	a.playerList.SetCurrentItem(index)
	a.updateTrackInfo()
	a.updateStatus(fmt.Sprintf("Selected: %s", a.selectedPlayer.Name))
}

// updateTrackInfo fetches and displays current track information
func (a *App) updateTrackInfo() {
	if a.selectedPlayer == nil {
		a.trackInfo.SetText("[gray]No player selected[-]")
		a.progressView.SetText("[gray]--:-- / --:--[-]")
		return
	}

	status, err := a.controller.GetPlayerStatus(a.selectedPlayer.ID)
	if err != nil {
		a.trackInfo.SetText(fmt.Sprintf("[red]Error: %v[-]", err))
		return
	}

	info := fmt.Sprintf("[yellow::b]%s[-::-]\n\n", status.State)

	if status.Title != "" {
		info += fmt.Sprintf("[white::b]Title:[-::-]  %s\n", status.Title)
	} else {
		info += "[gray]Title:  Unknown[-]\n"
	}

	if status.Artist != "" {
		info += fmt.Sprintf("[white::b]Artist:[-::-] %s\n", status.Artist)
	} else {
		info += "[gray]Artist: Unknown[-]\n"
	}

	if status.Album != "" {
		info += fmt.Sprintf("[white::b]Album:[-::-]  %s\n", status.Album)
	} else {
		info += "[gray]Album:  Unknown[-]\n"
	}

	a.trackInfo.SetText(info)

	a.updateProgress(status)
}
//
// func (a *App) updateProgress(status controller.PlayerStatus) {
// 	if status.Duration == 0 {
// 		if status.Position > 0 {
// 			a.progressView.SetText(fmt.Sprintf("[yellow]%s[-]", 
// 				controller.FormatDuration(status.Position)))
// 		} else {
// 			a.progressView.SetText("[gray]--:-- / --:--[-]")
// 		}
// 		return
// 	}
//
// 	timeText := fmt.Sprintf("[yellow]%s[-]", 
// 		controller.FormatProgress(status.Position, status.Duration))
//
// 	progressBar := controller.ProgressBar(status.Position, status.Duration, 30)
//
// 	a.progressView.SetText(fmt.Sprintf("%s\n\n[cyan]%s[-]", timeText, progressBar))
// }

func (a *App) updateProgress(status controller.PlayerStatus) {
    // GetRect may return 0 before first draw -> fallback width
    _, _, width, _ := a.progressView.GetRect()
    if width < 15 {
        width = 30 // fallback default
    }

    barWidth := width - 2
    if barWidth < 5 {
        barWidth = 5
    }

    var text string

    if status.Duration == 0 {
        if status.Position > 0 {
            text = fmt.Sprintf("[yellow]%s[-]",
                controller.FormatDuration(status.Position))
        } else {
            text = "[gray]--:-- / --:--[-]"
        }
    } else {
        timeText := fmt.Sprintf("[yellow]%s[-]",
            controller.FormatProgress(status.Position, status.Duration))

        progressBar := controller.ProgressBar(status.Position, status.Duration, barWidth)

        text = fmt.Sprintf("%s\n\n[cyan]%s[-]", timeText, progressBar)
    }

    a.progressView.SetText(text)
}
