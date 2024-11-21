package main

import (
	"fmt"
	"time"
)

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

	go player.move()

	buffer := make([]string, screen.height*screen.width)
	fmt.Println(buffer)

	fmt.Printf("\x1b[2J\x1b[?25l")
	defer fmt.Printf("\x1b[?25h")

	time_point_1 := time.Now()
	for true {
		time_point_2 := time.Now()
		elapsedTime = time_point_2.Sub(time_point_1).Seconds()
		time_point_1 = time_point_2

		select {
		case <-exitChan:
			fmt.Println("Exiting game loop...")
			return
		default:
			time.Sleep(settings.sleepTime)
		}
	}

}
