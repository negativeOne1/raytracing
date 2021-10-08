package main

import (
	"fmt"
	"math"
	"os"

	"github.com/negativeOne1/raytracing/ray"
	"github.com/negativeOne1/raytracing/vec"
)

func hitSphere(center vec.Vec3, radius float64, r ray.Ray) float64 {
	oc := vec.Sub(r.Origin(), center)
	a := vec.Dot(r.Direction(), r.Direction())
	b := 2.0 * vec.Dot(oc, r.Direction())
	c := vec.Dot(oc, oc) - radius*radius
	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return -1.0
	}
	return (-b - math.Sqrt(discriminant)) / (2.0 * a)
}

func rayColor(r ray.Ray) vec.Vec3 {
	t := hitSphere(vec.New(0, 0, -1), 0.5, r)

	if t > 0 {
		N := vec.Sub(r.At(t), vec.New(0, 0, -1)).AsUnit()
		return vec.New(N.X()+1, N.Y()+1, N.Z()+1).Mul(0.5)
	}
	unitDirection := r.Direction().AsUnit()

	t = 0.5 * (unitDirection.Y() + 1.0)
	v1 := vec.New(1.0, 1.0, 1.0).Mul(1.0 - t)
	v2 := vec.New(0.5, 0.7, 1.0).Mul(t)
	return vec.Add(v2, v1)
}

var (
	//Image
	aspectRatio float64 = 16.0 / 9.0
	imageWidth  int     = 400
	imageHeight int     = int(float64(imageWidth) / aspectRatio)

	//Camera
	viewportHeight float64 = 2.0
	viewportWidth  float64 = aspectRatio * viewportHeight
	focalLength    float64 = 1.0

	origin          vec.Vec3 = vec.New(0, 0, 0)
	horizontal      vec.Vec3 = vec.New(float64(viewportWidth), 0, 0)
	vertical        vec.Vec3 = vec.New(0, float64(viewportHeight), 0)
	lowerLeftCorner          = vec.Sub(origin, horizontal.Dev(2), vertical.Dev(2), vec.New(0, 0, focalLength))
)

func main() {
	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)
	fmt.Fprintf(os.Stderr, "\t %#v\n", vertical)

	for j := imageHeight - 1; j >= 0; j-- {
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %d\n", j)
		for i := 0; i < imageWidth; i++ {
			u := float64(i) / float64(imageWidth-1)
			v := float64(j) / float64(imageHeight-1)
			r := ray.New(origin, vec.Add(lowerLeftCorner, horizontal.Mul(u), vertical.Mul(v), origin.Neg()))
			pixelColor := rayColor(r)
			vec.WriteColor(os.Stdout, pixelColor)
		}
	}
}
