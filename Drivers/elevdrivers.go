
package drivers

/*
#include "elev.h"
*/
import "C"


type TagElevLampType int

const(
	ButtonCallUp TagElevLampType = iota
	ButtonCallDown
	ButtonCommand
)

func ElevInit() int{
	return int(C.elev_init(void))
}

func ElevSetSpeed(speed int){
	C.elev_set_speed(speed)
}

func ElevGetFloorSensorSignal() int{
	return int(C.elev_get_floor_sensor_signal())
}

func ElevGetButtonSignal(button TagElevLampType, floor int) int{
	return int(C.elev_get_button_signal(button,floor))
}

func ElevGetStopSignal() int{
	return int(C.elev_get_stop_signal())
}

func ElevSetFloorIndicator(floor int){
	C.elev_set_floor_indicator(floor)
}

func ElevSetButtonLamp(button TagElevLampType, floor int, value int){
	C.elev_set_button_lamp(button,floor,value)
}

func ElevSetStopLamp(value int){
	C.elev_set_stop_lamp(value)
}

func ElevSetDoorOpenLamp(value int){
	C.void elev_set_door_open_lamp(value)
}