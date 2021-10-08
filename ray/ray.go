package ray

import (
	vec "github.com/negativeOne1/raytracing/vec"
)

type Ray struct {
	orig vec.Vec3
	dir  vec.Vec3
}

func New(orig, dir vec.Vec3) Ray {
	return Ray{orig, dir}
}

func (r *Ray) Origin() vec.Vec3 {
	return r.orig
}

func (r *Ray) Direction() vec.Vec3 {
	return r.dir
}

func (r *Ray) At(t float64) vec.Vec3 {
	return vec.Add(r.orig, r.dir.Mul(t))
}
