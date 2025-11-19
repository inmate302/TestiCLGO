package gamepad

import (
	"github.com/veandco/go-sdl2/sdl"
)

type EventType int

const (
	EvtAxis EventType = iota
	EvtButton
	EvtHat
	EvtConnected
	EvtDisconnected
	EvtQuit
)

type Event struct {
	Type      EventType
	Which     int
	Axis      uint8
	AxisValue int16
	Button    uint8
	Pressed   bool
	Hat       uint8
	HatPos    uint8
}

func Run(out chan<- Event) error {
	if err := sdl.Init(sdl.INIT_JOYSTICK); err != nil {
		return err
	}
	sdl.JoystickEventState(sdl.ENABLE)

	var joysticks [16]*sdl.Joystick
	defer func() {
		for _, j := range joysticks {
			if j != nil {
				j.Close()
			}
		}
		sdl.Quit()
		close(out)
	}()

	running := true
	for running {
		for ev := sdl.PollEvent(); ev != nil; ev = sdl.PollEvent() {
			switch t := ev.(type) {
			case *sdl.QuitEvent:
				out <- Event{Type: EvtQuit}
				running = false
			case *sdl.JoyAxisEvent:
				out <- Event{Type: EvtAxis, Which: int(t.Which), Axis: t.Axis, AxisValue: t.Value}
			case *sdl.JoyButtonEvent:
				out <- Event{Type: EvtButton, Which: int(t.Which), Button: t.Button, Pressed: t.Type == sdl.JOYBUTTONDOWN}
			case *sdl.JoyHatEvent:
				out <- Event{Type: EvtHat, Which: int(t.Which), Hat: t.Hat, HatPos: t.Value}
			case *sdl.JoyDeviceAddedEvent:
				joysticks[int(t.Which)] = sdl.JoystickOpen(int(t.Which))
				out <- Event{Type: EvtConnected, Which: int(t.Which)}
			case *sdl.JoyDeviceRemovedEvent:
				if j := joysticks[int(t.Which)]; j != nil {
					j.Close()
				}
				out <- Event{Type: EvtDisconnected, Which: int(t.Which)}
			}
		}
		sdl.Delay(16)
	}
	return nil
}
