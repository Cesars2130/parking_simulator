package models

import (
	"math/rand"
	"time"
)

func GeneradorVehiculos() {
	id := 0
	for {
		id++
		vehiculo := CrearVehiculo(id)
		CanalVehiculos <- vehiculo
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)+500))
	}
}

func Inicializar() {
	CanalVehiculos = make(chan Vehiculo)
	go GeneradorVehiculos()
}