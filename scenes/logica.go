package scenes

import "carro/models"

func ActualizarLogica() {
	models.MutexVehiculos.Lock()
	models.LogicaMovimientoVehiculos()
	models.MutexVehiculos.Unlock()
}
