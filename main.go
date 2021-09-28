package main

import (
	"fmt"
	"os"

	"github.com/negativeOne1/raytracing/vec"
)

const (
	imageWidth  int = 256
	imageHeight int = 256
)

func main() {
	fmt.Printf("P3\n%d %d\n255\n", imageWidth, imageHeight)

	for j := imageHeight - 1; j >= 0; j-- {
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %d\n", j)
		for i := 0; i < imageWidth; i++ {
			v := vec.New(float64(i)/float64(imageWidth-1), float64(j)/float64(imageHeight-1), 0.25)
			vec.WriteColor(os.Stdout, v)
		}
	}
}
