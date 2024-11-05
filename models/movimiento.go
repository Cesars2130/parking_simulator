package models

import (
	"time"
	"github.com/faiface/pixel"
)

func EsperarPorPosicion(id int, destinoX float64) {
	for {
		posVehiculo := EncontrarPosicionVehiculo(id)
		if posVehiculo.X >= destinoX {
			break
		}
		time.Sleep(16 * time.Millisecond)
	}
}

func ReiniciarPosicionVehiculo(id int) {
	MutexVehiculos.Lock()
	defer MutexVehiculos.Unlock()
	for i := range Vehiculos {
		if Vehiculos[i].ID == id {
			Vehiculos[i].Posicion = pixel.V(0, 300)
		}
	}
}

func EncontrarPosicionVehiculo(id int) pixel.Vec {
	MutexVehiculos.Lock()
	defer MutexVehiculos.Unlock()
	for _, vehiculo := range Vehiculos {
		if vehiculo.ID == id {
			return vehiculo.Posicion
		}
	}
	return pixel.Vec{}
}

func EstacionarVehiculo(vehiculo *Vehiculo, destinoX, destinoY float64) {
	vehiculo.Posicion.X = destinoX
	vehiculo.Posicion.Y = destinoY
	vehiculo.Estacionado = true
	EstablecerHoraSalida(vehiculo)
}