package main

import (
	"TestiCLGO/internal/ascii_art"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	Header string
	Logo   string
}

func NewModel() model {
	return model{
		Header: ascii_art.HEADER,
		Logo:   ascii_art.LOGO,
	}
}

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("TestiCL GO")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model) View() string {
	return m.Logo
}

func main() {
	p := tea.NewProgram(NewModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Couldn't initialize program: %v", err)
		os.Exit(1)
	}
}
