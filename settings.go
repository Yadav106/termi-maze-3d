package main

type Settings struct {
  FOV float64
  depth float64
  showDetails bool
}

func (s *Settings) init() {
  s.FOV = -3.1415 / 4.0
  s.depth = 16
  s.showDetails = false
}
