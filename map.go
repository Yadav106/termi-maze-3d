package main

type Map struct {
	height  int
	width   int
	gameMap string
}

func (gameMap *Map) init() {

	gameMap.width = 16
	gameMap.height = 16

	gameMap.gameMap += "################"
	gameMap.gameMap += "#..............#"
	gameMap.gameMap += "#..............#"
	gameMap.gameMap += "#..............#"
	gameMap.gameMap += "#..............#"
	gameMap.gameMap += "#..............#"
	gameMap.gameMap += "#..............#"
	gameMap.gameMap += "#..............#"
	gameMap.gameMap += "#..............#"
	gameMap.gameMap += "#..............#"
	gameMap.gameMap += "#..............#"
	gameMap.gameMap += "#..............#"
	gameMap.gameMap += "#..............#"
	gameMap.gameMap += "#..............#"
	gameMap.gameMap += "#..............#"
	gameMap.gameMap += "################"
}
