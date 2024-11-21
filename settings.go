package main

import "time"

type Settings struct {
  FOV float64
  depth float64
  showDetails bool
  sleepTime time.Duration
}

func (s *Settings) init() {
  s.FOV = -3.1415 / 4.0
  s.depth = 16
  s.showDetails = false
  s.sleepTime = time.Microsecond * 8000 * 2
}
