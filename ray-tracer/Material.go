package main

import (
	"math"
	"math/rand"
)

type Material interface {
	Scatter(in Ray, hit HitRecord) (bool, Ray)
	MatColor() Vector
}


type Lambertian struct {
	C Vector
}


func (l Lambertian) Scatter(in Ray, hit HitRecord) (bool, Ray) {
	dir := randomHemisphere(hit.Normal)
	return true, Ray{hit.P, dir}
}

func (l Lambertian) MatColor() Vector {
	return l.C
}



type Metal struct {
	C Vector
	Fuzz float64
}


func (m Metal) Scatter(in Ray, hit HitRecord) (bool, Ray) {
	reflected := reflect(in.Direction, hit.Normal)
	newRay := Ray{hit.P, reflected.Add(randomHemisphere(hit.Normal).Multiply(m.Fuzz))}
	return newRay.Direction.Dot(hit.Normal) > 0 , newRay
}


func (m Metal) MatColor() Vector {
	return m.C
}

func reflect(v Vector , n Vector) Vector {
	return v.Subtract(n.Multiply(2*v.Dot(n)))
}




type Dielectric struct {
	I float64
}

func (d Dielectric) MatColor() Vector {
	return Vector{1,1,1}
}
func (d Dielectric) Scatter(in Ray, hit HitRecord) (bool , Ray) {

	uDir := in.Direction.UnitVector()
	cos := math.Min(uDir.Multiply(-1).Dot(hit.Normal), 1.0)
	sin := math.Sqrt(1 - cos*cos)

	cantRefract := d.I*sin > 1.0
	var direction Vector
	if cantRefract || d.schlick(cos) > rand.Float64() {
		direction = reflect(in.Direction, hit.Normal)
	} else {
		direction = refract(in.Direction, hit.Normal, d.I)
	}

	return true , Ray{hit.P, direction}

}

func refract(in , out Vector,  niOverNt float64)  Vector {

	theta := math.Min(in.Multiply(-1).Dot(out), 1)
	r_out_perp :=(in.Add(out.Multiply(theta))).Multiply( niOverNt)
	r_out_parallel := out.Multiply(-math.Sqrt(math.Abs(1.0 - r_out_perp.LengthSquared()))) ;
	return r_out_perp.Add(r_out_parallel)
}

func (d Dielectric) schlick(cos float64) float64 {
	r0 := (1-d.I) / (1+d.I)
	r0 = r0*r0
	return r0+(1-r0)* math.Pow(1-cos, 5)
}

