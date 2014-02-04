
package drivers

/*
#include "elev.h"
*/
import "C"


type TagElevLampType int

const(
	ButtonCallUp TagElevLamType = iota
	ButtonCallDown
	ButtonCommand
)


int elev_init(void);

void elev_set_speed(int speed);

int elev_get_floor_sensor_signal(void);

int elev_get_button_signal(elev_button_type_t button, int floor);

int elev_get_stop_signal(void);

int elev_get_obstruction_signal(void);

void elev_set_floor_indicator(int floor);

void elev_set_button_lamp(elev_button_type_t button, int floor, int value);

void elev_set_stop_lamp(int value);

void elev_set_door_open_lamp(int value);