package engine

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

type Scene struct {
	DefaultShaderProgram *ShaderProgram
	Renderer             *ForwardRenderer
	Window               *glfw.Window
	Camera               *Camera
	Objects              []*GameObject
	Lights               []*Light
}

func NewScene(window *glfw.Window) (*Scene, error) {

	width, height := window.GetFramebufferSize()
	gl.Viewport(0, 0, int32(width), int32(height))

	shaderProgram, err := NewShaderProgram("shaders/default.vert", "shaders/default.frag")
	if err != nil {
		return nil, err
	}

	scene := &Scene{
		DefaultShaderProgram: shaderProgram,
		Window:               window,
		Objects:              []*GameObject{},
	}

	return scene, nil
}

func (s *Scene) AddObject(obj *GameObject) {
	obj.Scene = s
	s.Objects = append(s.Objects, obj)
}

func (s *Scene) RemoveObject(obj *GameObject) {
	for i, o := range s.Objects {
		if o.Equals(obj) {
			obj.Scene = nil
			s.Objects = append(s.Objects[:i], s.Objects[i+1:]...)
			break
		}
	}
}

func (s *Scene) Update(dt float32) {
	for _, object := range s.Objects {
		object.Update(dt)
	}
}

func (s *Scene) Render(camera *Camera) {
	// Clear the screen
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	proj := camera.ProjectionMatrix()
	view := camera.ViewMatrix()

	// Loop over all objects and render them
	for _, obj := range s.Objects {
		// Render the object
		obj.Render(proj, view)
	}
}

func (s *Scene) CreateCamera(position mgl32.Vec3, target mgl32.Vec3, up mgl32.Vec3) *Camera {
	return NewCamera(s.Window, position, target, up)
}

func (s *Scene) SetRenderer(renderer *ForwardRenderer) {
	s.Renderer = renderer
}
