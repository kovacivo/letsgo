package main

import "fmt"

type car struct {
	gas_pedal uint16     //min: 0,      max: 65535    16bit
	brake_pedal uint16   //min: 0,      max: 65535
	steering_wheel int16 //min: -32768  max: 32768
	top_speed_kmh float64 //what's our top speed?
	}

	const usixteenbitmax float64 = 1234
	const kmh_multiple float64 = 1.60934

func main() {
	//a_car := car{gas_pedal: 16535, brake_pedal: 0, steering_wheel: 12562, top_speed_kmh: 225.0}	

	//a_car := car{22314,0,12562,225.0}
	//fmt.Println("gas_pedal:",a_car.gas_pedal, "brake_pedal:",a_car.brake_pedal,"steering_wheel:",a_car.steering_wheel)
	//fmt.Println("kmh func: ", a_car.kmh())
	//fmt.Println("mph func: ", a_car.mph())
	//fmt.Println("Setting new top speed: ")
	//a_car.set_new_top_speed(45)
	//fmt.Println("kmh func: ", a_car.kmh())
	//fmt.Println("mph func: ", a_car.mph())

	a_car := car{22314,0,12562,225.0}
	fmt.Println("gas_pedal:",a_car.gas_pedal, "brake_pedal:",a_car.brake_pedal,"steering_wheel:",a_car.steering_wheel)
	fmt.Println("Car is going",a_car.mph(),"MPH,",a_car.kmh(),"KMH, and top speed is",a_car.top_speed_kmh)
	a_car.set_new_top_speed(500)
	fmt.Println("Car is going",a_car.mph(),"MPH,",a_car.kmh(),"KMH, and top speed is",a_car.top_speed_kmh)

}

func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax)
}

func (c car) mph() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh/kmh_multiple/usixteenbitmax)
}

func (c *car) set_new_top_speed(newspeed float64) {
	c.top_speed_kmh = newspeed
}
