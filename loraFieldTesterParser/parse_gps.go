package loraFieldTesterParser

import "fmt"

/*
  Coordinates 45°15,9690N
  byte index  00 11 223 3
  last 3 chars are
  9 <- first 4 bits in bytes[3]
  0 <- is always 0
  N <- last bit in bytes[3], 0==N 1==S
*/
func parseLatitude(latitudeBytes []byte) string {
	if len(latitudeBytes) != 4 {
		panic("parseLatitude() only supports 4 bytes as input")
	}
	hemisphere := "N"
	if latitudeBytes[3]&0x01 == 1 {
		hemisphere = "S"
	}
	return fmt.Sprintf(
		"%d%d°%d%d.%d%d%d0%s",
		int8(latitudeBytes[0]&0xf0>>4),
		int8(latitudeBytes[0]&0x0f),
		int8(latitudeBytes[1]&0xf0>>4),
		int8(latitudeBytes[1]&0x0f),
		int8(latitudeBytes[2]&0xf0>>4),
		int8(latitudeBytes[2]&0x0f),
		int8(latitudeBytes[3]&0xf0>>4),
		hemisphere,
	)
}

/*
  Coordinates 005°34,500E
  byte index  001 12 23 3
  last 3 chars are
  9 <- first 4 bits in bytes[3]
  0 <- is always 0
  N <- last bit in bytes[3], E==N W==S
*/
func parseLongitude(longitudeBytes []byte) string {
	if len(longitudeBytes) != 4 {
		panic("parseLongitude() only supports 4 bytes as input")
	}
	hemisphere := "E"
	if longitudeBytes[3]&0x01 == 1 {
		hemisphere = "W"
	}
	return fmt.Sprintf(
		"%d%d%d°%d%d.%d%d0%s",
		int8(longitudeBytes[0]&0xf0>>4),
		int8(longitudeBytes[0]&0x0f),
		int8(longitudeBytes[1]&0xf0>>4),
		int8(longitudeBytes[1]&0x0f),
		int8(longitudeBytes[2]&0xf0>>4),
		int8(longitudeBytes[2]&0x0f),
		int8(longitudeBytes[3]&0xf0>>4),
		hemisphere,
	)
}
