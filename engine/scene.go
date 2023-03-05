package engine

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Scene struct {
	DefaultShaderProgram *ShaderProgram
	Window               *glfw.Window
	Camera               *Camera
	Objects              []*GameObject
	Lights               []*Light

	Vertices []float32
	Indices  []uint32
	Vao      uint32
	Vbo      uint32
	Ebo      uint32
}

func NewScene(window *glfw.Window) (*Scene, error) {
	// Set up the vertex array object (VAO)
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	// Set up the vertex buffer object (VBO)
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	// Set up the element buffer object (EBO)
	var ebo uint32
	gl.GenBuffers(1, &ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)

	defaultShaderProgram, err := NewShaderProgram("shaders/default.vert", "shaders/default.frag")
	if err != nil {
		return nil, err
	}

	scene := &Scene{
		DefaultShaderProgram: defaultShaderProgram,
		Window:               window,
		Objects:              []*GameObject{},
		Vertices:             []float32{},
		Indices:              []uint32{},
		Vao:                  vao,
		Vbo:                  vbo,
		Ebo:                  ebo,
	}
	/*
	   // Set up the vertex attribute pointers
	   positionAttribute := uint32(gl.GetAttribLocation(program, gl.Str("position\x00")))
	   gl.EnableVertexAttribArray(positionAttribute)
	   gl.VertexAttribPointer(positionAttribute, 3, gl.FLOAT, false, 0, nil)
	*/
	return scene, nil
}

func (s *Scene) AddObject(obj *GameObject) {
	obj.Scene = s
	s.Objects = append(s.Objects, obj)
	if obj.vertices != nil && len(obj.vertices) > 0 {
		s.Vertices = append(s.Vertices, obj.vertices...)
	}
	s.Indices = append(s.Indices, uint32(len(s.Vertices))/3-1)
	gl.BindVertexArray(s.Vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, s.Vbo)
	if obj.vertices != nil && len(obj.vertices) > 0 {
		gl.BufferData(gl.ARRAY_BUFFER, len(s.Vertices)*4, gl.Ptr(s.Vertices), gl.STATIC_DRAW)
	}
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, s.Ebo)
	if obj.vertices != nil && len(obj.vertices) > 0 {
		gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(s.Indices)*4, gl.Ptr(s.Indices), gl.STATIC_DRAW)
	}
}

func (s *Scene) RemoveObject(obj *GameObject) {
	for i, o := range s.Objects {
		if o.Equals(obj) {

			s.Objects = append(s.Objects[:i], s.Objects[i+1:]...)
			s.Vertices = []float32{}
			s.Indices = []uint32{}
			obj.Scene = nil
			for _, o := range s.Objects {
				s.Vertices = append(s.Vertices, o.vertices...)
				s.Indices = append(s.Indices, uint32(len(s.Vertices))/3-1)
			}
			gl.BindVertexArray(s.Vao)
			gl.BindBuffer(gl.ARRAY_BUFFER, s.Vbo)
			gl.BufferData(gl.ARRAY_BUFFER, len(s.Vertices)*4, gl.Ptr(s.Vertices), gl.STATIC_DRAW)
			gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, s.Ebo)
			gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(s.Indices)*4, gl.Ptr(s.Indices), gl.STATIC_DRAW)
			break
		}
	}
}

func (s *Scene) Update(dt float64) {
	for _, object := range s.Objects {
		object.Update(dt)
	}
}

func (s *Scene) Render(camera *Camera) {
	// Clear the screen
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	// Loop over all objects and render them
	for _, obj := range s.Objects {

		// Set the transform for the object
		/*model := obj.Transform.ToMat4()
		  if obj.Mesh != nil {
		  	obj.Mesh.SetTransform(model)
		  }*/

		// Render the object
		obj.Render(camera)

	}
}

func (s *Scene) CreateCamera(position Vec3, target Vec3, up Vec3) *Camera {
	return NewCamera(s.Window, position, target, up)
}
