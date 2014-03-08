package fsm

import (
	"../drivers"
	"../orders"
	"fmt"
	"time"
)

const brakeDur = 10   //Duration, in milliseconds, of the braking time when stopping at a floor
const doorOpenDur = 3 //Duration, in seconds, of the time the door stays open when arriving at a floor
const Speed = 300     //The speed of the motor

type (
	Event int
	State int
)

const (
	OrderReached Event = iota
	TimerFinished
	NewOrder
	SwitchDirection
)
const (
	Idle State = iota
	Running
	AtFloor
)

var state State
var DoorTimer <-chan time.Time
var BrakeTimer <-chan time.Time
var direction orders.Direction

func InitElev() int {
	if drivers.ElevInit() == 0 { //IO init failed
		return 0
	} else {
		direction = orders.Down
		if drivers.ElevGetFloorSensorSignal() != -1 { //Check if the elevator is at a floor
		} else { //else, run downwards until one is found
				drivers.ElevSetSpeed(int(direction) * Speed)
				floor := drivers.ElevGetFloorSensorSignal()
				for floor == -1 {
					floor = drivers.ElevGetFloorSensorSignal()
				}
				drivers.ElevSetSpeed(int(-1*direction) * Speed)
				brake()
		}
		//orders.InitOrderMod(floor)
		state = Idle
		fmt.Printf("Initialized\n")
		return 1
	}
}

//Reverse the direction to brake
func brake() {
	BrakeTimer = time.After(time.Millisecond * brakeDur)
}

// Checks for events and runs the state machine when some occur
func EventManager() {
	orderReachedEvent := make(chan bool)
	newOrderEvent 	  := make(chan bool)
	switchDirEvent 	  := make(chan bool)
	atEndEvent := make(chan bool)
	go orders.CheckForEvents(orderReachedEvent, newOrderEvent, switchDirEvent, atEndEvent)
	for {
		select {
		case <-BrakeTimer:
			drivers.ElevSetSpeed(int(orders.Stop))
			fmt.Printf("Ferdig \n")
		case <-newOrderEvent:
			fmt.Printf("New order event\n")
			stateMachine(NewOrder)
		case direction:= <-switchDirEvent:
			fmt.Printf("Switch direction event\n")
		case <-atEndEvent:
			stateMachine(AtEndFloor)
		case <-orderReachedEvent:
			fmt.Printf("Order reached event\n")
			stateMachine(OrderReached)
		case <-DoorTimer:
			fmt.Printf("Door timer finished\n")
			stateMachine(TimerFinished)
		}
	}
}

func stateMachine(event Event) {
	switch state {
	case Idle:
		switch event {
		case NewOrder:
			if direction != 0 {
				drivers.ElevSetSpeed(int(direction) * Speed)
				state = Running
			} else {
				DoorTimer = time.After(time.Second * doorOpenDur)
				drivers.ElevSetDoorOpenLamp(1)
				state = AtFloor
			}
		}
	case Running:
		switch event {
		case AtEndFloor:
			drivers.ElevSetSpeed(int(direction) * Speed)
		case OrderReached:
			drivers.ElevSetSpeed(-1 * int(direction) * Speed)
			brake()
			DoorTimer = time.After(time.Second * doorOpenDur)
			drivers.ElevSetDoorOpenLamp(1)
			state = AtFloor
			fmt.Printf("Atfloor \n")
		}
	case AtFloor:
		switch event {

		case OrderReached:
			fmt.Printf("At floor again\n")
			state = AtFloor
		case TimerFinished:
			if orders.IsLocOrdMatEmpty() {
				drivers.ElevSetDoorOpenLamp(0)
				state = Idle
				fmt.Printf("Idle \n")
			} else if direction == orders.Stop {
				DoorTimer = time.After(time.Second * doorOpenDur)
			} else {
				drivers.ElevSetDoorOpenLamp(0)
				state = Running
				fmt.Printf("Runing \n")
				drivers.ElevSetSpeed(int(direction) * Speed)
			}

		}
	}
}
