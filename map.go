package main

type Map struct {
	height  int
	width   int
	gameMap string
}

func (gameMap *Map) init() {

	gameMap.width = 32
	gameMap.height = 32

	gameMap.gameMap += "################################"
	gameMap.gameMap += "#  #        #s          # #    #"
	gameMap.gameMap += "# ## #####  #  ######## # #    #"
	gameMap.gameMap += "#  #   #    #         # #   ## #"
	gameMap.gameMap += "## #   #    #         # #####  #"
	gameMap.gameMap += "#  ### #  ######## #  #        #"
	gameMap.gameMap += "# ## # #       # # #  #      # #"
	gameMap.gameMap += "#    # #       # # #  ######## #"
	gameMap.gameMap += "#### # #   ##  # # #         # #"
	gameMap.gameMap += "#    # #   #   # # #         # #"
	gameMap.gameMap += "## ### #####   #   ########### #"
	gameMap.gameMap += "#    #     #   # ###  #        #"
	gameMap.gameMap += "# ####     #   #   #  #        #"
	gameMap.gameMap += "#   #  #   #   # # #       #####"
	gameMap.gameMap += "#   #  ### #   # # #    ####   #"
	gameMap.gameMap += "# #    #   #   # ####  ##      #"
	gameMap.gameMap += "# #    #   #####    #      # ###"
	gameMap.gameMap += "# # #  ##      ###  #    ###   #"
	gameMap.gameMap += "# # #  #         ## #      #   #"
	gameMap.gameMap += "# # # #####       ###########  #"
	gameMap.gameMap += "# #####  #   ######            #"
	gameMap.gameMap += "#   #    #                   ###"
	gameMap.gameMap += "#   #    #    ######  ##  ##   #"
	gameMap.gameMap += "#   #  ###    #       #    #   #"
	gameMap.gameMap += "#   #    # #  #  ###########   #"
	gameMap.gameMap += "#   #      #  #  #   #    #  ###"
	gameMap.gameMap += "# #### #####  #  #   #    ##   #"
	gameMap.gameMap += "# #    #   #  #  # ######  ##  #"
	gameMap.gameMap += "# #   ## ###  ####      #      #"
	gameMap.gameMap += "# #  ##          ####   #####  #"
	gameMap.gameMap += "#     #                e#      #"
	gameMap.gameMap += "################################"
}
