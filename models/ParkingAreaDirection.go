package models

type ParkingSpaceAddress struct {
	Direction string
	Point     float64
}

func NewParkingSpaceAddress(direction string, point float64) *ParkingSpaceAddress {
	return &ParkingSpaceAddress{
		Direction: direction,
		Point:     point,
	}
}
