package main

import (
	"fmt"
	"math"
)

type Vec3 struct {
	x, y, z float64
}

func makeVec3(x, y, z float64) Vec3 {
	return Vec3{x, y, z}
}

// changes the input vector
func (v *Vec3) normalize() {
	nor2 := v.length2()
	if nor2 > 0 {
		invNor := 1.0 / math.Sqrt(nor2)
		v.x *= invNor
		v.y *= invNor
		v.z *= invNor
	}
}

func (v Vec3) multScaler(f float64) Vec3 {
	return Vec3{v.x * f, v.y * f, v.z * f}
}
func (v Vec3) multVec3(w Vec3) Vec3 {
	return Vec3{v.x * w.x, v.y * w.y, v.z * w.z}
}
func (v Vec3) dot(w Vec3) float64 {
	return v.x*w.x + v.y*w.y + v.z*w.z
}

func (v Vec3) length2() float64 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

type Sphere struct {
	center                      Vec3
	radius, radius2             float64
	surfaceColor, emissionColor Vec3
	transparency, reflection    float64
}

func newSphere(c Vec3, r float64, sc Vec3, ref, tr float64) Sphere {
	return Sphere{center: c, radius: r, radius2: r * r,
		surfaceColor: sc, reflection: ref, transparency: tr}
}

func render(spheres []Sphere) {
	const width = 640
	const height = 480
	//var image []Vec3 = make([]Vec3, width*height)
	//invWidth := 1 / width
	//invHeight := 1 / height
}

func main() {
	var spheres []Sphere
	spheres = append(spheres, newSphere(Vec3{0, -10004, -20}, 10000,
		Vec3{0.20, 0.20, 0.20}, 0, 0.0))
	spheres = append(spheres, newSphere(Vec3{0, 0, -20}, 4,
		Vec3{1.0, 0.32, 0.36}, 1, 0.5))
	fmt.Println(spheres)

}
