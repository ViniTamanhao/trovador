# 🎵 Trovador

> A powerful, terminal-based connector and manager for MPRIS media players.

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)

**Trovador** is a CLI and TUI utility written in Go. It acts as a bridge to your system's media players via DBus and MPRIS (Media Player Remote Interfacing Specification). Whether you want to control Spotify, VLC, mpv, or Firefox from a script, or manage them interactively via a terminal UI, Trovador handles the orchestration.

Built with [Cobra](https://github.com/spf13/cobra) for CLI commands and [tview](https://github.com/rivo/tview) for the interactive interface.

---

## ✨ Features

* **Interactive TUI:** A full terminal user interface to visualize and control all active players simultaneously.
* **MPRIS Discovery:** Automatically detects any MPRIS-compatible player running on your system.
* **Granular Control:** Play, pause, stop, and skip tracks for specific players.
* **Script Friendly:** CLI commands allow you to bind Trovador to global hotkeys (e.g., inside i3, Hyprland, or tmux).
* **Lightweight:** Written in pure Go, compiling to a single static binary.

## 📦 Installation

### From Source
Ensure you have Go installed on your machine.

```bash
# Clone the repository
git clone [https://github.com/yourusername/trovador.git](https://github.com/yourusername/trovador.git)

# Enter the directory
cd trovador

# Build the binary
go build -o trovador main.go

# (Optional) Move to your path
sudo mv trovador /usr/local/bin/
