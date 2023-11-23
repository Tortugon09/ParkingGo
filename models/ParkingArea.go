package models

import (
	"github.com/oakmound/oak/v4/alg/floatgeom"
)

type ParkingSpace struct {
	Area                      *floatgeom.Rect2
	ParkingDirections         *[]ParkingSpaceAddress
	ExitDirections            *[]ParkingSpaceAddress
	Number                    int
	Available                 bool
}

func NewParkingSpace(x, y, x2, y2 float64, row, number int) *ParkingSpace {
	parkingDirections := getParkingDirections(x, y, row)
	exitDirections := getExitDirections()
	area := floatgeom.NewRect2(x, y, x2, y2)

	return &ParkingSpace{
		Area:             &area,
		ParkingDirections: parkingDirections,
		ExitDirections:    exitDirections,
		Number:           number,
		Available:        true,
	}
}

func getParkingDirections(x, y float64, row int) *[]ParkingSpaceAddress {
	var directions []ParkingSpaceAddress

	if row == 1 {
		directions = append(directions, *NewParkingSpaceAddress("left", 445))
	} else if row == 2 {
		directions = append(directions, *NewParkingSpaceAddress("left", 355))
	} else if row == 3 {
		directions = append(directions, *NewParkingSpaceAddress("left", 265))
	} else if row == 4 {
		directions = append(directions, *NewParkingSpaceAddress("left", 175))
	}

	directions = append(directions, *NewParkingSpaceAddress("down", y+5))
	directions = append(directions, *NewParkingSpaceAddress("left", x+5))

	return &directions
}

func getExitDirections() *[]ParkingSpaceAddress {
	var directions []ParkingSpaceAddress

	directions = append(directions, *NewParkingSpaceAddress("down", 450))
	directions = append(directions, *NewParkingSpaceAddress("right", 475))
	directions = append(directions, *NewParkingSpaceAddress("up", 185))

	return &directions
}

func (p *ParkingSpace) GetArea() *floatgeom.Rect2 {
	return p.Area
}

func (p *ParkingSpace) GetNumber() int {
	return p.Number
}

func (p *ParkingSpace) GetParkingDirections() *[]ParkingSpaceAddress {
	return p.ParkingDirections
}

func (p *ParkingSpace) GetExitDirections() *[]ParkingSpaceAddress {
	return p.ExitDirections
}

func (p *ParkingSpace) IsAvailable() bool {
	return p.Available
}

func (p *ParkingSpace) SetAvailability(available bool) {
	p.Available = available
}
