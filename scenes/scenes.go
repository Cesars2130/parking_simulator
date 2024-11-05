package scenes

import (
	"time"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"carro/models"
)

func Run() {
	models.Inicializar()

	ventana, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  "Estacionamiento concurrente",
		Bounds: pixel.R(0, 0, 800, 600),
	})
	if err != nil {
		panic(err)
	}

	go GestionarVehiculos()

	for !ventana.Closed() {
		ventana.Clear(colornames.White)
		Renderizar(ventana)
		ActualizarLogica()

		time.Sleep(16 * time.Millisecond)
	}
}
