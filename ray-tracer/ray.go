package main

import "math"

type Ray struct {
	Origin, Direction Vector
}


func (r Ray) At(k float64) Vector {
	// Computes P(k) = A + k*B where A -> origin and B -> direction
	return r.Origin.Add(r.Direction.Multiply(k))
}


func (r Ray) Color(h Hittable, depth int) Vector {

	contact, rec := h.Hit(r,0.001, math.MaxFloat64)

	if contact {
		if depth > 0 {
			scattered, newRay := rec.Scatter(r,rec)
			if scattered {
				newColor := newRay.Color(h, depth-1)
				return rec.Material.MatColor().MultiplyVector(newColor)
			}
			return Vector{}
		}

	}

	t := 0.5 * (r.Direction.Y + 1.0)
	white := Vector{1,1,1}.Multiply(1-t)
	blue := Vector{0.5, 0.7, 1}.Multiply(t)

	return white.Add(blue)
}