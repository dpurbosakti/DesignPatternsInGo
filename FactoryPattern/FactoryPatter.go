package main

import "fmt"

// logic
type Vehicle interface {
	getVehicleType() string
}

type Car struct {
	Type string
}

type Bike struct {
	Type string
}

func (c *Car) getVehicleType() string {
	return "car"
}

func (c *Bike) getVehicleType() string {
	return "bike"
}

func newCar() *Car {
	return &Car{}
}

func newBike() *Bike {
	return &Bike{}
}

// factory Class/Object
func getVehicleFactory(vehicleType int) {
	var vehicle Vehicle
	if vehicleType == 1 {
		vehicle = newCar()
	} else {
		vehicle = newBike()
	}

	fmt.Println(vehicle.getVehicleType())
}

// client
func main() {
	getVehicleFactory(1)
	getVehicleFactory(2)
}
