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

func TestAddTuple(t *testing.T) {
    t1 := CreateTuple(3, -2, 5, 1)
    t2 := CreateTuple(-2, 3, 1, 0)
    t3 := CreateTuple(1, 1, 6, 1)
    t4 := t1.Add(t2)
    if !t3.Equal(t4) {
        t.Errorf("Tuples should be equal: %f, %f, %f, %f", t4.x, t4.y, t4.z, t4.w)
    }
}

func TestSubPointPoint(t * testing.T) {
    t1 := CreatePoint(3, 2, 1)
    t2 := CreatePoint(5, 6, 7)
    t3 := CreateVector(-2, -4, -6)
    t4 := t1.Sub(t2)
    if !t3.Equal(t4) {
        t.Errorf("Tuples should be equal: %f, %f, %f, %f", t4.x, t4.y, t4.z, t4.w)
    }
}

func TestSubPointVector(t * testing.T) {
    t1 := CreatePoint(3, 2, 1)
    t2 := CreateVector(5, 6, 7)
    t3 := CreatePoint(-2, -4, -6)
    t4 := t1.Sub(t2)
    if !t3.Equal(t4) {
        t.Errorf("Tuples should be equal: %f, %f, %f, %f", t4.x, t4.y, t4.z, t4.w)
    }
}

func TestSubVectorVector(t * testing.T) {
    t1 := CreateVector(3, 2, 1)
    t2 := CreateVector(5, 6, 7)
    t3 := CreateVector(-2, -4, -6)
    t4 := t1.Sub(t2)
    if !t3.Equal(t4) {
        t.Errorf("Tuples should be equal: %f, %f, %f, %f", t4.x, t4.y, t4.z, t4.w)
    }
}
