package loraFieldTesterParser

import "testing"

func TestPresenceOfRssiAndSnrInformation(t *testing.T) {
	status := parseStatus(0x01) //01 00000001
	if !status.PresenceOfRssiAndSnrInformation {
		t.Errorf("Expected PresenceOfRssiAndSnrInformation to be true")
	}
	status = parseStatus(0xfe) //fe 11111110
	if status.PresenceOfRssiAndSnrInformation {
		t.Errorf("Expected PresenceOfRssiAndSnrInformation to be false")
	}
}
func TestPresenceOfBatteryLevelInformation(t *testing.T) {
	status := parseStatus(0x02) //02 00000010
	if !status.PresenceOfBatteryLevelInformation {
		t.Errorf("Expected PresenceOfBatteryLevelInformation to be true")
	}
	status = parseStatus(0xfd) //fd 11111101
	if status.PresenceOfBatteryLevelInformation {
		t.Errorf("Expected PresenceOfBatteryLevelInformation to be false")
	}
}
func TestPresenceOfDownlinkFrameCounter(t *testing.T) {
	status := parseStatus(0x04) //04 00000100
	if !status.PresenceOfDownlinkFrameCounter {
		t.Errorf("Expected PresenceOfDownlinkFrameCounter to be true")
	}
	status = parseStatus(0xfb) //fb 11111011
	if status.PresenceOfDownlinkFrameCounter {
		t.Errorf("Expected PresenceOfDownlinkFrameCounter to be false")
	}
}
func TestPresenceOfUplinkFrameCounter(t *testing.T) {
	status := parseStatus(0x08) //08 00001000
	if !status.PresenceOfUplinkFrameCounter {
		t.Errorf("Expected PresenceOfUplinkFrameCounter to be true")
	}
	status = parseStatus(0xf7) //f7 11110111
	if status.PresenceOfUplinkFrameCounter {
		t.Errorf("Expected PresenceOfUplinkFrameCounter to be false")
	}
}
func TestPresenceOfGpsInformation(t *testing.T) {
	status := parseStatus(0x10) //10 00010000
	if !status.PresenceOfGpsInformation {
		t.Errorf("Expected PresenceOfGpsInformation to be true")
	}
	status = parseStatus(0xef) //ef 11101111
	if status.PresenceOfGpsInformation {
		t.Errorf("Expected PresenceOfGpsInformation to be false")
	}
}
func TestTransmissionTriggeredByPressingPushbutton1(t *testing.T) {
	status := parseStatus(0x20) //20 00100000
	if !status.TransmissionTriggeredByPressingPushbutton1 {
		t.Errorf("Expected TransmissionTriggeredByPressingPushbutton1 to be true")
	}
	status = parseStatus(0xdf) //df 11011111
	if status.TransmissionTriggeredByPressingPushbutton1 {
		t.Errorf("Expected TransmissionTriggeredByPressingPushbutton1 to be false")
	}
}
func TestTransmissionTriggeredByTheAccelerometer(t *testing.T) {
	status := parseStatus(0x40) //40 01000000
	if !status.TransmissionTriggeredByTheAccelerometer {
		t.Errorf("Expected TransmissionTriggeredByTheAccelerometer to be true")
	}
	status = parseStatus(0xbf) //bf 10111111
	if status.TransmissionTriggeredByTheAccelerometer {
		t.Errorf("Expected TransmissionTriggeredByTheAccelerometer to be false")
	}
}
func TestPresenceOfTemperatureInformation(t *testing.T) {
	status := parseStatus(0x80) //80 10000000
	if !status.PresenceOfTemperatureInformation {
		t.Errorf("Expected PresenceOfTemperatureInformation to be true")
	}
	status = parseStatus(0x7f) //7f 01111111
	if status.PresenceOfTemperatureInformation {
		t.Errorf("Expected PresenceOfTemperatureInformation to be false")
	}
}
