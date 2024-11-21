package main

import "strconv"

func returnWonScreen(openMapCount int) string {
  var won_screen string

  won_screen += "                                                  "
  won_screen += "                                                  "
  won_screen += "   #            #     #########    ###        #   "
  won_screen += "   #            #         #        # ##       #   "
  won_screen += "   #            #         #        #  ##      #   "
  won_screen += "   #            #         #        #   ##     #   "
  won_screen += "   #            #         #        #    ##    #   "
  won_screen += "   #     ##     #         #        #    ##    #   "
  won_screen += "   #     ##     #         #        #     ##   #   "
  won_screen += "   #     ##     #         #        #      ##  #   "
  won_screen += "   #     ##     #         #        #      ##  #   "
  won_screen += "   #     ##     #         #        #       ## #   "
  won_screen += "   #     ##     #         #        #        ###   "
  won_screen += "   ##############     #########    #         ##   "
  won_screen += "                                                  "
  won_screen += "   Map used : "+ strconv.Itoa(openMapCount) +" times                             "
  won_screen += "   Press 'c' to quit                              "

  return won_screen
}
