//go:build !nano_33_ble
// +build !nano_33_ble

package hts221

import "tinygo.org/x/drivers"

// New creates a new HTS221 connection. The I2C bus must already be
// configured.
//
// This function only creates the Device object, it does not touch the device.
func New(bus drivers.I2C) Device {
	return Device{bus: bus, Address: HTS221_ADDRESS}
}
