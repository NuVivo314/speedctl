package speedctl

// Byte is abstract for convert
// TODO: Add (f Byte) String() string look like time.Duration
type Byte int64

// Convert in byte
func (f Byte) Byte() int64 {
  return int64(f)
}

// Convert in Kilobyte (Unite: kB)
func (f Byte) Kilobyte() int64 {
  return int64(f / Kilobyte)
}

// Convert in Megabyte (Unite: MB)
func (f Byte) Megabyte() int64 {
  return int64(f / Megabyte)
}

// Convert in Gigabyte (Unite: GB)
func (f Byte) Gigabyte() int64 {
  return int64(f / Gigabyte)
}

// Convert in Terabyte (Unite: TB)
func (f Byte) Terabyte() int64 {
  return int64(f / Terabyte)
}

// Convert in Petabyte (Unite: PB)
func (f Byte) Petabyte() int64 {
  return int64(f / Petabyte)
}


// Convert Byte/Duration to Byte/second
func BytePerSeconds(bytes Byte, dur time.Duration) Byte {
  f := float64(bytes.Byte()) / dur.Seconds()

  return Byte(f)
}
