package engine

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

var QuatIdent = mgl32.QuatIdent()

type Camera struct {
	*GameObject
	Window *glfw.Window
	Target Vec3
	Up     Vec3
}

// NewCamera creates a new camera object with the given position, target, up vector and Field of view
func NewCamera(window *glfw.Window, position, target, up Vec3) *Camera {
	return &Camera{
		Window: window,
		GameObject: &GameObject{
			Position: position,
			Rotation: QuatIdent,
		},
		Target: target,
		Up:     up,
	}
}

// ViewMatrix returns the view matrix for the camera
func (c *Camera) ViewMatrix() mgl32.Mat4 {
	// Apply rotation to camera position
	position := c.Rotation.Inverse().Rotate(c.Position.ToMGl32())
	return mgl32.LookAtV(position, c.Target.ToMGl32(), c.Up.ToMGl32())
}

// ProjectionMatrix returns the projection matrix for the camera
func (c *Camera) ProjectionMatrix(width, height float32) mgl32.Mat4 {
	return mgl32.Perspective(mgl32.DegToRad(45), width/height, 0.1, 100.0)
}

// Update updates the camera's position and target based on the given input
func (c *Camera) Update(window *glfw.Window, dt float64) {
	// Camera controls
	cameraSpeed := float64(0.05)
	if window.GetKey(glfw.KeyA) == glfw.Press {
		c.Position.X -= cameraSpeed
		c.Target.X -= cameraSpeed
	}
	if window.GetKey(glfw.KeyD) == glfw.Press {
		c.Position.X += cameraSpeed
		c.Target.X += cameraSpeed
	}
	if window.GetKey(glfw.KeyW) == glfw.Press {
		c.Position.Y += cameraSpeed
		c.Target.Y += cameraSpeed
	}
	if window.GetKey(glfw.KeyS) == glfw.Press {
		c.Position.Y -= cameraSpeed
		c.Target.Y -= cameraSpeed
	}

	if window.GetKey(glfw.KeyDown) == glfw.Press {
		c.Position.Z += cameraSpeed
		c.Target.Z += cameraSpeed
	}

	if window.GetKey(glfw.KeyUp) == glfw.Press {
		c.Position.Z -= cameraSpeed
		c.Target.Z -= cameraSpeed
	}

	if window.GetKey(glfw.KeyLeft) == glfw.Press {
		c.Rotation = c.Rotation.Mul(mgl32.QuatRotate(mgl32.DegToRad(1), c.Up.ToMGl32()))
	}

	if window.GetKey(glfw.KeyRight) == glfw.Press {
		c.Rotation = c.Rotation.Mul(mgl32.QuatRotate(mgl32.DegToRad(-1), c.Up.ToMGl32()))
	}

}
