package scenes

import (
	"github.com/faiface/pixel/pixelgl"
	"carro/views"
	"carro/models"
)

func Renderizar(ventana *pixelgl.Window) {
	views.DibujarEstacionamiento(ventana, models.ObtenerVehiculos())
	ventana.Update()
}
