package object

type Intersection struct {
	T      float64
	Object Object
}

type Intersections struct {
	Xs []Intersection
}

func (i Intersections) Defined() bool {
	return len(i.Xs) != 0
}

func (i Intersections) Hit() Intersection {
	var result Intersection
	for _, in := range i.Xs {
		if result.Object == nil && in.T >= 0 {
			result = in
		} else if in.T >= 0 && in.T < result.T {
			result = in
		}
	}
	return result
}
