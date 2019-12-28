package main

type Tuple struct {
    x float64
    y float64
    z float64
    w float64
}

func CreateTuple(x float64, y float64, z float64, w float64) Tuple {
    result := Tuple{
        x,
        y,
        z,
        w,
    }
    return result
}

func CreatePoint(x float64, y float64, z float64) Tuple {
    return CreateTuple(x, y, z, 1.0)
}

func CreateVector(x float64, y float64, z float64) Tuple {
    return CreateTuple(x, y, z, 0.0)
}

func abs(x float64) float64 {
    if x < 0 {
        return -x
    }
    return x
}

func equal(a float64, b float64) bool {
    e := 0.00001
    return abs(a - b) < e
}

func (t Tuple) Equal(o Tuple) bool {
    return equal(t.x, o.x) && equal(t.y, o.y) && equal(t.z, o.z) && equal(t.w, o.w)
}

func (t Tuple) Add(o Tuple) Tuple {
    return CreateTuple(t.x + o.x, t.y + o.y, t.z + o.z, t.w + o.w)
}

func (t Tuple) Sub(o Tuple) Tuple {
    return CreateTuple(t.x - o.x, t.y - o.y, t.z - o.z, t.w - o.w)
}
