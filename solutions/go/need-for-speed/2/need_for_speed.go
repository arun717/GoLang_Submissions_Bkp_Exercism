package speed

type Car struct {
    battery int
    batteryDrain int
    speed int
    distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car {
        battery : 100,
        speed : speed,
        batteryDrain: batteryDrain,
        distance: 0,
    }
}

type Track struct {
    distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track {
        distance : distance,
    }
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if car.battery - car.batteryDrain >= 0 {
        return Car {
            speed : car.speed,
            batteryDrain : car.batteryDrain,
            battery : car.battery - car.batteryDrain,
            distance : car.distance + car.speed,
    	}
    } else {
        return Car {
            speed : car.speed ,
            batteryDrain : car.batteryDrain ,
            battery : car.battery,
            distance : car.distance,
        }
    }
    
	
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	distanceRem := track.distance
    batteryRem := car.battery
    for distanceRem >=0 && batteryRem >= car.batteryDrain {
        if batteryRem - car.batteryDrain >= 0 {
            distanceRem = distanceRem - car.speed
            batteryRem = batteryRem - car.batteryDrain
        }
    }
    if distanceRem > 0 {
        return false
    } else {
        return true
    }
}
