package orders

import (
      "fmt"
		//"../drivers"
		//"../orders"
		"math"
)

const floorValue 	 = 1 	//Time to travel from one floor to the next
const waitValue      = 3 	//Time spent for each stop
const directionValue = 20 	//Extra cost if the order is in a conflicting direction


func cost(orderFloor int, orderType int, locOrdMat [Floors][Buttons] int, prevFloor int, direction Direction) (cost int) {
	if isLocOrdMatEmpty(locOrdMat){
		cost = getTravelCost(orderFloor, prevFloor)
		fmt.Printf("Cost:%d\n", cost)
		return
	} else{
		cost = getTravelCost(orderFloor, prevFloor)
		cost += getWaitCost(orderFloor, orderType, locOrdMat, prevFloor, direction)
		cost += getDirectionCost(orderFloor, orderType, direction)
		fmt.Printf("Cost:%d\n", cost)
		return
	}
}

func getTravelCost(orderFloor int, prevFloor int) (travelCost int) {
		travelDistance := prevFloor - orderFloor
        travelCost = int(math.Abs(float64(travelDistance))*float64(floorValue))
        return
}

func getWaitCost(orderFloor int, orderType int, locOrdMat [Floors][Buttons] int, prevFloor int, direction Direction) (waitCost int) {
	waitCount := 0
	for i := range locOrdMat{
		for j := range locOrdMat[i]{
			if(locOrdMat[i][j]==1){
				if(direction == Up && j != DownButton && orderFloor > prevFloor && i < orderFloor){
					waitCount++
					break
				} else if(direction == Down && j != UpButton && orderFloor < prevFloor){
					waitCount++
					break
				}
			}
		}
	}
	waitCost = waitCount*waitValue
	return
}

func getDirectionCost(orderFloor int, orderType int, direction Direction) (directionCost int) {
	if((orderType == UpButton && direction == Up)||(orderType == DownButton && direction == Down)){
		directionCost = 0
	} //else if(orderFloor == 
	return
}

