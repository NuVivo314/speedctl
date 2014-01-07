package speedctl

// Const value conversion from Byte to Petabyte
const (
  Byteunite	Byte	= 1
  Kilobyte		= 1000
  Megabyte		= Kilobyte * Kilobyte
  Gigabyte		= Megabyte * Kilobyte
  Terabyte		= Gigabyte * Kilobyte
  Petabyte		= Terabyte * Kilobyte
)

// Value use for configuration package
var (
  BuffSize	Byte	= 10 * Kilobyte // Arbitrary initial buff size in no speed control configuration
  BuffStep		= 5  * Kilobyte // Arbitrary step
)
