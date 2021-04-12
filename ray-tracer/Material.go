package main

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