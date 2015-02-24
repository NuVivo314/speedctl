package speedctl

import (
	"testing"
)

func TestByte(t *testing.T) {
	b := Byte(1)
	bs := Byte(10)
	kb := 10 * Kilobyte
	mb := 1 * Megabyte
	gb := 10 * Gigabyte
	tb := 10 * Terabyte
	pb := 10 * Petabyte

	t.Logf("Byte %s and Bytes %s", b, bs)
	t.Logf("Kilobyte %s", kb)
	t.Logf("Megabyte %s", mb)
	t.Logf("Gigabyte %s", gb)
	t.Logf("Terabyte %s", tb)
	t.Logf("Petabyte %s", pb)
}
