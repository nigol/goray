package main

import (
    "testing"
)

func TestCreateTuple(t *testing.T) {
    // Test for point w = 1.0
    tup := CreateTuple(4.3, -4.2, 3.1, 1.0)
    comp := tup.x == 4.3 && tup.y == -4.2 && tup.z == 3.1 && tup.w == 1.0
    if !comp {
        t.Errorf("Tuple is not a point")
    }
    // Test for vector w = 0.0
    tup = CreateTuple(4.3, -4.2, 3.1, 0.0)
    comp = tup.x == 4.3 && tup.y == -4.2 && tup.z == 3.1 && tup.w == 0.0
    if !comp {
        t.Errorf("Tuple is not a vector")
    }
}

func TestCreatePoint(t *testing.T) {
    tup := CreatePoint(4.3, -4.2, 3.1)
    comp := tup.x == 4.3 && tup.y == -4.2 && tup.z == 3.1 && tup.w == 1.0
    if !comp {
        t.Errorf("Tuple is not a point")
    }
}

func TestCreateVector(t *testing.T) {
    tup := CreateVector(4.3, -4.2, 3.1)
    comp := tup.x == 4.3 && tup.y == -4.2 && tup.z == 3.1 && tup.w == 0.0
    if !comp {
        t.Errorf("Tuple is not a vector")
    }
}

func TestEqual(t *testing.T) {
    t1 := CreateTuple(1, 2, 3, 4)
    t2 := CreateTuple(1, 2, 3, 4)
    if !t1.Equal(t2) {
        t.Errorf("Tuples should be equal")
    }
    t1 = CreateTuple(1, 2, 3, 4)
    t2 = CreateTuple(12, 2, 3, 4)
    if t1.Equal(t2) {
        t.Errorf("Tuples should be different")
    }
}
