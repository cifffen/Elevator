
package order

import (
	"network"
	
)


func OrderHandler(orderReachedEvent chan<- bool, newOrderEvent chan<- bool, switchDirEvent chan<- bool) {
	var direction 		Direction 	// Keeps the last direction the elevator was heading. Can only be changed in atOrder() and GetDir()
	var prevFloor 		int      	// Holds the previous floor the elevator past. Can only be changed at atOrder()
	var orderCount 		int      	// Keeps track of the number of active orders.
	var firstOrderFloor int 		// Keeps the floor where the first order came from when the elevator was Idle
	var atEndFloor 		bool     	// True if the elevator is at the lowest or highest floor. Used to change direction in case it got "lost"
	var newOrder		bool   	
	// Set high to launch NewOrderEvent if an order is made and the orderMatrix is empyty
	var locOrdMat [Floors][3]int 	// Holds the orders that the elevator has accepted and will carry out
	var activeTenders map[OrderType] 	TenderType
	var lostTenders map[OrderType] 		time.Time
	
	//---- Start Init--------//
	direction 		= Up
	prevFloor 		= floor
	activeTenders 	= make(map[OrderType] TenderType)
	lostTenders 	= make(map[OrderType] time.Time)
	msgChan 	:= make(chan orderMsg)
	orderChan 	:= make(chan orderMsg)
	go network.ListenOnNetwork(msgChan)
	//---- Init complete------//
	for {
		select {
		case <-time.After(time.Millisecond * SamplingTime):
			if newOrders, msgSlice := getOrders(&locOrdMat); newOrders {  // Check for new orders.
				for _, msg := range msgSlice {  // Go through all new orders and process them in msgHandler
					if newOrder:= msgHandler(msg, &locOrdMat); newOrder{
						newOrderEvent <-true  // New order from an empty order matrix has occured
					}
				}
			}
			if orderReached := <- atOrder(); orderReached {
				orderReachedEvent <- true
			}
			
			
			
			
			if currDir <- 
			checkTenderMaps()
			if atEndFloor {
				atEndEvent <- true
			}
			
		case msg:= <-msgChan:  // Received message on the network
			if newOrder := msgHandler(msg, &locOrdMat); newOrder{
				newOrderEvent <-true  // New order from an empty order matrix has occured
			}
		}
	}
}

//Handles orders both locally and over the network
func msgHandler(msg OrderMsg, locOrdMat *[Floors][Buttons] int, aTenders *map[OrderType] TenderType, lTenders *map[OrderType] time.Time )(newOrder bool) {
	newOrder = false
	if checkMsg(msg) {
		order := msg.Order
		floor, button := order.Floor, order.Button
		switch msg.Action {
			case NewOrder:
				if (*locOrdMat)[floor][button] == 0 {
					drivers.ElevSetButtonLamp(drivers.TagElevLampType(button), floor, 1)
					if button == PanelButton {
						(*locOrdMat)[floor][button]=1
						if IsLocOrdMatEmpty(*locOrdMat) {
							newOrder = true
						}
					} else {
						msg.Action = Tender
						msg.TenderVal = cost(floor, button)
						*aTenders[order] = TenderType{time.Now(), msg.TenderVal}
						network.BroadcastOnNet(msg)  // Send tender for order on network
					}
				}			
			case DeleteOrder:
				delete(*aTenders, order)
				delete(*lTenders , order)
				drivers.ElevSetButtonLamp(drivers.TagElevLampType(button), floor, 0)
				if (*locOrdMat)s[floor][button] == 1 {
					(*locOrdMat)[floor][button]=0
				}	
			case Tender:
				if tender, ok := *aTenders[order]; ok { // Check if we already have a tender there
					if tender.val > msg.TenderVal {			// If our tender is worse than the one received -
						delete(*aTenders, order)		// we delete it from active tenders -
						*lTenders[order] = time.Now()	// and add it to lost tenders 
					} 
				} else {
					if tenderVal := cost(floor, button); tenderVal < msg.TenderVal {
						msg.TenderVal = tenderVal
						*aTenders[order] = TenderType{time.Now(), tenderVal}
						network.BroadcastOnNet(msg)  // Send tender for order on network
					}
				}
			case AddOrder:
				delete(*aTenders, order)
				delete(*lTenders , order)
				if (*locOrdMat)[floor][button]  == 0 {
					drivers.ElevSetButtonLamp(drivers.TagElevLampType(button), floor, 1)
					(*locOrdMat)[floor][button] = 1
					if IsLocOrdMatEmpty(*locOrdMat) {
						newOrder = true
					}
				}
		}
	}
	return
}

//Checks that the message is valid
func checkMsg(msg OrderMsg) bool {
	switch msg.Action {
		case NewOrder, DeleteOrder, Tender, AddOrder:
			order 		  := msg.Order
			floor, button := order.Floor, order.Button
			if((floor != 0 && floor != Floors-1) || (floor == 0 && button != DownButton) || (floor == Floors-1 && button != UpButton)){
				if (floor>=0 && floor<Floors) && (button >=0 && button<Buttons) && msg.TenderVal>=0 { 
					return true
				}
			}
	}
	return false
}

func IsLocOrdMatEmpty(locOrdMat [Floors][Buttons] int) bool {
	for i := range locOrdMat {
		for _, order := range locOrdMat[i] {
			if value == 1 {
				return false
			}
		}
	}
	return true
}

// Check for orders
func getOrders(locOrdMat *[Floors][Buttons] int, aTenders map[OrderType] TenderType, lTenders map[OrderType] time.Time )(newOrders bool, orders []OrderMsg ) {
	newOrders = false
	var msg orderMsg
	msg.Action = NewOrder
	for i := range *locOrdMat {
		for j := range (*locOrdMat)[i] {
			if (i != 0 && i != Floors-1) || (i == 0 && j != 1) || (i == Floors-1 && j != 0) { // Statement that makes sure that we don't check the Down button at the groud floor and the Up button at the top floor, as they don't exist.
				if drivers.ElevGetButtonSignal(j, i) == 1 && (*locOrdMat)[i][j] == 0 {
					order = OrderType{j, i}
					_, lostOk   := lTenders[order]; 
					_, activeOk := aTenders[order]; 
					if !lostOk && !activeOk{ 	//Check that those order are not already active on the network, either as an active- or lost tender
						newOrders = true
						msg.Order = order
						orders = append(orders, msg)
					}
				}
			}
		}
	}
	return
}

