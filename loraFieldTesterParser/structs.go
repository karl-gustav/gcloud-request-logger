package loraFieldTesterParser

type LoraFieldTester struct {
	Status       LoraStatus `json:"status"`
	Temperature  int8       `json:"temperature"`
	Gps          Gps        `json:"gps"`
	GpsQuality   bool       `json:"gpsQuality"`
	UlCounter    uint8      `json:"ulCounter"`
	DlCounter    uint8      `json:"dlCounter"`
	BatteryLevel Battery    `json:"batteryLevel"`
	Rssi         uint8      `json:"rssi"`
	Snr          int8       `json:"snr"`
	Source       string     `json:"source"`
}

type LoraStatus struct {
	PresenceOfRssiAndSnrInformation            bool `json:"presenceOfRssiAndSnrInformation"`
	PresenceOfBatteryLevelInformation          bool `json:"presenceOfBatteryLevelInformation"`
	PresenceOfDownlinkFrameCounter             bool `json:"presenceOfDownlinkFrameCounter"`
	PresenceOfUplinkFrameCounter               bool `json:"presenceOfUplinkFrameCounter"`
	PresenceOfGpsInformation                   bool `json:"presenceOfGpsInformation"`
	TransmissionTriggeredByPressingPushbutton1 bool `json:"transmissionTriggeredByPressingPushbutton1"`
	TransmissionTriggeredByTheAccelerometer    bool `json:"transmissionTriggeredByTheAccelerometer"`
	PresenceOfTemperatureInformation           bool `json:"presenceOfTemperatureInformation"`
}

type Gps struct {
	Latitude          string `json:"latitude"`
	Longitude         string `json:"longitude"`
	ReceptionScale    string `json:"receptionScale"`
	NumberOfSatelites uint8  `json:"numberOfSatelites"`
}

type Battery struct {
	MsbValueOfTheBatteryLevelInMillivolt uint8 `json:"msbValueOfTheBatteryLevelInMillivolt"`
	LsbValueOfTheBatteryLevelInMillivolt uint8 `json:"lsbValueOfTheBatteryLevelInMillivolt"`
}
