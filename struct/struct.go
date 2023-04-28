package speed

// Car implements new car
type Car struct {
	speed        int
	batteryDrain int
	battery      int
	distance     int
}

func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
		distance:     0,
	}
}

// Track is a struct containing distance field
type Track struct {
	distance int
}

// NewTrack creates new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives car one time depending on how much battery is left
func Drive(car Car) Car {
	switch {
	case car.battery > car.batteryDrain:
		return Car{
			speed:        car.speed,
			batteryDrain: car.batteryDrain,
			battery:      car.battery - car.batteryDrain,
			distance:     car.speed + car.distance,
		}
	default:
		return car
	}
}

func CanFinish(car Car, track Track) bool {
	carStatus := (car.battery / car.batteryDrain) * car.speed
	if carStatus >= track.distance {
		return true
	} else {
		return false
	}
}
