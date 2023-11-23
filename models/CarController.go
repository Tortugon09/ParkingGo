package models

import "sync"

type CarController struct {
	Cars  []*Car
	Mutex sync.Mutex
}

func NewCarController() *CarController {
	return &CarController{
		Cars: make([]*Car, 0),
	}
}

func (cc *CarController) AddCar(car *Car) {
	cc.Mutex.Lock()
	defer cc.Mutex.Unlock()
	cc.Cars = append(cc.Cars, car)
}

func (cc *CarController) DeleteCar(car *Car) {
	cc.Mutex.Lock()
	defer cc.Mutex.Unlock()
	for i, c := range cc.Cars {
		if c == car {
			cc.Cars = append(cc.Cars[:i], cc.Cars[i+1:]...)
			break
		}
	}
}

func (cc *CarController) GetCars() []*Car {
	cc.Mutex.Lock()
	defer cc.Mutex.Unlock()
	return cc.Cars
}
