package models

import (
	"math/rand"
	"time"
)

func Carril(id int) {
	CrearVehiculo(id)
	EsperarPorPosicion(id, 100)
	carril, carrilEncontrado := EncontrarCarrilDisponible()
	if !carrilEncontrado {
		ReiniciarPosicionVehiculo(id)
		return
	}
	AsignarCarrilAVehiculo(id, carril)
}


func EncontrarCarrilDisponible() (int, bool) {
	MutexCarril.Lock()
	defer MutexCarril.Unlock()
	rand.Seed(time.Now().UnixNano())
	carriles := rand.Perm(numCarriles)
	for _, c := range carriles {
		if !EstadoCarril[c] {
			EstadoCarril[c] = true
			return c, true
		}
	}
	return -1, false
}

func actualizarEstadoCarril(carril int, estado bool) {
	MutexCarril.Lock()
	defer MutexCarril.Unlock()
	EstadoCarril[carril] = estado
}
