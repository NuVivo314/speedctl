package speedctl

import (
  "io"
  "log"
  "time"
)

// Struct is send in Update and Done chan
type SpeedInfo struct {
  Size Byte
  Duration time.Duration
}

// Main struct
type Speed struct {
  Update chan SpeedInfo // Send Update information
  Done chan SpeedInfo // Send globale speed/time information

  speed Byte
  speedControl bool
  r io.Reader
  w io.Writer
}

// write is dst
// read is src
// speed is speed of progression in Byte
// Return: Speed "object"
func CopyControl(write io.Reader, read io.Writer, speed Byte) *Speed {
  up := make(chan SpeedInfo, 1000)
  done := make(chan SpeedInfo)

  speedControl := false

  if speed > 0 {
    speedControl = true
  } else {
    speed = BuffStep
  }

  return &Speed {
    Update: up,
    Done: done,
    speed: speed,
    speedControl: speedControl,
    r: read,
    w: write,
  }
}

// Start transfer, return nil or error never io.EOF
// You want to lunch in goroutine !
func (s *Speed) Lunch() error {
  var err error
  var tmpSize int64
  var completSize Byte

  speed := s.speed
  speedAvg := BuffStep

  // Starte "Globale compteur"
  startGlobal := time.Now()

  for {
    // Starte "Local" Compteur
    start := time.Now()

    log.Println("Internal speed:", speed.Kilobyte())
    log.Println("Avrange speed:", speedAvg.Kilobyte())

    tmpSize, err = io.CopyN(s.w, s.r, speed.Byte())
    if err != nil && err != io.EOF {
      log.Println(err.Error())
      break
    }

    completSize += Byte(tmpSize)
    sleep := time.Now().Sub(start)
    s.Update <- SpeedInfo{ Byte(tmpSize), sleep }


    if err == io.EOF {
      break
    }

    speedTmp := SpeedByteSeconds(Byte(tmpSize), sleep)

    // Update Avrange speed. And "low" weighting the new element
    // 1/3 
    speedAvg = ((speedAvg * 2) + speedTmp) / 3

    if s.speedControl {
      if sleep < time.Second && speedAvg >= s.speed {
	speed = s.speed
	time.Sleep(time.Second - sleep)
      }
    } else {
      speed += BuffStep
    }

  }

  if err == nil || err == io.EOF {
    finalGlobal := time.Now().Sub(startGlobal)
    s.Done <- SpeedInfo{ completSize, finalGlobal }

    return nil
  }
  return err
}

