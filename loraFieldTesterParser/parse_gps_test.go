package loraFieldTesterParser

import "testing"

func TestParseLatitudeInputValidation(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Errorf("Expected parseLatitude() to panic because byte array was wrong size")
		}
	}()
	parseLatitude([]byte{0, 0, 0})
}

func TestParseLatitudeManualExampleCoordinatesN(t *testing.T) {
	coordinateBytes := []byte{0x45, 0x15, 0x96, 0x90}
	coordinateString := "45°15.9690N"

	parsedCoordinates := parseLatitude(coordinateBytes)
	if parsedCoordinates != coordinateString {
		t.Errorf(
			"Expected coordinates to be %s but was %s",
			coordinateString,
			parsedCoordinates,
		)
	}
}

func TestParseLatitudeManualExampleCoordinatesS(t *testing.T) {
	coordinateBytes := []byte{0x19, 0x50, 0x91, 0x61}
	coordinateString := "19°50.9160S"

	parsedCoordinates := parseLatitude(coordinateBytes)
	if parsedCoordinates != coordinateString {
		t.Errorf(
			"Expected coordinates to be %s but was %s",
			coordinateString,
			parsedCoordinates,
		)
	}
}

func TestParseLongitudeInputValidation(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Errorf("Expected parseLongitude() to panic because byte array was wrong size")
		}
	}()
	parseLongitude([]byte{0, 0, 0})
}

func TestParseLongitudeManualExampleCoordinatesE(t *testing.T) {
	coordinateBytes := []byte{0x00, 0x53, 0x45, 0x00}
	coordinateString := "005°34.500E"

	parsedCoordinates := parseLongitude(coordinateBytes)
	if parsedCoordinates != coordinateString {
		t.Errorf(
			"Expected coordinates to be %s but was %s",
			coordinateString,
			parsedCoordinates,
		)
	}
}

func TestParseLongitudeManualExampleCoordinatesW(t *testing.T) {
	coordinateBytes := []byte{0x01, 0x82, 0x80, 0x51}
	coordinateString := "018°28.050W"

	parsedCoordinates := parseLongitude(coordinateBytes)
	if parsedCoordinates != coordinateString {
		t.Errorf(
			"Expected coordinates to be %s but was %s",
			coordinateString,
			parsedCoordinates,
		)
	}
}
