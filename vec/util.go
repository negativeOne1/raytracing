package vec

import (
	"fmt"
	"io"
)

func Add(v ...Vec3) Vec3 {
	d := New(v[0].e[0], v[0].e[1], v[0].e[2])
	for _, u := range v[1:] {
		d.e[0] += u.e[0]
		d.e[1] += u.e[1]
		d.e[2] += u.e[2]
	}
	return d
}

func Sub(v ...Vec3) Vec3 {
	d := New(v[0].e[0], v[0].e[1], v[0].e[2])
	for _, u := range v[1:] {
		d.e[0] -= u.e[0]
		d.e[1] -= u.e[1]
		d.e[2] -= u.e[2]
	}
	return d
}

func Mul(u, v Vec3) Vec3 {
	return New(
		u.e[0]*v.e[0],
		u.e[1]*v.e[1],
		u.e[2]*v.e[2],
	)
}

func Dot(u, v Vec3) float64 {
	return u.e[0]*v.e[0] + u.e[1]*v.e[1] + u.e[2]*v.e[2]
}

func Cross(u, v Vec3) Vec3 {
	return New(
		u.e[1]*v.e[2]-u.e[2]*v.e[1],
		u.e[2]*v.e[0]-u.e[0]*v.e[2],
		u.e[0]*v.e[1]-u.e[1]*v.e[0],
	)
}

func WriteColor(w io.Writer, v Vec3) {
	c := 255.999
	b := fmt.Sprintf("%d %d %d\n",
		int(c*v.X()),
		int(c*v.Y()),
		int(c*v.Z()),
	)
	w.Write([]byte(b))
}
