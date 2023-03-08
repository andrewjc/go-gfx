package engine

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

type ForwardRenderer struct {
	window *glfw.Window
	camera *Camera
}

func NewForwardRenderer(window *glfw.Window, camera *Camera) *ForwardRenderer {
	return &ForwardRenderer{
		window: window,
		camera: camera,
	}
}
