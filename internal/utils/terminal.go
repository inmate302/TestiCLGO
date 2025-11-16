package utils

import (
	"TestiCLGO/internal/ascii_art"
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// -------------------------------------------------------------------
// Model
// -------------------------------------------------------------------
type model struct {
	artLines []string // all lines of the picture
	current  []string // lines that have been shown so far
	index    int      // next line to reveal
}

// -------------------------------------------------------------------
// Init – start a ticker that drives the animation
// -------------------------------------------------------------------
func (m model) Init() tea.Cmd {
	// Tick every 100 ms (adjust speed as you like)
	return tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
		return tickMsg{}
	})
}

// -------------------------------------------------------------------
// Message type for the ticker
// -------------------------------------------------------------------
type tickMsg struct{}

// -------------------------------------------------------------------
// Update – reveal the next line on each tick
// -------------------------------------------------------------------
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		if m.index < len(m.artLines) {
			// Append the next line
			m.current = append(m.current, m.artLines[m.index])
			m.index++
			// Continue ticking until the art is complete
			return m, tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
				return tickMsg{}
			})
		}
		// No more lines – stop ticking
		return m, nil
	}
	return m, nil
}

// -------------------------------------------------------------------
// View – render the lines that have been revealed so far
// -------------------------------------------------------------------
func (m model) View() string {
	return fmt.Sprintf("%s\n", joinLines(m.current))
}

// Helper to join slice of strings with newlines
func joinLines(lines []string) string {
	return fmt.Sprint(strings.Join(lines, "\n"))
}

// -------------------------------------------------------------------
// Main – define the ASCII art and start the program
// -------------------------------------------------------------------
func main() {
	art := ascii_art.LOGO

	p := tea.NewProgram(model{artLines: art})
	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
