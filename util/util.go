package util

import "math"

type Vec2 struct {
	X float64
	Y float64
}

func (v1 *Vec2) ToInts() (X, Y int) {
	return int(v1.X), int(v1.Y)
}

func (v1 *Vec2) ToVals() (X, Y float64) {
	return v1.X, v1.Y
}

func (v1 *Vec2) Add(v2 *Vec2) *Vec2 {
	return &Vec2{v1.X + v2.X, v1.Y + v2.Y}
}

func (v1 *Vec2) Sub(v2 *Vec2) *Vec2 {
	return &Vec2{v1.X - v2.X, v1.Y - v2.Y}
}

func (v1 *Vec2) Dist(v2 *Vec2) float64 {
	return math.Sqrt(math.Pow(v1.Sub(v2).X, 2) + math.Pow(v1.Sub(v2).Y, 2))
}

func (v1 *Vec2) Dist2(v2 *Vec2) float64 {
	return (math.Pow(v1.Sub(v2).X, 2) + math.Pow(v1.Sub(v2).Y, 2))
}

func (v1 *Vec2) Norm(v2 *Vec2) *Vec2 {
	return v2.Sub(v1).Div(v1.Dist(v2))
}

func (v1 *Vec2) Norm2(v2 *Vec2) *Vec2 {
	return v2.Sub(v1).Div(v1.Dist2(v2))
}

func (v1 *Vec2) Div(val float64) *Vec2 {
	return &Vec2{X: v1.X / val, Y: v1.Y / val}
}

func (v1 *Vec2) Mul(val float64) *Vec2 {
	return &Vec2{X: v1.X * val, Y: v1.Y * val}
}

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func (v1 Vec3) ToInts() (X, Y, Z int) {
	return int(v1.X), int(v1.Y), int(v1.Z)
}

func (v1 Vec3) ToVals() (X, Y, Z float64) {
	return v1.X, v1.Y, v1.Z
}

func (v1 Vec3) Add(v2 Vec3) *Vec3 {
	return &Vec3{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}
