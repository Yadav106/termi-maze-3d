package main

type Map struct {
	height  int
	width   int
	gameMap string
}

func (gameMap *Map) init() {

	gameMap.width = 36
	gameMap.height = 36

	gameMap.gameMap += "####################################"
	gameMap.gameMap += "#  #        #s          # #    #   #"
	gameMap.gameMap += "# ## #####  #  ######## # #    #   #"
	gameMap.gameMap += "#  #   #    #         # #   ## # a #"
	gameMap.gameMap += "## #   #    #         # #####  # l #"
	gameMap.gameMap += "#  ### #  ######## #  #        # l #"
	gameMap.gameMap += "# ## # #       # # #  #      # #   #"
	gameMap.gameMap += "#    # #       # # #  ######## # t #"
	gameMap.gameMap += "#### # #   ##  # # #         # # h #"
	gameMap.gameMap += "#    # #   #   # # #         # # e #"
	gameMap.gameMap += "## ### #####   #   ########### #   #"
	gameMap.gameMap += "#    #     #   # ###  #        # b #"
	gameMap.gameMap += "# ####     #   #   #  #        # e #"
	gameMap.gameMap += "#   #  #   #   # # #       ##### s #"
	gameMap.gameMap += "#   #  ### #   # # #    ####   # t #"
	gameMap.gameMap += "# #    #   #   # ####  ##      #   #"
	gameMap.gameMap += "# #    #   #####    #      # ###   #"
	gameMap.gameMap += "# # #  ##      ###  #    ###   #   #"
	gameMap.gameMap += "# # #  #         ## #      #   # p #"
	gameMap.gameMap += "# # # #####       ###########  # r #"
	gameMap.gameMap += "# #####  #   ######            # e #"
	gameMap.gameMap += "#   #    #                   ### s #"
	gameMap.gameMap += "#   #    #    ######  ##  ##   # s #"
	gameMap.gameMap += "#   #  ###    #       #    #   #   #"
	gameMap.gameMap += "#   #    # #  #  ###########   # c #"
	gameMap.gameMap += "#   #      #  #  #   #    #  ###   #"
	gameMap.gameMap += "# #### #####  #  #   #    ##   # t #"
	gameMap.gameMap += "# #    #   #  #  # ######  ##  # o #"
	gameMap.gameMap += "# #   ## ###  ####      #      #   #"
	gameMap.gameMap += "# #  ##          ####   #####  # q #"
	gameMap.gameMap += "#     #                e#      # u #"
	gameMap.gameMap += "################################ i #"
  gameMap.gameMap += "# w : up        # s : down     # t #"
  gameMap.gameMap += "# a : left      # d : right    #   #"
  gameMap.gameMap += "# q : rotate(l) # e : rotate(r)#   #"
	gameMap.gameMap += "####################################"
}
