package xsens

import "encoding/binary"

// OrientationResetCode is the two-byte value indicating which orientation reset to perform.
// See Table 33 (Available orientation resets) in the Xsens MT documentation.
// To store the present settings, enter Config state and send ResetOrientation again with
// OrientationResetCodeStoreCurrentSettings.
type OrientationResetCode uint16

const (
	// OrientationResetCodeStoreCurrentSettings stores current settings. Only valid in Config state.
	OrientationResetCodeStoreCurrentSettings OrientationResetCode = 0x0000
	// OrientationResetCodeHeadingReset is a heading reset. Not supported by GNSS/INS devices.
	OrientationResetCodeHeadingReset OrientationResetCode = 0x0001
	// OrientationResetCodeReserved is reserved.
	OrientationResetCodeReserved OrientationResetCode = 0x0002
	// OrientationResetCodeObjectInclinationReset is an object or inclination reset.
	OrientationResetCodeObjectInclinationReset OrientationResetCode = 0x0003
	// OrientationResetCodeAlignmentReset is an alignment reset (heading and inclination).
	OrientationResetCodeAlignmentReset OrientationResetCode = 0x0004
	// OrientationResetCodeDefaultHeading sets default heading.
	OrientationResetCodeDefaultHeading OrientationResetCode = 0x0005
	// OrientationResetCodeDefaultInclination sets default inclination.
	OrientationResetCodeDefaultInclination OrientationResetCode = 0x0006
	// OrientationResetCodeDefaultAlignment sets default alignment.
	OrientationResetCodeDefaultAlignment OrientationResetCode = 0x0007
)

// Marshal returns the wire representation of the orientation reset code (2 bytes, big-endian).
func (c OrientationResetCode) Marshal() []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(c))
	return b
}
