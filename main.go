package main

import (
	"TestiCLGO/internal/ascii_art"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	Logo   string
	Header string
}

func MainScreen() model {
	return model{
		Logo:   lipgloss.NewStyle().Margin(1, 2).Render(ascii_art.LOGO),
		Header: lipgloss.NewStyle().Margin(1, 20).Render(ascii_art.HEADER),
	}
}

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("TestiCL GO")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Left, m.Logo, m.Header)
}

func main() {
	p := tea.NewProgram(MainScreen(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Couldn't initialize program: %v", err)
		os.Exit(1)
	}
}
