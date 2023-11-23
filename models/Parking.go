package models

import (
	"sync"
)

type ParkingLot struct {
	Spaces          []*ParkingSpace
	CarQueue        *CarQueue
	mu              sync.Mutex
	CondAvailable   *sync.Cond
}

func NewParkingLot(spaces []*ParkingSpace) *ParkingLot {
	queue := NewCarQueue()
	p := &ParkingLot{
		Spaces:         spaces,
		CarQueue:       queue,
	}
	p.CondAvailable = sync.NewCond(&p.mu)
	return p
}

func (p *ParkingLot) GetSpaces() []*ParkingSpace {
	return p.Spaces
}

func (p *ParkingLot) GetAvailableParkingSpace() *ParkingSpace {
	p.mu.Lock()
	defer p.mu.Unlock()

	for {
		for _, space := range p.Spaces {
			if space.IsAvailable() {
				space.SetAvailability(false)
				return space
			}
		}
		p.CondAvailable.Wait()
	}
}

func (p *ParkingLot) ReleaseParkingSpace(space *ParkingSpace) {
	p.mu.Lock()
	defer p.mu.Unlock()

	space.SetAvailability(true)
	p.CondAvailable.Signal()
}

func (p *ParkingLot) GetCarQueue() *CarQueue {
	return p.CarQueue
}
