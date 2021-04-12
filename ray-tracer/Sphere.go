package main

import "math"

type Sphere struct {
	Center Vector
	Radius float64
}

func (s Sphere) Hit(r Ray, tMin, tMax float64) (bool,HitRecord) {
	oc := r.Origin.Subtract(s.Center)
	a := r.Direction.Dot(r.Direction)
	b := oc.Dot(r.Direction)
	c := oc.Dot(oc) - s.Radius*s.Radius

	rec := HitRecord{}
	delta := b*b - a*c
	if delta <= 0 {
		return false,rec
	}



	t := (-b - math.Sqrt(delta)) / a
	if t > tMin && t < tMax {
		rec.P = r.At(t)
		rec.T = t
		rec.Normal =   r.At(t).Subtract(s.Center).UnitVector()
		return true,rec
	}

	t = (-b + math.Sqrt(delta)) / a
	if t > tMin && t < tMax {
		rec.P = r.At(t)
		rec.T = t
		rec.Normal = r.At(t).Subtract(s.Center).UnitVector()
		return true, rec
	}
	return false,rec
}
