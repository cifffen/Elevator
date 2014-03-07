	/*
		case Stop:
			if firstOrderFloor != -1 {
				if firstOrderFloor == prevFloor {
					return Stop
				} else if firstOrderFloor > prevFloor {
					direction = Up
				} else if firstOrderFloor < prevFloor {
					direction = Down
				}
				return direction
			}
	*/
	
	
	
	/*
//Delete given orders at current floor
func deleteFloorOrders(floor int, button int) {
	if locOrdMat[floor][button] == 1 { // Double checking to see that there is an order there. Remove?
		locOrdMat[floor][PanelButton] = 0
		drivers.ElevSetButtonLamp(drivers.TagElevLampType(button), floor, 0)
		orderCount--
	}/*
	switch direction {
	case Up:
		if locOrdMat[floor][0] == 1 {
			locOrdMat[floor][0] = 0
			drivers.ElevSetButtonLamp(drivers.TagElevLampType(0), floor, 0)
			orderCount--
		}
	case Down:
		if locOrdMat[floor][1] == 1 {
			drivers.ElevSetButtonLamp(drivers.TagElevLampType(1), floor, 0)
			locOrdMat[floor][1] = 0
			orderCount--
		}
	} */
}
*/


From getOrder():
					/*
					if j==2{	
						locOrdMat[i][2] = 1
						drivers.ElevSetButtonLamp(drivers.TagElevLampType(2), i, 1)
						orderCount++ // Increment number of active orders.
					} else {
					}
					
					if orderCount == 0 { //set  newOrderEvent if there is made an order to an empty locOrdMat
						firstOrderEvent = true
						firstOrderFloor = i //remember where to first order was made for. Might not be necessary with more elevators.
					}
					*/