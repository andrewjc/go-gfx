package engine

import "C"
import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

var QuatIdent = mgl32.QuatIdent()

type Camera struct {
	*GameObject
	Window *glfw.Window
	Target mgl32.Vec3
	Up     mgl32.Vec3
	Left   mgl32.Vec3
}

// NewCamera creates a new camera object with the given position, target, up vector and Field of view
func NewCamera(window *glfw.Window, position, target, up mgl32.Vec3) *Camera {
	return &Camera{
		Window: window,
		GameObject: &GameObject{
			Position: position,
			Rotation: QuatIdent,
		},
		Target: target,
		Up:     up,
		Left:   up.Cross(target.Sub(position)).Normalize(),
	}
}

// ViewMatrix returns the view matrix for the camera
func (c *Camera) ViewMatrix() mgl32.Mat4 {
	// Apply rotation to camera position
	position := c.Rotation.Inverse().Rotate(c.Position)
	return mgl32.LookAtV(position, c.Target, c.Up)
}

// ProjectionMatrix returns the projection matrix for the camera
func (c *Camera) ProjectionMatrix(width, height float32) mgl32.Mat4 {
	return mgl32.Perspective(mgl32.DegToRad(45), width/height, 0.1, 100.0)
}

// Update updates the camera's position and target based on the given input
func (c *Camera) Update(window *glfw.Window, dt float32) {
	// Camera controls
	cameraSpeed := float32(5.0)
	rotateSpeed := float32(35.0)
	if window.GetKey(glfw.KeyW) == glfw.Press {
		c.Position = c.Position.Add(mgl32.Vec3{0, -cameraSpeed * dt, 0})
		c.Target = c.Target.Add(mgl32.Vec3{0, -cameraSpeed * dt, 0})
	}
	if window.GetKey(glfw.KeyS) == glfw.Press {
		c.Position = c.Position.Add(mgl32.Vec3{0, cameraSpeed * dt, 0})
		c.Target = c.Target.Add(mgl32.Vec3{0, cameraSpeed * dt, 0})
	}
	if window.GetKey(glfw.KeyA) == glfw.Press {
		c.Position = c.Position.Add(mgl32.Vec3{-cameraSpeed * dt, 0, 0})
		c.Target = c.Target.Add(mgl32.Vec3{-cameraSpeed * dt, 0, 0})

	}
	if window.GetKey(glfw.KeyD) == glfw.Press {
		c.Position = c.Position.Add(mgl32.Vec3{cameraSpeed * dt, 0, 0})
		c.Target = c.Target.Add(mgl32.Vec3{cameraSpeed * dt, 0, 0})
	}

	if window.GetKey(glfw.KeyDown) == glfw.Press {
		c.Rotation = c.Rotation.Mul(mgl32.QuatRotate(mgl32.DegToRad(-rotateSpeed*dt), c.Up))
	}

	if window.GetKey(glfw.KeyUp) == glfw.Press {
		c.Rotation = c.Rotation.Mul(mgl32.QuatRotate(mgl32.DegToRad(rotateSpeed*dt), c.Up))
	}

	if window.GetKey(glfw.KeyLeft) == glfw.Press {
		c.Rotation = c.Rotation.Mul(mgl32.QuatRotate(mgl32.DegToRad(rotateSpeed*dt), c.Left))
	}

	if window.GetKey(glfw.KeyRight) == glfw.Press {
		c.Rotation = c.Rotation.Mul(mgl32.QuatRotate(mgl32.DegToRad(-rotateSpeed*dt), c.Left))
	}

}
