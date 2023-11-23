package scenes

import (
	"estacionamiento/models"
	"image/color"
	"math/rand"
	"sync"
	"time"

	"github.com/oakmound/oak/v4"
	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/entities"
	"github.com/oakmound/oak/v4/event"
	"github.com/oakmound/oak/v4/render"
	"github.com/oakmound/oak/v4/scene"
)

var (
	ParkingSpaces = []*models.ParkingSpace{
		models.NewParkingSpace(410, 210, 440, 240, 1, 1),
		models.NewParkingSpace(410, 255, 440, 285, 1, 2),
		models.NewParkingSpace(410, 300, 440, 330, 1, 3),
		models.NewParkingSpace(410, 345, 440, 375, 1, 4),
		models.NewParkingSpace(410, 390, 440, 420, 1, 5),

		models.NewParkingSpace(320, 210, 350, 240, 2, 6),
		models.NewParkingSpace(320, 255, 350, 285, 2, 7),
		models.NewParkingSpace(320, 300, 350, 330, 2, 8),
		models.NewParkingSpace(320, 345, 350, 375, 2, 9),
		models.NewParkingSpace(320, 390, 350, 420, 2, 10),

		models.NewParkingSpace(230, 210, 260, 240, 3, 11),
		models.NewParkingSpace(230, 255, 260, 285, 3, 12),
		models.NewParkingSpace(230, 300, 260, 330, 3, 13),
		models.NewParkingSpace(230, 345, 260, 375, 3, 14),
		models.NewParkingSpace(230, 390, 260, 420, 3, 15),

		models.NewParkingSpace(140, 210, 170, 240, 4, 16),
		models.NewParkingSpace(140, 255, 170, 285, 4, 17),
		models.NewParkingSpace(140, 300, 170, 330, 4, 18),
		models.NewParkingSpace(140, 345, 170, 375, 4, 19),
		models.NewParkingSpace(140, 390, 170, 420, 4, 20),
	}
	estacionamiento = models.NewParkingLot(ParkingSpaces)
	colaAutomoviles = estacionamiento.GetCarQueue()
	doorMutex       sync.Mutex
	carManager      = models.NewCarController()
)

type ParkingScene struct{}

func NewParkingScene() *ParkingScene {
	return &ParkingScene{}
}

func (ps *ParkingScene) Start() {
	firstTime := true

	_ = oak.AddScene("parkingScene", scene.Scene{
		Start: func(ctx *scene.Context) {
			_ = ctx.Window.SetBorderless(true)
			PrepareScene(ctx)

			event.GlobalBind(ctx, event.Enter, func(enterPayload event.EnterPayload) event.Response {
				if !firstTime {
					return 0
				}

				firstTime = false

				for i := 0; i < 100; i++ {
					go CarCycle(ctx)

					time.Sleep(time.Millisecond * time.Duration(GetRandomNumber(1000, 2000)))
				}

				return 0
			})
		},
	})
}

func PrepareScene(ctx *scene.Context) {
	parkingLotArea := floatgeom.NewRect2(0, 0, 1000, 1000)
	entities.New(ctx, entities.WithRect(parkingLotArea), entities.WithColor(color.RGBA{128, 128, 128, 128}))

	for _, parkingSpace := range ParkingSpaces {
		spritePath := "assets/Parking.jpg"
		sprite, _ := render.LoadSprite(spritePath)
		entities.New(ctx, entities.WithRect(*parkingSpace.GetArea()), entities.WithRenderable(sprite))
	}
}

func CarCycle(ctx *scene.Context) {
	car := models.NewCar(ctx)

	carManager.AddCar(car)

	car.Enqueue(carManager)

	availableSpace := estacionamiento.GetAvailableParkingSpace()

	doorMutex.Lock()

	car.JoinGate(carManager)

	doorMutex.Unlock()

	car.Park(availableSpace, carManager)

	time.Sleep(time.Millisecond * time.Duration(GetRandomNumber(40000, 50000)))

	car.LeaveSpace(carManager)

	estacionamiento.ReleaseParkingSpace(availableSpace)

	car.Leave(availableSpace, carManager)

	doorMutex.Lock()

	car.LeaveGate(carManager)

	doorMutex.Unlock()

	car.MoveAway(carManager)

	car.Destroy()

	carManager.DeleteCar(car)
}

func GetRandomNumber(min, max int) float64 {
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)
	return float64(generator.Intn(max-min+1) + min)
}
