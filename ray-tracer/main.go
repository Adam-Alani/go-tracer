package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {

	aspectRatio := 16.0 / 9.0

	imgWidth := 400
	imgHeight := int(float64(imgWidth)/aspectRatio)

	min := image.Point{}
	max := image.Point{X: imgWidth, Y: imgHeight}

	img := image.NewRGBA(image.Rectangle{Min: min, Max: max})


	// Camera Settings

	vh := 2.0
	vw := aspectRatio * vh
	focalLength := -1.0

	origin := Vector{0,0,0}
	horizontal := Vector{vw,0,0}
	vertical := Vector{0,vh,0}
	// origin - horizontal/2 - vertical/2 - focalLength
	lowerLeft := origin.Subtract(horizontal.Divide(2)).Subtract(vertical.Divide(2)).Add(Vector{0,0, focalLength})


	sphere := Sphere{Center: Vector{0,0,-1}, Radius: 0.5}
	floor := Sphere{Center: Vector{0,-100.5,-1},Radius: 100}

	list := List{[]Hittable{sphere,floor}}
	f, _ := os.Create("out.ppm")
	fmt.Fprintf(f, "P3\n%d %d\n255\n", imgWidth, imgHeight)

	for i := imgHeight - 1; i >= 0; i-- {
		fmt.Println(i)
		for j := 0 ; j < imgWidth; j++ {


			u := float64(j)/ float64(imgWidth-1)
			v := float64(i)/ float64(imgHeight-1)


			r := Ray{origin,lowerLeft.Add(horizontal.Multiply(u)).Add(vertical.Multiply(v)).Subtract(origin)}
			pxColor := r.Color(list)

			writeColor(i,j,pxColor, img)

			fmt.Fprintf(f, "%d %d %d\n", int(pxColor.X*255.99),  int(pxColor.Y*255.99),  int(pxColor.Z*255.99))

		}
	}

	out, _ := os.Create("scene.png")
	png.Encode(out, img)

}



func writeColor(i,j int , colorVector Vector, img *image.RGBA ) {
	img.Set(j,i,
		color.RGBA{
		R: uint8(colorVector.X * 255.999),
		G: uint8(colorVector.Y * 255.999),
		B: uint8(colorVector.Z * 255.999),
		A: 255,
		})
}
