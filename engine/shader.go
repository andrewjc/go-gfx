package engine

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type ShaderProgram struct {
	program                uint32
	UniformLocationCache   map[string]int32
	shaderReferenceTracker int64
}

func (s *ShaderProgram) GetUniformLocation(name string) int32 {

	if s.UniformLocationCache == nil {
		s.UniformLocationCache = make(map[string]int32)
	}

	if loc, ok := s.UniformLocationCache[name]; ok {
		return loc
	}

	loc := gl.GetUniformLocation(s.program, gl.Str(name+"\x00"))
	if loc == -1 {
		fmt.Printf("Could not find uniform %s", name)
	} else {
		s.UniformLocationCache[name] = loc
	}
	return loc
}

func (s *ShaderProgram) Use() {
	print("Using shader program: ", s.program, " reference tracker ", s.shaderReferenceTracker)
	gl.UseProgram(s.program)
}

func (s *ShaderProgram) Unuse() {
	gl.UseProgram(0)
}

func (s *ShaderProgram) SetFloat32(key string, value *float32) {
	loc := s.GetUniformLocation(key)
	gl.UniformMatrix4fv(loc, 1, false, value)
}

func (s *ShaderProgram) SetMat4(key string, value *mgl32.Mat4) {
	loc := s.GetUniformLocation(key)
	gl.UniformMatrix4fv(loc, 1, false, &value[0])
}

func (s *ShaderProgram) SetVec3(key string, value *float32) {
	loc := s.GetUniformLocation(key)
	gl.Uniform3fv(loc, 1, value)
}

func (s *ShaderProgram) SetFloat(key string, value float32) {
	loc := s.GetUniformLocation(key)
	gl.Uniform1f(loc, value)
}

func (s *ShaderProgram) SetInt(key string, value int32) {
	loc := s.GetUniformLocation(key)
	gl.Uniform1i(loc, value)
}

func LoadShader(vertShader string, fragShader string) *ShaderProgram {
	shaderProgram, err := NewShaderProgram(vertShader, fragShader)
	if err != nil {
		panic(err)
	}
	return shaderProgram
}

func NewShaderProgram(vertexShaderPath, fragmentShaderPath string) (*ShaderProgram, error) {
	GLContext().check()

	vertexShaderSource, err := LoadFile(vertexShaderPath)
	if err != nil {
		return nil, err
	}

	fragmentShaderSource, err := LoadFile(fragmentShaderPath)
	if err != nil {
		return nil, err
	}

	// Compile the vertex shader
	vertexShader := gl.CreateShader(gl.VERTEX_SHADER)
	cvertexShaderSource, freeVertexShader := gl.Strs(vertexShaderSource)
	gl.ShaderSource(vertexShader, 1, cvertexShaderSource, nil)
	freeVertexShader()
	gl.CompileShader(vertexShader)
	var status int32
	gl.GetShaderiv(vertexShader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(vertexShader, gl.INFO_LOG_LENGTH, &logLength)
		log := make([]byte, logLength)
		gl.GetShaderInfoLog(vertexShader, logLength, nil, &log[0])
		gl.DeleteShader(vertexShader)
		return nil, fmt.Errorf("failed to compile vertex shader: %s", log)
	}
	// Compile the fragment shader
	fragmentShader := gl.CreateShader(gl.FRAGMENT_SHADER)
	cfragmentShaderSource, freeFragmentShader := gl.Strs(fragmentShaderSource)
	gl.ShaderSource(fragmentShader, 1, cfragmentShaderSource, nil)
	freeFragmentShader()
	gl.CompileShader(fragmentShader)
	gl.GetShaderiv(fragmentShader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(fragmentShader, gl.INFO_LOG_LENGTH, &logLength)
		log := make([]byte, logLength)
		gl.GetProgramInfoLog(fragmentShader, logLength, nil, &log[0])
		gl.DeleteShader(vertexShader)
		gl.DeleteShader(fragmentShader)
		return nil, fmt.Errorf("failed to link shader program: %s", log)
	}
	// Link the shader program
	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)
		log := make([]byte, logLength)
		gl.GetProgramInfoLog(program, logLength, nil, &log[0])
		gl.DeleteShader(vertexShader)
		gl.DeleteShader(fragmentShader)
		gl.DeleteProgram(program)
		return nil, fmt.Errorf("failed to link shader program: %s", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return &ShaderProgram{program, map[string]int32{}, 0}, nil
}
