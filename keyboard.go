// This code here is heavly based on the keypad4x4 driver from
// https://github.com/tinygo-org/drivers/ made by @Nerzal

package main

import (
	"machine"
)

// NoKeyPressed is used, when no key was pressed
const NoKeyPressed = 255

// device is a driver for 4x4 keypads
type device struct {
	inputEnabled bool
	lastColumn   int
	lastRow      int
	columns      [6]machine.Pin
	rows         [6]machine.Pin
	mapping      [6][6]uint8
}

// takes [r6-r1] pins and [c6-c1] pins
func newDevice(rows [6]machine.Pin, columns [6]machine.Pin) *device {
	result := &device{}
	result.columns = columns
	result.rows = rows

	return result
}

// configure sets the column pins as input and the row pins as output
func (keypad *device) configure() {
	inputConfig := machine.PinConfig{Mode: machine.PinInputPullup}
	for i := range keypad.columns {
		keypad.columns[i].Configure(inputConfig)
	}

	outputConfig := machine.PinConfig{Mode: machine.PinOutput}
	for i := range keypad.rows {
		keypad.rows[i].Configure(outputConfig)
		keypad.rows[i].High()
	}

	keypad.mapping = [6][6]uint8{
		{0, 1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10, 11},
		{12, 13, 14, 15, 16, 17},
		{18, 19, 20, 21, 22, 23},
	}

	keypad.inputEnabled = true
	keypad.lastColumn = -1
	keypad.lastRow = -1
}

// getKey returns the code for the given key.
// returns 255 for no keyPressed
func (keypad *device) getKey() uint8 {
	row, column := keypad.getIndices()
	if row == -1 && column == -1 {
		return NoKeyPressed
	}

	return keypad.mapping[row][column]
}

// getIndices returns the position of the pressed key
func (keypad *device) getIndices() (int, int) {
	for rowIndex, rowPin := range keypad.rows {
		rowPin.Low()

		for columnIndex := range keypad.columns {
			columnPin := keypad.columns[columnIndex]

			if !columnPin.Get() && keypad.inputEnabled {
				keypad.inputEnabled = false

				keypad.lastColumn = columnIndex
				keypad.lastRow = rowIndex

				return keypad.lastRow, keypad.lastColumn
			}

			if columnPin.Get() &&
				columnIndex == keypad.lastColumn &&
				rowIndex == keypad.lastRow &&
				!keypad.inputEnabled {
				keypad.inputEnabled = true
			}
		}

		rowPin.High()
	}

	return -1, -1
}
