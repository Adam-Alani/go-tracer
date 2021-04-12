package main

type Camera struct {
	LowerLeft, Horizontal, Vertical, Origin Vector
	VH,VW,FL float64
}

func makeCamera(vh,focalLength float64) Camera{
	aspectRatio := 16.0 / 9.0
	vw := aspectRatio * vh
	cam :=  Camera{
		VW: vw,
		VH: vh,
		FL: focalLength,
		Origin: Vector{0,0,0},
		Horizontal: Vector{vw,0,0},
		Vertical: Vector{0,vh,0},
	}
	cam.LowerLeft = cam.Origin.Subtract(cam.Horizontal.Divide(2)).Subtract(cam.Vertical.Divide(2)).Subtract(Vector{0,0,focalLength})
	return cam
}



func (c Camera) getRay(u,v float64) Ray {
	return Ray{Vector{0,0,0}, c.LowerLeft.Add(c.Horizontal.Multiply(u)).Add(c.Vertical.Multiply(v)).Subtract(c.Origin)}
}
