package utils

import (
	"github.com/charmbracelet/lipgloss"
)

var Colorpairs = map[int]lipgloss.Style{
	0: lipgloss.NewStyle().
		Background(lipgloss.Color("#000000")).
		Foreground(lipgloss.Color("#ffffff")),
	1: lipgloss.NewStyle().
		Background(lipgloss.Color("#b22222")).
		Foreground(lipgloss.Color("#ffffff")),
	2: lipgloss.NewStyle().
		Background(lipgloss.Color("#0000cd")).
		Foreground(lipgloss.Color("#ffa500")),
	3: lipgloss.NewStyle().
		Background(lipgloss.Color("#8a2be2")).
		Foreground(lipgloss.Color("#ffd700")),
	4: lipgloss.NewStyle().
		Background(lipgloss.Color("#708090")).
		Foreground(lipgloss.Color("#32cd32")),
	5: lipgloss.NewStyle().
		Background(lipgloss.Color("#000000")).
		Foreground(lipgloss.Color("#00ff00")),
	6: lipgloss.NewStyle().
		Background(lipgloss.Color("#ff69b4")).
		Foreground(lipgloss.Color("#fff0f5")),
	7: lipgloss.NewStyle().
		Background(lipgloss.Color("#ff69b4")).
		Foreground(lipgloss.Color("#fff0f5")),
	8: lipgloss.NewStyle().
		Background(lipgloss.Color("#ffd700")).
		Foreground(lipgloss.Color("#000000")),
}
