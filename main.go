package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)

var elapsedTime float64
var exitChan = make(chan bool)

var player Player
var gameMap Map
var settings Settings
var screen Screen

var openMapCount int = 0
var won bool = false

var called_loop bool = false
func loop_screen() {
	// show looping animation
	player.x = 15.0
	player.y = 30.6
	player.angle = 1.60

  for {
    player.x += 0.01
    if player.x > 20.0 {
      player.x = 15.0
    }
    time.Sleep(settings.sleepTime)
  }
}

func main() {
	gameMap.init()
	settings.init()
	screen.init()
	player.init()

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
			fmt.Printf("\x1b[H")
			for x := 0; x < screen.width; x++ {
				fRayAngle := (player.angle - settings.FOV/2.0) + (float64(x)/float64(screen.width))*settings.FOV
				fDistanceToWall := 0.0
				bHitWall := false
				bBoundary := false

				if int(player.y) == 30 && int(player.x) == 23 {
					won = true
				}

				var fEyeX float64 = math.Sin(fRayAngle)
				var fEyeY float64 = math.Cos(fRayAngle)

				for !bHitWall && (fDistanceToWall < settings.depth) {
					fDistanceToWall += 0.1

					var nTestX int = int(player.x + fEyeX*fDistanceToWall)
					var nTestY int = int(player.y + fEyeY*fDistanceToWall)

					// Test if ray is out of bounds
					if nTestX < 0 || nTestX > gameMap.width || nTestY < 0 || nTestY >= gameMap.height {
						bHitWall = true // set the distance to max depth
						fDistanceToWall = settings.depth
					} else {
						// Ray is in bounds, test to see if ray cell is wall block
						if string(gameMap.gameMap[nTestY*gameMap.width+nTestX]) == string("#") {
							bHitWall = true

							pairs := []Pair{}

							for tx := 0; tx < 2; tx++ {
								for ty := 0; ty < 2; ty++ {
									vy := float64(nTestY) + float64(ty) - player.y
									vx := float64(nTestX) + float64(tx) - player.x
									d := math.Sqrt(vx*vx + vy*vy)
									dot := (fEyeX * vx / d) + (fEyeY * vy / d)
									pairs = append(pairs, Pair{
										first:  d,
										second: dot,
									})
								}
							}

							sort.Slice(pairs, func(i, j int) bool {
								return pairs[i].first < pairs[j].first
							})

							fBound := 0.01
							if math.Acos(pairs[0].second) < fBound {
								bBoundary = true
							}
							if math.Acos(pairs[1].second) < fBound {
								bBoundary = true
							}
						}
					}
				}

				// Calculate distance to ceiling and floor
				var nCeiling int = int(float64(screen.height/2.0) - float64(screen.height)/float64(fDistanceToWall))
				var nFloor int = screen.height - nCeiling

				var shade rune = ' '

				if fDistanceToWall <= settings.depth/4.0 {
					shade = 0x2588
				} else if fDistanceToWall <= settings.depth/3.0 {
					shade = 0x2593
				} else if fDistanceToWall <= settings.depth/2.0 {
					shade = 0x2592
				} else if fDistanceToWall <= settings.depth {
					shade = 0x2591
					// shade = rune('.')
				} else {
					shade = ' '
				}

				if bBoundary {
					shade = rune(' ')
				}

				for y := 0; y < screen.height; y++ {
					if y < nCeiling {
						buffer[y*screen.width+x] = " "
					} else if y > nCeiling && y < nFloor {
						buffer[y*screen.width+x] = string(shade)
					} else {
						b := 1.0 - ((float64(y) - float64(screen.height)/2.0) / (float64(screen.height) / 2.0))
						fShade := " "
						if b < 0.25 {
							fShade = "#"
						} else if b < 0.5 {
							fShade = "x"
						} else if b < 0.75 {
							fShade = "."
						} else if b < 0.9 {
							fShade = "-"
						} else {
							fShade = " "
						}
						buffer[y*screen.width+x] = fShade
					}
				}

			}

			if !won {
				if settings.showDetails {
					// Display Stats
					formattedString := fmt.Sprintf("Map opened %d times", openMapCount)
					for i, char := range formattedString {
						if i < screen.width { // Ensure we don't exceed the screen width
							buffer[i] = string(char)
						}
					}

					// Display Map
					for nx := 0; nx < gameMap.width; nx++ {
						for ny := 0; ny < gameMap.width; ny++ {
							buffer[(ny+1)*screen.width+nx] = string(gameMap.gameMap[ny*gameMap.width+nx])
						}
					}

					// Display Player
					buffer[(int(player.y)+1)*screen.width+int(player.x)] = "P"
				}
			} else {
				// display win screen in center
        won_screen := returnWonScreen(openMapCount)
				offset_x := screen.width/2 - 50/2
				offset_y := screen.width * (screen.height/2 - 8)
				for y := 0; y < 17; y++ {
					for x := 0; x < 50; x++ {
						buffer[offset_y+offset_x+y*screen.width+x] = string(won_screen[y*50+x])
					}
				}

        // show looping screen at the end
        if !called_loop {
          go loop_screen()
          called_loop = true
        }

			}

			// Output the buffer to the screen
			for y := 0; y < screen.height; y++ {
				for x := 0; x < screen.width; x++ {
					fmt.Printf("%s", buffer[y*screen.width+x])
				}
				fmt.Println()
			}

			time.Sleep(settings.sleepTime)
		}
	}

}
