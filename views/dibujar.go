package views

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"carro/models"
	_"image/png"
)


func DibujarEstacionamiento(ventana *pixelgl.Window, vehiculos []models.Vehiculo) {
    if fondo == nil {
        cargarFondo()
    }
    if imagenCarro == nil {
        cargarCarro()
    }

    fondo.Draw(ventana, pixel.IM.Moved(ventana.Bounds().Center()))

    imd := imdraw.New(nil)
    imd.Color = colornames.White


	imd.Push(pixel.V(100, 500), pixel.V(700, 500))
	imd.Line(2)
	imd.Push(pixel.V(100, 100), pixel.V(700, 100))
	imd.Line(2)

	imd.Push(pixel.V(700, 100), pixel.V(700, 500))
	imd.Line(2)

	anchoEstacionamiento := 600.0 
	anchoCarril := anchoEstacionamiento / 10 

	for i := 0.0; i < 10.0; i++ {
		desplazamientoX := 100.0 + i*anchoCarril 
		imd.Push(pixel.V(desplazamientoX, 500), pixel.V(desplazamientoX, 350))
		imd.Line(2)  

		imd.Push(pixel.V(desplazamientoX, 250), pixel.V(desplazamientoX, 100))
		imd.Line(2)  
	}

    anchoVehiculo := anchoCarril / 4
    for _, vehiculo := range vehiculos {
        matriz := pixel.IM.Scaled(pixel.ZV, anchoVehiculo*2/imagenCarro.Frame().W()).Moved(vehiculo.Posicion)
        imagenCarro.Draw(ventana, matriz)
    }

    imd.Draw(ventana)
}
