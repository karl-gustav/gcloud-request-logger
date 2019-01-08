package loraFieldTesterParser

import (
	"fmt"
	"log"
)

func ParseFieldTester(fieldTesterBytes []byte) (device LoraFieldTester) {
	device.Source = fmt.Sprintf("%x", fieldTesterBytes)

	device.Status = parseStatus(fieldTesterBytes[0])
	log.Printf("Status %+v", device.Status)
	if device.Status.PresenceOfTemperatureInformation {
		device.Temperature = int8(fieldTesterBytes[1])
	}
	if device.Status.PresenceOfGpsInformation {
		device.Gps.Latitude = parseLatitude(fieldTesterBytes[2:6])
		device.Gps.Longitude = parseLongitude(fieldTesterBytes[6:10])
		switch uint8(fieldTesterBytes[10] & 0xf0 >> 4) {
		case 1:
			device.Gps.ReceptionScale = "Good"
		case 2:
			device.Gps.ReceptionScale = "Average"
		case 3:
			device.Gps.ReceptionScale = "Poor"
		}
		device.Gps.NumberOfSatelites = uint8(fieldTesterBytes[10] & 0x0f)
	}
	if device.Status.PresenceOfUplinkFrameCounter {
		device.UlCounter = uint8(fieldTesterBytes[11])
	}
	if device.Status.PresenceOfDownlinkFrameCounter {
		device.DlCounter = uint8(fieldTesterBytes[12])
	}
	if device.Status.PresenceOfBatteryLevelInformation {
		device.BatteryLevel.MsbValueOfTheBatteryLevelInMillivolt = uint8(fieldTesterBytes[13])
		device.BatteryLevel.LsbValueOfTheBatteryLevelInMillivolt = uint8(fieldTesterBytes[14])
	}
	if device.Status.PresenceOfRssiAndSnrInformation {
		device.Rssi = uint8(fieldTesterBytes[15])
		device.Snr = int8(fieldTesterBytes[16])
	}
	return
}
