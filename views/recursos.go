package views

import (
	"image"
	_ "image/png"
	"os"

	"github.com/faiface/pixel"
)

var (
	fondo       *pixel.Sprite
	imagenFondo pixel.Picture
	imagenCarro *pixel.Sprite
)

func cargarFondo() {
	archivo, err := os.Open("Assets/Parking.png")
	if err != nil {
		panic(err)
	}
	defer archivo.Close()

	img, _, err := image.Decode(archivo)
	if err != nil {
		panic(err)
	}

	imagenFondo = pixel.PictureDataFromImage(img)
	fondo = pixel.NewSprite(imagenFondo, imagenFondo.Bounds())
}

func cargarCarro() {
	archivo, err := os.Open("Assets/carro2.png")
	if err != nil {
		panic(err)
	}
	defer archivo.Close()

	img, _, err := image.Decode(archivo)
	if err != nil {
		panic(err)
	}

	pic := pixel.PictureDataFromImage(img)
	imagenCarro = pixel.NewSprite(pic, pic.Bounds())
}
