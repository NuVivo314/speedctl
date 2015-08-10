package speedctl

import (
	"fmt"
	"time"
)

// Byte is abstract for convert. By decimal multiplication(1000) not power of 2 (2^10 or 1024).
type Byte int64

// Convert into byte
func (f Byte) Byte() int64 {
	return int64(f)
}

// Convert into Kilobyte (Unite: kB)
func (f Byte) Kilobyte() int64 {
	return int64(f / Kilobyte)
}

// Convert into Megabyte (Unite: MB)
func (f Byte) Megabyte() int64 {
	return int64(f / Megabyte)
}

// Convert into Gigabyte (Unite: GB)
func (f Byte) Gigabyte() int64 {
	return int64(f / Gigabyte)
}

// Convert into Terabyte (Unite: TB)
func (f Byte) Terabyte() int64 {
	return int64(f / Terabyte)
}

// Convert into Petabyte (Unite: PB)
func (f Byte) Petabyte() int64 {
	return int64(f / Petabyte)
}

func (b Byte) String() string {
	var unit string
	var valu int64

	switch {
	case b == Byteunite && b < Kilobyte:
		unit = "Byte"
		valu = b.Byte()
	case b > Byteunite && b < Kilobyte:
		unit = "Bytes"
		valu = b.Byte()
	case b >= Kilobyte && b < Megabyte:
		unit = "kB"
		valu = b.Kilobyte()
	case b >= Megabyte && b < Gigabyte:
		unit = "MB"
		valu = b.Megabyte()
	case b >= Gigabyte && b < Terabyte:
		unit = "GB"
		valu = b.Gigabyte()
	case b >= Terabyte && b < Petabyte:
		unit = "TB"
		valu = b.Terabyte()
	case b >= Petabyte:
		unit = "PB"
		valu = b.Petabyte()
	}

	return fmt.Sprintf("%d %s", valu, unit)
}

func (b *Byte) ParseByte(parse string) error {
	speed := 0
	unit := ""

	fmt.Sscanf(parse, "%d%s", &speed, &unit)

	if v, ok := convMap[unit]; ok {
		*b = Byte(speed) * v
		return nil
	}

	return UnknowUnit
}

func (b *Byte) UnmarshalJSON(js []byte) error {
	p := string(js[1 : len(js)-1])
	return b.ParseByte(p)
}

func (b *Byte) UnmarshalText(toml []byte) error {
	return b.UnmarshalJSON(toml)
}

func (b *Byte) MarshalJSON() ([]byte, error) {
	return []byte("\"" + b.String() + "\""), nil
}

func (b *Byte) MarshalText() ([]byte, error) {
	return b.MarshalJSON()
}

// Convert Byte/Duration into Byte/second
func BytePerSeconds(bytes Byte, dur time.Duration) Byte {
	f := float64(bytes.Byte()) / dur.Seconds()

	return Byte(f)
}
