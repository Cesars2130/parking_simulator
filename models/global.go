package models

import (
	"sync"
	"github.com/faiface/pixel"
	"time"
)

const (
	numCarriles  = 20
	AnchoCarril = 150.0
)

type Vehiculo struct {
	ID             			int
	Posicion       			pixel.Vec
	PosicionAnterior 		pixel.Vec
	Carril         			int
	Estacionado    			bool
	HoraSalida     			time.Time
	Entrando       			bool
	Teletransportando   	bool
	InicioTeletransporte 	time.Time
}

var CanalVehiculos chan Vehiculo


var (
	EstadoCarril 				[numCarriles]bool
	Vehiculos    				[]Vehiculo
	MutexVehiculos  			sync.Mutex
	MutexCarril  				sync.Mutex
	VehiculoEntrandoOSaliendo 	bool
)