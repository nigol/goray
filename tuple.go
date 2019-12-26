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
    result := Tuple{
        x,
        y,
        z,
        1.0,
    }
    return result
}

func CreateVector(x float64, y float64, z float64) Tuple {
    result := Tuple{
        x,
        y,
        z,
        0.0,
    }
    return result
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
    return t.x == o.x && t.y == o.y && t.z == o.z && t.w == o.w
}
