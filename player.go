package main

import (
	"fmt"
	"log"
	"math"

	"github.com/eiannone/keyboard"
)

type Player struct {
	x     float64
	y     float64
	angle float64
}

func (player *Player) init() {
	player.x = 13.0
	player.y = 1.0
	player.angle = 0.0
}

func (player *Player) move() {
	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}

	defer keyboard.Close()
	for {
		char, _, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		switch string(char) {
		case "q":
			player.angle += 3 * elapsedTime
			settings.showDetails = false

		case "e":
			player.angle -= 3 * elapsedTime
			settings.showDetails = false

		case "a": // Move player left (adjust angle)
			settings.showDetails = false
			strafeA := player.angle - math.Pi/2
			player.x -= math.Sin(strafeA) * 5.0 * elapsedTime
			player.y -= math.Cos(strafeA) * 5.0 * elapsedTime

			if string(gameMap.gameMap[int(player.y)*gameMap.width+int(player.x)]) == "#" {
				player.x += math.Sin(strafeA) * 5.0 * elapsedTime
				player.y += math.Cos(strafeA) * 5.0 * elapsedTime
			}

		case "d": // Move player right (adjust angle)
			settings.showDetails = false
			strafeA := player.angle + math.Pi/2
			player.x -= math.Sin(strafeA) * 5.0 * elapsedTime
			player.y -= math.Cos(strafeA) * 5.0 * elapsedTime

			if string(gameMap.gameMap[int(player.y)*gameMap.width+int(player.x)]) == "#" {
				player.x += math.Sin(strafeA) * 5.0 * elapsedTime
				player.y += math.Cos(strafeA) * 5.0 * elapsedTime
			}

		case "w":
			settings.showDetails = false
			player.x += math.Sin(player.angle) * 5.0 * elapsedTime
			player.y += math.Cos(player.angle) * 5.0 * elapsedTime

			if string(gameMap.gameMap[int(player.y)*gameMap.width+int(player.x)]) == "#" {
				player.x -= math.Sin(player.angle) * 5.0 * elapsedTime
				player.y -= math.Cos(player.angle) * 5.0 * elapsedTime
			}

		case "s":
			settings.showDetails = false
			player.x -= math.Sin(player.angle) * 5.0 * elapsedTime
			player.y -= math.Cos(player.angle) * 5.0 * elapsedTime

			if string(gameMap.gameMap[int(player.y)*gameMap.width+int(player.x)]) == "#" {
				player.x += math.Sin(player.angle) * 5.0 * elapsedTime
				player.y += math.Cos(player.angle) * 5.0 * elapsedTime
			}

		case "m":
			settings.showDetails = !settings.showDetails
      if settings.showDetails == true {
        openMapCount++
      }

		case "c": // Exit the game
			fmt.Println("Exiting...")
			exitChan <- true // Signal the main loop to exit
			return
		}
	}
}
