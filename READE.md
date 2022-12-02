# Tinygo Programmable KeyPad

This project is a small idea to create my own programmable keypad using tinygo
and cheap hardware.

I used a cheap keypad (3.5€) and a [Waveshare
RP2040-Zero](https://www.waveshare.com/wiki/RP2040-Zero) board (6€) and a FPC
Connector. Total budge has been around 10€.

The way it works, is using the `machine/usb/hid/keyboard` library from tinygo
to simulate a keyboard and convert the key pressed in the keypad to a sequence
of key strokes that we program. For example, the code in this repo use the
gnome-shell for run programs like firefox, chromium or gimp.

## How I did it

The keypad included its own pcb with its own controller, but the keys where
connected with FPC connector, so I remove the entire original board and replace
it with the RP2040-zero connecting it with the FPC connector with 12 pins.

After that I created the code that you can find in this repo (Also I have to
add support for the RP2040-Zero board to TinyGo for this). As part of the code
I created the "driver" for the keypad heavily based on the existing keypad4x4
driver of tinygo.
