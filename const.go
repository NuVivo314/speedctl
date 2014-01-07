package speedctl

// Const value conversion from Byte to Petabyte
const (
  Byteunite	Byte	= 1
  Kilobyte		= 1000 * Byteunite
  Megabyte		= Kilobyte * Kilobyte
  Gigabyte		= Megabyte * Kilobyte
  Terabyte		= Gigabyte * Kilobyte
  Petabyte		= Terabyte * Kilobyte
)

// Value used for the configuration of this package
var (
  BuffSize	Byte	= 10 * Kilobyte // Arbitrary initial buff size in no speed control configuration
  BuffStep	Byte	= 5  * Kilobyte // Arbitrary step for detect the best speed
)
