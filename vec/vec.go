package vec

import (
	"fmt"
	"math"
)

//Point3 TODO: aliases don't work like in C++
type Point3 Vec3
type Color Vec3

type Vec3 struct {
	e []float64
}

func New(x, y, z float64) Vec3 {
	return Vec3{[]float64{x, y, z}}
}

func (v *Vec3) X() float64 {
	return v.e[0]
}

func (v *Vec3) Y() float64 {
	return v.e[1]
}

func (v *Vec3) Z() float64 {
	return v.e[2]
}

func (v Vec3) Neg() Vec3 {
	return New(-v.X(), -v.Y(), -v.Z())
}

func (v *Vec3) At(i int) float64 {
	return v.e[i]
}

func (v Vec3) Mul(d float64) Vec3 {
	return New(
		v.e[0]*d,
		v.e[1]*d,
		v.e[2]*d,
	)
}

func (v Vec3) Dev(d float64) Vec3 {
	return v.Mul(1 / d)
}

func (v *Vec3) LengthSquared() float64 {
	return v.e[0]*v.e[0] + v.e[1]*v.e[1] + v.e[2]*v.e[2]
}

func (v *Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v *Vec3) String() string {
	return fmt.Sprintf("%f %f %f", v.e[0], v.e[1], v.e[2])
}

func (v Vec3) AsUnit() Vec3 {
	return v.Dev(v.Length())
}
