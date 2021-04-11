package main

type Ray struct {
	Origin, Direction Vector
}


func (r Ray) At(k float64) Vector {
	// Computes P(k) = A + k*B where A -> origin and B -> direction
	return r.Origin.Add(r.Direction.Multiply(k))
}


func (r Ray) Color() Vector {
	unitDir := r.Direction.UnitVector()
	k := 0.5*(unitDir.Y + 1.0)

	white := Vector{1.0,1.0,1.0}
	blue := Vector{0.5, 0.7, 1.0}
	return white.Multiply(1.0-k).Add(blue.Multiply(k))
}