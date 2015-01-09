// Copyright (c) 2006, 2007 Sony Computer Entertainment Inc.
// Copyright (c) 2012 James Helferty
// All rights reserved.

package vectormath

type Scalar float32

type Vector2 struct {
	X Scalar	`json:"x"`
	Y Scalar	`json:"y"`
}

type Vector3 struct {
	X float32	`json:"x"`
	Y float32	`json:"y"`
	Z float32	`json:"z"`
}

type Vector4 struct {
	X, Y, Z, W float32
}

type Point3 struct {
	X, Y, Z float32
}

type Quat struct {
	X, Y, Z, W float32
}

type Matrix3 struct {
	col0, col1, col2 Vector3
}

type Matrix4 struct {
	col0, col1, col2, col3 Vector4
}

type Transform3 struct {
	col0, col1, col2, col3 Vector3
}
