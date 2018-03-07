package main

import (
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

func main() {
}
