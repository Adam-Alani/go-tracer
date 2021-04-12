package main

type Material interface {
	Scatter(in Ray, hit HitRecord) (bool, Ray, Vector)
}

type Lambertian struct {
	C Vector
}

//type Metal struct {
//	C Vector
//}

func (l Lambertian) Scatter(in Ray, hit HitRecord) (bool, Ray, Vector) {
	dir := randomHemisphere(hit.Normal)
	return true, Ray{hit.P, dir}, l.C
}

func (l Lambertian) MatColor() Vector {
	return l.C
}

//func (m Metal) Scatter(in Ray, hit HitRecord) (bool, Ray) {
//	dir := reflect(in.Direction, hit.Normal)
//	newRay := Ray{hit.P, dir}
//	scattered := dir.Dot(hit.Normal) > 0
//	return scattered, newRay
//}
//
//
//func (m Metal) MatColor() Vector {
//	return m.C
//}
//
//func reflect(v Vector , n Vector) Vector {
//	return v.Subtract(n.Multiply(2*v.Dot(n)))
//}