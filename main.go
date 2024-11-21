package main

import "fmt"

var elapsedTime float64
var exitChan = make(chan bool)

var player Player
var gameMap Map

var showDetails bool

func main() {
  gameMap.init()
  fmt.Println(gameMap)
}
