package loraFieldTesterParser

const (
	statusPresenceOfRssiAndSnrInformation            = 1 << iota // bit 0
	statusPresenceOfBatteryLevelInformation          = 1 << iota // bit 1
	statusPresenceOfDownlinkFrameCounter             = 1 << iota // bit 2
	statusPresenceOfUplinkFrameCounter               = 1 << iota // bit 3
	statusPresenceOfGpsInformation                   = 1 << iota // bit 4
	statusTransmissionTriggeredByPressingPushbutton1 = 1 << iota // bit 5
	statusTransmissionTriggeredByTheAccelerometer    = 1 << iota // bit 6
	statusPresenceOfTemperatureInformation           = 1 << iota // bit 7
)

func parseStatus(status byte) (loraStatus LoraStatus) {
	loraStatus.PresenceOfRssiAndSnrInformation = status&statusPresenceOfRssiAndSnrInformation > 0
	loraStatus.PresenceOfBatteryLevelInformation = status&statusPresenceOfBatteryLevelInformation > 0
	loraStatus.PresenceOfDownlinkFrameCounter = status&statusPresenceOfDownlinkFrameCounter > 0
	loraStatus.PresenceOfUplinkFrameCounter = status&statusPresenceOfUplinkFrameCounter > 0
	loraStatus.PresenceOfGpsInformation = status&statusPresenceOfGpsInformation > 0
	loraStatus.TransmissionTriggeredByPressingPushbutton1 = status&statusTransmissionTriggeredByPressingPushbutton1 > 0
	loraStatus.TransmissionTriggeredByTheAccelerometer = status&statusTransmissionTriggeredByTheAccelerometer > 0
	loraStatus.PresenceOfTemperatureInformation = status&statusPresenceOfTemperatureInformation > 0
	return
}
