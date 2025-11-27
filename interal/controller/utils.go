package controller

import "fmt"

// FormatDuration converts microseconds to human-readable time (mm:ss or hh:mm:ss)
func FormatDuration(microseconds int64) string {
	if microseconds == 0 {
		return "0:00"
	}

	seconds := microseconds / 1000000
	minutes := seconds / 60
	hours := minutes / 60

	seconds = seconds % 60
	minutes = minutes % 60

	if hours > 0 {
		return fmt.Sprintf("%d:%02d:%02d", hours, minutes, seconds)
	}
	return fmt.Sprintf("%d:%02d", minutes, seconds)
}

// FormatProgress returns a progress string with the following format "2:34 / 4:12 (61%)"
func FormatProgress(position, duration int64) string {
	if duration == 0 {
		return FormatDuration(position) + " / --:--"
	}

	percentage := float64(position) / float64(duration) * 100

	return fmt.Sprintf("%s / %s (%.0f%%)", 
		FormatDuration(position),
		FormatDuration(duration),
		percentage,
	)
}

// ProgressBar returns a visual progress bar like "[=====>    ]"
func ProgressBar (position, duration int64, width int) string {
	if duration == 0 || width < 3 {
		return ""
	}

	percentage := float64(position) / float64(duration)
	filled := int(float64(width-2) * percentage)

	bar := "["
	for i := 0; i < width-2; i++ {
		if i < filled {
			bar += "="
		} else if i == filled {
			bar += ">"
		} else {
			bar += " "
		}
	}
	bar += "]"

	return bar
}
