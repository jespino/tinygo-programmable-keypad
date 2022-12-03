package main

import (
	"machine"
	"machine/usb/hid/keyboard"
	"time"
)

func main() {
	keypad := newDevice(
		[6]machine.Pin{machine.D2, machine.D3, machine.D4, machine.D5, machine.D6, machine.D7},
		[6]machine.Pin{machine.D8, machine.D9, machine.D10, machine.D11, machine.D12, machine.D13},
	)
	keypad.configure()
	for {
		handleKeyPress(keypad)
		time.Sleep(100 * time.Millisecond)
	}
}

func handleKeyPress(keypad *device) error {
	kb := keyboard.New()
	value := keypad.getKey()
	switch value {
	case 2:
		// "+" & backspace & Enter
	case 6:
		// "9"
		openProgram("mattermost")
	case 7:
		// "6"
	case 8:
		// "3"
	case 9:
		// "."
	case 10:
		// "*"
		textPaste()
	case 11:
		// "-"
		writeText("jespinog")
		kb.Down(keyboard.KeyModifierRightAlt)
		kb.Press(keyboard.Key2)
		kb.Up(keyboard.KeyModifierRightAlt)
		writeText("gmail.com")
	case 12:
		// "8"
		openProgram("chromium")
	case 13:
		// "5"
		openProgram("terminal")
	case 14:
		// "2"
	case 15:
		// "0"
	case 16:
		// "/"
		textCut()
	case 18:
		// "7"
		openProgram("firefox")
	case 19:
		// "4"
		openProgram("gimp")
	case 20:
		// "1"
	case 22:
		// Num lock
		textCopy()
	}
	return nil
}

func writeText(text string) {
	kb := keyboard.Keyboard
	kb.Write([]byte(text))
}

func openProgram(programName string) {
	kb := keyboard.Keyboard
	kb.Press(keyboard.KeyLeftGUI)
	time.Sleep(100 * time.Millisecond)
	kb.Write([]byte(programName))
	time.Sleep(100 * time.Millisecond)
	kb.Press(keyboard.KeyEnter)
	time.Sleep(100 * time.Millisecond)
}

func textCopy() {
	kb := keyboard.Keyboard
	kb.Down(keyboard.KeyModifierLeftCtrl)
	kb.Press(keyboard.KeyC)
	kb.Up(keyboard.KeyModifierLeftCtrl)
}

func textCut() {
	kb := keyboard.Keyboard
	kb.Down(keyboard.KeyModifierLeftCtrl)
	kb.Press(keyboard.KeyX)
	kb.Up(keyboard.KeyModifierLeftCtrl)
}

func textPaste() {
	kb := keyboard.Keyboard
	kb.Down(keyboard.KeyModifierLeftCtrl)
	kb.Press(keyboard.KeyV)
	kb.Up(keyboard.KeyModifierLeftCtrl)
}
