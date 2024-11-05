package scenes

import "carro/models"

func GestionarVehiculos() {
	for vehiculo := range models.CanalVehiculos {
		models.MutexCarril.Lock()
		for _, ocupado := range models.EstadoCarril {
			if !ocupado {
				break
			}
		}
		models.MutexCarril.Unlock()

		go models.Carril(vehiculo.ID)
	}
}
