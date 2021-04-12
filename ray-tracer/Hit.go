package main



type HitRecord struct {
	P, Normal Vector
	T float64
	//frontFace bool
}

type Hittable interface {
	Hit(r Ray, tMin, tMax float64) (bool,HitRecord)
}


type List struct {
	Elements []Hittable
}

func (l List) Hit(r Ray, tMin, tMax float64) (bool,HitRecord) {
	hitAny := false
	closest := tMax
	rec := HitRecord{}

	for _, el := range l.Elements{
		contact, tRec := el.Hit(r, tMin, closest)

		if contact {
			hitAny = true
			closest = tRec.T
			rec = tRec
		}
	}
	return hitAny, rec
}
