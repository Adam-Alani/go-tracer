package main



type HitRecord struct {
	P, Normal Vector
	T float64
	Material
	//frontFace bool
}

type Hittable interface {
	Hit(r Ray, tMin, tMax float64) (bool,HitRecord)

}


type List struct {
	Elements []Hittable
}

func (l List) Add(h Hittable) {
	l.Elements = append(l.Elements, h)
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
