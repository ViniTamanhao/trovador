package cmd

import "trovador/interal/controller"

// findPlayer searches for a player by name
func findPlayer(ctrl controller.MediaController, name string) *controller.Player {
	players, err := ctrl.GetPlayers()
	if err != nil {
		return nil
	}

	for _, p := range players {
		if p.Name == name {
			return &p
		}
	}
	return nil
}
