// +build qtpy

package main

import "machine"

// Replace neo in the code below to match the pin
// that you are using if different.
var neo machine.Pin = machine.NEOPIXELS

func init() {
	pwr := machine.NEOPIXELS_POWER
	pwr.Configure(machine.PinConfig{Mode: machine.PinOutput})
	pwr.High()
}
