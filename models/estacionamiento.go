package models

import (
	"time"
)

func LogicaMovimientoVehiculos() {
    for i := len(Vehiculos) - 1; i >= 0; i-- {
        if Vehiculos[i].Carril == -1 && !Vehiculos[i].Estacionado {
            Vehiculos[i].Posicion.X = 100
        } else if Vehiculos[i].Carril != -1 && !Vehiculos[i].Estacionado {
            var destinoX, destinoY float64
            anchoCarril := 600.0 / 10
            if Vehiculos[i].Carril < 10 {
                destinoX = 100.0 + float64(Vehiculos[i].Carril)*anchoCarril + anchoCarril/2
                destinoY = 400 + (500-350)/2
            } else {
                destinoX = 100.0 + float64(Vehiculos[i].Carril-10)*anchoCarril + anchoCarril/2
                destinoY = 100 + (250-100)/2
            }
            EstacionarVehiculo(&Vehiculos[i], destinoX, destinoY)
        }
    }
    LogicaSalidaVehiculo()
}


func VerificarTodosEstacionados() bool {
	todosEstacionados := true
	for _, vehiculo := range Vehiculos {
		if !vehiculo.Estacionado {
			todosEstacionados = false
			break
		}
	}
	return todosEstacionados
}

func LogicaSalidaVehiculo() {
	for i := len(Vehiculos) - 1; i >= 0; i-- {
		if Vehiculos[i].Estacionado && time.Now().After(Vehiculos[i].HoraSalida) && !Vehiculos[i].Entrando {
			if !Vehiculos[i].Teletransportando {
				Vehiculos[i].Teletransportando = 		true
				Vehiculos[i].InicioTeletransporte = 	time.Now()
				Vehiculos[i].Posicion.X = 				50
				Vehiculos[i].Posicion.Y = 				200
			} else if time.Since(Vehiculos[i].InicioTeletransporte) >= time.Millisecond*500 {
				actualizarEstadoCarril(Vehiculos[i].Carril, false)
				EliminarVehiculo(i)
			}
		}
	}
}

