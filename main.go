package main

import (
	"fmt"
	"os"

	"github.com/negativeOne1/raytracing/ray"
	"github.com/negativeOne1/raytracing/vec"
)

func rayColor(r ray.Ray) vec.Vec3 {
	unitDirection := r.Direction().AsUnit()

	t := 0.5 * (unitDirection.Y() + 1.0)
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
	focalLength    float64 = 2.0

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
