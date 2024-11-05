package models

import (
	"github.com/faiface/pixel"
	"math/rand"
	"time"
)



func CrearVehiculo(id int) Vehiculo {
	MutexVehiculos.Lock()
	defer MutexVehiculos.Unlock()
	Vehiculo := Vehiculo{
		ID:       		id,
		Posicion: 		pixel.V(50, 300),
		Carril:   		-1,
		Estacionado: 	false,
	}
	Vehiculos = append(Vehiculos, Vehiculo)
	return Vehiculo
}

func EstablecerHoraSalida(vehiculo *Vehiculo) {
	rand.Seed(time.Now().UnixNano())
	salidaEn := time.Duration(rand.Intn(5)+1) * time.Second
	vehiculo.HoraSalida = time.Now().Add(salidaEn)
}

func ObtenerVehiculos() []Vehiculo {
	return Vehiculos
}

func AsignarCarrilAVehiculo(id int, carril int) {
	MutexVehiculos.Lock()
	defer MutexVehiculos.Unlock()
	for i := range Vehiculos {
		if Vehiculos[i].ID == id {
			Vehiculos[i].Carril = carril
		}
	}
}

func EliminarVehiculo(indice int) {
	Vehiculos = append(Vehiculos[:indice], Vehiculos[indice+1:]...)
}


