package speed

// TODO: define the 'Car' type struct
type Car struct {
	speed        int
	batteryDrain int
	battery      int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	NewCar := Car{speed: speed, batteryDrain: batteryDrain, battery: 100, distance: 0}
	return NewCar
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	newTrack := Track{distance: distance}
	return newTrack
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if car.battery == 100 {
		actualBatery := car.battery - car.batteryDrain
		car.battery = actualBatery
		actualDistance := car.distance + car.speed
		car.distance = actualDistance
	}
	car = Car{speed: car.speed, batteryDrain: car.batteryDrain, battery: car.battery, distance: car.distance}
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	calculateBateryDrain := (((track.distance - car.distance) / (car.speed)) * car.batteryDrain)
	if calculateBateryDrain > car.battery {
		return false
	} else {
		return true
	}
}
