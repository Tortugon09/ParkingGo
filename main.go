package main

import (
	"github.com/oakmound/oak/v4"
	"estacionamiento/scenes"
)

func main() {
	parkingScene := scenes.NewParkingScene()

	parkingScene.Start()

	_ = oak.Init("parkingScene")
}
