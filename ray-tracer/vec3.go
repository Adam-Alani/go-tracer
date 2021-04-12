package main

import (
	"math"
	"math/rand"
)

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
func (v1 Vector) MultiplyVector(v2 Vector) Vector {
	return Vector{
		v1.X * v2.X ,
		v1.Y * v2.Y,
		v1.Z * v2.Z,
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

func (v1 Vector) LengthSquared() float64  {
	return v1.X*v1.X + v1.Y*v1.Y + v1.Z + v1.Z
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
	k := 1.0 / v1.Length()
	return Vector{
		v1.X * k,
		v1.Y * k,
		v1.Z * k,
	}
}


func randomUnitSphere() Vector {
	for true {
		randomVector := Vector{rand.Float64(),rand.Float64(),rand.Float64()}.Multiply(2).Subtract(Vector{1,1,1})
		if randomVector.LengthSquared() >= 1.0 {
			return randomVector
		}
	}
	return Vector{1,0,0}
}

func randomUnitVector() Vector {
	return randomUnitSphere().UnitVector()
}

func randomHemisphere(normal Vector) Vector {
	UnitSphere := randomUnitSphere()
	if UnitSphere.Dot(normal) > 0.0 {
		return UnitSphere
	} else {
		return UnitSphere.Multiply(-1)
	}
}