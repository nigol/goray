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
