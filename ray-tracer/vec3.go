package main

import "math"

type Vector struct {
	X,Y,Z float64
}


func (v1 Vector) Add(v2 Vector) Vector {
	return Vector{
		v1.X+v2.X,
		v1.Y+v2.Y,
		v1.Z+v2.Z,
	}
}


func (v1 Vector) AddScalar(k float64) Vector {
	return Vector{
		v1.X + k,
		v1.Y + k,
		v1.Z + k,
	}
}

func (v1 Vector) Subtract(v2 Vector) Vector {
	return Vector{
		v1.X-v2.X,
		v1.Y-v2.Y,
		v1.Z-v2.Z,
	}
}


func (v1 Vector) SubtractScalar(k float64) Vector {
	return Vector{
		v1.X - k,
		v1.Y - k,
		v1.Z - k,
	}
}


func (v1 Vector) Multiply(k float64) Vector {
	return Vector{
		v1.X * k,
		v1.Y * k,
		v1.Z * k,
	}
}

func (v1 Vector) Divide(k float64) Vector {
	return Vector{
		v1.X / k,
		v1.Y / k,
		v1.Z / k,
	}
}


func (v1 Vector) Length() float64  {
	return math.Sqrt(v1.X*v1.X + v1.Y*v1.Y + v1.Z + v1.Z)
}


func (v1 Vector) Dot(v2 Vector) float64 {
	return v1.X * v2.X + v1.Y * v2.Y + v1.Z * v2.Z
}

func (v1 Vector) Cross(v2 Vector) Vector {
	return Vector{
		v1.Y * v2.Z - v1.Z * v2.Y,
		v1.Z * v2.X - v1.X * v2.Z,
		v1.X * v2.Y - v1.Y - v2.X,
	}
}

func (v1 Vector) UnitVector() Vector {
	return v1.Divide(v1.Length())
}
