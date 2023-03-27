package main

import (
	//    "github.com/go-gl/gl/v4.1-core/gl"
	mgl "github.com/go-gl/mathgl/mgl32"
    "github.com/municorn/chicky-chicky-go/maths"
)


// Camera is a camera
type Camera struct {
	position    maths.Vec3
	fov         float32
	perspective mgl.Mat4 // perspective matrix
	orientation mgl.Mat4 // stores position and rotation of the camera
	aspectRatio float32
}

// NewCamera constructs and returns a new Camera object
func NewCamera(position maths.Vec3, fov float32, aspectRatio float32) Camera {
	c := Camera{position: position, fov: fov, aspectRatio: aspectRatio}
	c.UpdateAllMatrices()
	return c
}

// SetPosition sets the position of Camera c
func (c *Camera) SetPosition(p maths.Vec3) {
	c.position = p
	c.UpdateOrientationMatrix()
}

// Position returns the position of Camera c
func (c *Camera) Position() maths.Vec3 {
	return c.position
}

// UpdateAllMatrices updates the camera's perspective and
// camera matrices
func (c *Camera) UpdateAllMatrices() {
	c.UpdatePerspectiveMatrix()
	c.UpdateOrientationMatrix()
}

// UpdatePerspectiveMatrix updates the perspective matrix
// of Camera c
func (c *Camera) UpdatePerspectiveMatrix() {
	c.perspective = mgl.Perspective(c.fov, c.aspectRatio, 0.1, 10000)
}

// UpdateOrientationMatrix updates the orientation (position,
// rotation) matrix of Camera c
func (c *Camera) UpdateOrientationMatrix() mgl.Mat4 {
	translationMatrix := mgl.Translate3D(c.position.X, c.position.Y, c.position.Z);
	return mgl.Ident4().Mul4(translationMatrix)
}

// Matrices returns both the perspective and orientation
// matrices of Camera c.
func (c *Camera) Matrices() (perspective mgl.Mat4, orientation mgl.Mat4) {
	perspective = c.perspective
	orientation = c.orientation
	return
}
