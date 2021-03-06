package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
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


	samples := 100
	depth := 50
	cam := makeCamera(2.0,1.0)


	sphere := Sphere{Center: Vector{0,0,-1}, Radius: 0.5, Material: Lambertian{Vector{0.4,0.4,0.4}}}
	floor := Sphere{Center: Vector{0,-100.5,-1},Radius: 100,  Material: Lambertian{Vector{0.6,0.1,0.4}}}
	left := Sphere{Vector{-1, 0, -1}, 0.5, Metal{Vector{0.8, 0.8, 0.8},0.3}}
	right := Sphere{Vector{1, 0, -1}, -0.4, Dielectric{1.5}}
	right1 := Sphere{Vector{1, 0, -1}, 0.5, Dielectric{1.5}}

	list := List{[]Hittable{sphere,floor, left, right, right1}}

	f, _ := os.Create("out.ppm")
	fmt.Fprintf(f, "P3\n%d %d\n255\n", imgWidth, imgHeight)

	for i := imgHeight - 1; i >= 0; i-- {
		fmt.Println(i)
		for j := 0 ; j < imgWidth; j++ {

			pxColor := Vector{}
			for s := 0; s < samples; s++ {
				u := (float64(j)+rand.Float64())/ float64(imgWidth)
				v := (float64(i)+rand.Float64())/ float64(imgHeight)
				r := cam.getRay(u,v)
				pxColor = pxColor.Add(r.Color(list,depth))
			}

			go writeColor(i,j,pxColor, img, samples)

			r := math.Sqrt(pxColor.X * (1.0/float64(samples)))
			g := math.Sqrt(pxColor.Y * (1.0/float64(samples)))
			b := math.Sqrt(pxColor.Z * (1.0/float64(samples)))
			fmt.Fprintf(f, "%d %d %d\n", int(clamp(r,0.0,0.99) * 256), int(clamp(g,0.0,0.99) * 256),int(clamp(b,0.0,0.99) * 256))

		}
	}

	out, _ := os.Create("scene.png")
	png.Encode(out, img)

}



func writeColor(i,j int , colorVector Vector, img *image.RGBA, samples int ) {
	r := math.Sqrt( colorVector.X * (1.0/ float64(samples)))
	g := math.Sqrt( colorVector.Y * (1.0/ float64(samples)))
	b := math.Sqrt( colorVector.Z * (1.0/ float64(samples)))
	img.Set(j,i,
		color.RGBA{
		R: uint8(clamp(r,0.0,0.99) * 256),
		G: uint8(clamp(g,0.0,0.99) * 256),
		B: uint8(clamp(b,0.0,0.99) * 256),
		A: 255,
		})
}

func clamp(x, min, max float64) float64 {
	if x < min {return min}
	if x > max {return max}
	return x
}
