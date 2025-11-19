package main

import (
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/inmate302/TestiCLGO/internal/ascii_art"
	"github.com/inmate302/TestiCLGO/internal/gamepad"
	"github.com/inmate302/TestiCLGO/internal/utils"
)

type gpMsg gamepad.Event

// hold the channel in the model
type model struct {
	width, height int
	state         ControllerState
	gpCh          <-chan gamepad.Event
	colorIndex    int
}

// dimensions for the controller panel should be 57x17
type ControllerState struct {
	Buttons map[uint8]bool
	Axes    map[uint8]int16
	HatPos  map[uint8]uint8
}

func initialModel(ch <-chan gamepad.Event) model {
	return model{
		state: ControllerState{
			Buttons: map[uint8]bool{},
			Axes:    map[uint8]int16{},
			HatPos:  map[uint8]uint8{},
		},
		gpCh:       ch,
		colorIndex: 0,
	}
}

// subscribe waits for exactly one event, then returns it to Update
func subscribeGamepad(ch <-chan gamepad.Event) tea.Cmd {
	return func() tea.Msg {
		ev, ok := <-ch
		if !ok {
			return nil // channel closed, stop
		}
		return gpMsg(ev)
	}
}

func renderAt(x, y int, s string) string {
	return lipgloss.NewStyle().
		MarginTop(y).
		MarginLeft(x).
		Render(s)
}

func renderControllerASCII(s ControllerState, style lipgloss.Style) string {
	if sdl.NumJoysticks() == 0 {
		return style.Copy().Italic(true).Padding(8, 14).Render("Please, connect a controller...")
	}
	base := utils.NewCanvas(64, 17)
	base.Blit(0, 0, ascii_art.Lines(ascii_art.BASE))
	overlay := utils.NewCanvas(64, 17)
	//btn := lipgloss.NewStyle().Blink(true).Render("██")
	// Buttons (example)
	if s.HatPos[0] == 1 {
		overlay.Blit(8, 4, []string{"███"})
	}
	if s.HatPos[0] == 8 {
		overlay.Blit(4, 6, []string{"███"})
	}
	if s.HatPos[0] == 4 {
		overlay.Blit(8, 8, []string{"███"})
	}
	if s.HatPos[0] == 2 {
		overlay.Blit(12, 6, []string{"███"})
	}

	if s.Buttons[0] {
		overlay.Blit(46, 8, []string{"███"})
	}
	if s.Buttons[1] {
		overlay.Blit(51, 6, []string{"███"})
	}
	if s.Buttons[2] {
		overlay.Blit(41, 6, []string{"███"})
	}
	if s.Buttons[3] {
		overlay.Blit(46, 4, []string{"███"})
	}
	if s.Buttons[4] {
		overlay.Blit(7, 1, ascii_art.Lines(ascii_art.PRESSED))
	} else {
		overlay.Blit(7, 1, ascii_art.Lines(ascii_art.UNPRESSED))
	}
	if s.Buttons[5] {
		overlay.Blit(45, 1, ascii_art.Lines(ascii_art.PRESSED))
	} else {
		overlay.Blit(45, 1, ascii_art.Lines(ascii_art.UNPRESSED))
	}

	if s.Axes[2] > 0 {
		overlay.Blit(7, 0, ascii_art.Lines(ascii_art.PRESSED))
	} else {
		overlay.Blit(7, 0, ascii_art.Lines(ascii_art.UNPRESSED))
	}
	if s.Axes[5] > 0 {
		overlay.Blit(45, 0, ascii_art.Lines(ascii_art.PRESSED))
	} else {
		overlay.Blit(45, 0, ascii_art.Lines(ascii_art.UNPRESSED))
	}

	if s.Buttons[6] {
		overlay.Blit(20, 6, []string{"███"})
	}
	if s.Buttons[7] {
		overlay.Blit(33, 6, []string{"███"})
	}

	// Thumbsticks: anchor positions for LS/RS (tune these)
	lsAx, lsAy := 17, 9 // top-left where STICK’s top-left should be when centered
	rsAx, rsAy := 32, 9

	// Map axis -32768..32767 to small pixel offsets, e.g., [-2..2]
	mapAxis := func(v int16) int {
		f := float64(v) / 32767.0
		if f < -1 {
			f = -1
		}
		if f > 1 {
			f = 1
		}
		return int(f * 2) // -2..2
	}

	lsdx, lsdy := mapAxis(s.Axes[0]), mapAxis(s.Axes[1]) // invert Y
	rsdx, rsdy := mapAxis(s.Axes[3]), mapAxis(s.Axes[4])
	if s.Buttons[9] {
		overlay.Blit(lsAx+lsdx, lsAy+lsdy, ascii_art.Lines(ascii_art.L3R3))
	} else {
		overlay.Blit(lsAx+lsdx, lsAy+lsdy, ascii_art.Lines(ascii_art.STICK))
	}
	if s.Buttons[10] {
		overlay.Blit(rsAx+rsdx, rsAy+rsdy, ascii_art.Lines(ascii_art.L3R3))
	} else {
		overlay.Blit(rsAx+rsdx, rsAy+rsdy, ascii_art.Lines(ascii_art.STICK))
	}
	// Blit sticks onto overlay with transparency for spaces

	// Composite overlay onto base, skipping spaces
	base.BlitTransparent(0, 0, strings.Split(overlay.String(), "\n"), ' ')
	return base.String()
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		tea.SetWindowTitle("TestiCL GO"),
		subscribeGamepad(m.gpCh), // start listening
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		return m, nil

	case gpMsg:
		switch msg.Type {
		case gamepad.EvtButton:
			m.state.Buttons[msg.Button] = msg.Pressed
			if msg.Button == 6 && msg.Pressed {
				m.colorIndex = (m.colorIndex + 1) % len(utils.Colorpairs)
			}
		case gamepad.EvtAxis:
			m.state.Axes[msg.Axis] = msg.AxisValue
		case gamepad.EvtHat:
			m.state.HatPos[msg.Hat] = msg.HatPos
		case gamepad.EvtQuit:
			return m, tea.Quit
		}
		// immediately subscribe for the next event
		return m, subscribeGamepad(m.gpCh)

	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	style := utils.Colorpairs[m.colorIndex]
	ctrl := renderControllerASCII(m.state, style)
	logo := lipgloss.NewStyle().Render(ascii_art.LOGO)
	right := lipgloss.NewStyle().Padding(0, 24).Render(ctrl)
	joy := style.Copy().Padding(7, 33).Render(string(sdl.JoystickNameForIndex(0)))
	MOTD := style.Copy().Italic(true).Padding(4, 2).Render("It's pronounced Testi Cee El. He's chilean you see.")
	return style.Render(lipgloss.JoinHorizontal(lipgloss.Top, logo, right) + lipgloss.JoinHorizontal(lipgloss.Bottom, MOTD, joy))
}

func main() {
	// Create the channel and start SDL once
	ch := make(chan gamepad.Event, 64)
	go func() {
		if err := gamepad.Run(ch); err != nil {
			log.Println("gamepad error:", err)
		}
	}()

	if _, err := tea.NewProgram(initialModel(ch), tea.WithAltScreen()).Run(); err != nil {
		log.Fatal(err)
	}
}
