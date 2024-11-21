package main

var elapsedTime float64
var exitChan = make(chan bool)

var player Player
var gameMap Map
var settings Settings
var screen Screen

func main() {
  gameMap.init()
  settings.init()
  screen.init()
}
