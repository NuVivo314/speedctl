package speedctl

import (
	"os"
	"testing"
)

const (
	null = "/dev/null"
	zero = "/dev/zero"
)

func TestDevNullSpeed10Kb(t *testing.T) {
	testZero(zero, null, 10*Kilobyte, t)
}

func TestDevNullSpeed100Kb(t *testing.T) {
	testZero(zero, null, 100*Kilobyte, t)
}

func TestDevNullSpeed10MB(t *testing.T) {
	testZero(zero, null, 10*Megabyte, t)
}

func TestDevNullSpeedUnlimite(t *testing.T) {
	testZero(zero, null, 0, t)
}

func testZero(src string, dst string, speed Byte, t *testing.T) {
	dstFile, err := os.OpenFile(dst, os.O_WRONLY, 0)
	if err != nil {
		t.Fatal(err.Error())
		return
	}

	defer dstFile.Close()

	srcFile, err := os.Open(src)
	if err != nil {
		t.Fatal(err.Error())
		return
	}

	defer srcFile.Close()

	sc := CopyControl(dstFile, srcFile, speed)

	go func() {
		err := sc.Copy()
		if err != nil {
			t.Error(err.Error())
		}
	}()

	i := 0
	var speedInfo Byte
LOOP:
	for {
		select {
		case update := <-sc.Update:
			t.Logf("Update: %s/s", BytePerSeconds(update.Size, update.Duration))
			i += 1
			continue
		case done := <-sc.Done:
			speedInfo = BytePerSeconds(done.Size, done.Duration)
			t.Logf("Done: %s/s", speedInfo)
			break LOOP
		default:
			if i >= 10 {
				sc.Done <- SpeedInfo{}
			}
		}
	}
	speedMax := speed + (speed * 10 / 100)
	speedMin := speed - (speed * 10 / 100)
	if speed > 0 && (speedInfo > speedMax || speedInfo < speedMin) {
		t.Errorf("Speed control fail: Max[%d]/Min[%d] allow Current: %d", speedMax, speedMin, speedInfo)
	} else if speed > 0 {
		t.Logf("Speed control done: Max[%s]/Min[%s] allow Current: %s", speedMax, speedMin, speedInfo)
	} else {
		t.Logf("Speed is %s/s", speedInfo)
	}
}
