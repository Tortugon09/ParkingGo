package models

import (
	"image/color"
	"sync"
	"time"

	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/entities"
	"github.com/oakmound/oak/v4/scene"
	"github.com/oakmound/oak/v4/render"
)

const (
	entrada   = 185.00
	velocidad = 10
)

const carImagePath = "assets/car6.png"

type Car struct {
	area    floatgeom.Rect2
	entity  *entities.Entity
	mu      sync.Mutex
}

func NewCar(ctx *scene.Context) *Car {
	area := floatgeom.NewRect2(445, -20, 465, 0)
	sprite, _ := render.LoadSprite(carImagePath)
	entity := entities.New(ctx,
		entities.WithRect(area),
		entities.WithColor(color.RGBA{255, 0, 0, 255}),
		entities.WithRenderable(sprite),
		entities.WithDrawLayers([]int{1, 2}),
	)

	return &Car{
		area:   area,
		entity: entity,
	}
}

func (c *Car) Enqueue(controller *CarController) {
	for c.Y() < 145 {
		if !c.Collision("down", controller.GetCars()) {
			c.MoveY(1)
			time.Sleep(velocidad * time.Millisecond)
		}
	}
}

func (c *Car) JoinGate(controller *CarController) {
	for c.Y() < entrada {
		if !c.Collision("down", controller.GetCars()) {
			c.MoveY(1)
			time.Sleep(velocidad * time.Millisecond)
		}
	}
}

func (c *Car) LeaveGate(controller *CarController) {
	for c.Y() > 145 {
		if !c.Collision("up", controller.GetCars()) {
			c.MoveY(-1)
			time.Sleep(velocidad * time.Millisecond)
		}
	}
}

func (c *Car) Park(space *ParkingSpace, controller *CarController) {
	for index := 0; index < len(*space.GetParkingDirections()); index++ {
		directions := *space.GetParkingDirections()
		switch directions[index].Direction {
		case "right":
			for c.X() < directions[index].Point {
				if !c.Collision("right", controller.GetCars()) {
					c.MoveX(1)
					time.Sleep(velocidad * time.Millisecond)
				}
			}
		case "down":
			for c.Y() < directions[index].Point {
				if !c.Collision("down", controller.GetCars()) {
					c.MoveY(1)
					time.Sleep(velocidad * time.Millisecond)
				}
			}
		case "left":
			for c.X() > directions[index].Point {
				if !c.Collision("left", controller.GetCars()) {
					c.MoveX(-1)
					time.Sleep(velocidad * time.Millisecond)
				}
			}
		case "up":
			for c.Y() > directions[index].Point {
				if !c.Collision("up", controller.GetCars()) {
					c.MoveY(-1)
					time.Sleep(velocidad * time.Millisecond)
				}
			}
		}
	}
}

func (c *Car) Leave(space *ParkingSpace, controller *CarController) {
	for index := 0; index < len(*space.GetExitDirections()); index++ {
		directions := *space.GetExitDirections()
		switch directions[index].Direction {
		case "left":
			for c.X() > directions[index].Point {
				if !c.Collision("left", controller.GetCars()) {
					c.MoveX(-1)
					time.Sleep(velocidad * time.Millisecond)
				}
			}
		case "right":
			for c.X() < directions[index].Point {
				if !c.Collision("right", controller.GetCars()) {
					c.MoveX(1)
					time.Sleep(velocidad * time.Millisecond)
				}
			}
		case "up":
			for c.Y() > directions[index].Point {
				if !c.Collision("up", controller.GetCars()) {
					c.MoveY(-1)
					time.Sleep(velocidad * time.Millisecond)
				}
			}
		case "down":
			for c.Y() < directions[index].Point {
				if !c.Collision("down", controller.GetCars()) {
					c.MoveY(1)
					time.Sleep(velocidad * time.Millisecond)
				}
			}
		}
	}
}

func (c *Car) LeaveSpace(controller *CarController) {
	initialX := c.X()
	for c.X() > initialX-30 {
		if !c.Collision("left", controller.GetCars()) {
			c.MoveX(-1)
			time.Sleep(20 * time.Millisecond)
		}
	}
}

func (c *Car) MoveAway(controller *CarController) {
	for c.Y() > -20 {
		if !c.Collision("up", controller.GetCars()) {
			c.MoveY(-1)
			time.Sleep(velocidad * time.Millisecond)
		}
	}
}

func (c *Car) MoveY(dy float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entity.ShiftY(dy)
}

func (c *Car) MoveX(dx float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entity.ShiftX(dx)
}

func (c *Car) X() float64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.entity.X()
}

func (c *Car) Y() float64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.entity.Y()
}

func (c *Car) Destroy() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entity.Destroy()
}

func (c *Car) Collision(direction string, cars []*Car) bool {
	minimumDistance := 30.0
	for _, car := range cars {
		switch direction {
		case "left":
			if c.X() > car.X() && c.X()-car.X() < minimumDistance && c.Y() == car.Y() {
				return true
			}
		case "right":
			if c.X() < car.X() && car.X()-c.X() < minimumDistance && c.Y() == car.Y() {
				return true
			}
		case "up":
			if c.Y() > car.Y() && c.Y()-car.Y() < minimumDistance && c.X() == car.X() {
				return true
			}
		case "down":
			if c.Y() < car.Y() && car.Y()-c.Y() < minimumDistance && c.X() == car.X() {
				return true
			}
		}
	}
	return false
}
