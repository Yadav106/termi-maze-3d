package main

type Map struct {
	height  int
	width   int
	gameMap string
}

func (gameMap *Map) init() {

	gameMap.width = 30
	gameMap.height = 30

	gameMap.gameMap += "##############################"
	gameMap.gameMap += "#............................#"
	gameMap.gameMap += "#............................#"
	gameMap.gameMap += "#............................#"
	gameMap.gameMap += "#............................#"
	gameMap.gameMap += "#............................#"
	gameMap.gameMap += "#............................#"
	gameMap.gameMap += "#............................#"
	gameMap.gameMap += "#............................#"
	gameMap.gameMap += "#............................#"
	gameMap.gameMap += "#............................#"
	gameMap.gameMap += "#............................#"
	gameMap.gameMap += "#............................#"
	gameMap.gameMap += "#............................#"
	gameMap.gameMap += "#............................#"
	gameMap.gameMap += "#............................#"
	gameMap.gameMap += "#............................#"
	gameMap.gameMap += "#............................#"
	gameMap.gameMap += "#............................#"
	gameMap.gameMap += "#............................#"
	gameMap.gameMap += "#............................#"
	gameMap.gameMap += "#............................#"
	gameMap.gameMap += "#............................#"
	gameMap.gameMap += "#............................#"
	gameMap.gameMap += "#............................#"
	gameMap.gameMap += "#............................#"
	gameMap.gameMap += "#............................#"
	gameMap.gameMap += "#............................#"
	gameMap.gameMap += "#............................#"
	gameMap.gameMap += "##############################"
}
