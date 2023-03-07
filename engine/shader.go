package engine

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"log"
	"unsafe"
)

type ShaderProgram struct {
	program                uint32
	UniformLocationCache   map[string]uint32
	AttribLocationCache    map[string]uint32
	shaderReferenceTracker int64
}

func (s *ShaderProgram) GetUniformLocation(name string) uint32 {

	if s.UniformLocationCache == nil {
		s.UniformLocationCache = make(map[string]uint32)
	}

	if loc, ok := s.UniformLocationCache[name]; ok {
		return loc
	}

	loc := gl.GetUniformLocation(s.program, gl.Str(name+"\x00"))
	if loc == -1 {
		fmt.Printf("Could not find uniform %s", name)
	} else {
		s.UniformLocationCache[name] = uint32(loc)
	}
	return uint32(loc)
}

func (s *ShaderProgram) GetAttribLocation(name string) uint32 {

	if s.AttribLocationCache == nil {
		s.AttribLocationCache = make(map[string]uint32)
	}

	if loc, ok := s.AttribLocationCache[name]; ok {
		return loc
	}

	loc := gl.GetAttribLocation(s.program, gl.Str(name+"\x00"))
	if loc == -1 {
		fmt.Printf("Could not find attrib %s", name)
	} else {
		s.AttribLocationCache[name] = uint32(loc)
	}
	return uint32(loc)
}

func (s *ShaderProgram) Use() {
	gl.UseProgram(s.program)
}

func (s *ShaderProgram) Unuse() {
	gl.UseProgram(0)
}

func (s *ShaderProgram) SetFloat32UniformLocation(key string, value *float32) {
	loc := s.GetUniformLocation(key)
	gl.UniformMatrix4fv(int32(loc), 1, false, value)
}

func (s *ShaderProgram) SetMat4UniformLocation(key string, value *mgl32.Mat4) {
	loc := s.GetUniformLocation(key)
	gl.UniformMatrix4fv(int32(loc), 1, false, &value[0])
}

func (s *ShaderProgram) SetVec3UniformLocation(key string, value *float32) {
	loc := s.GetUniformLocation(key)
	gl.Uniform3fv(int32(loc), 1, value)
}

func (s *ShaderProgram) SetFloatUniformLocation(key string, value float32) {
	loc := s.GetUniformLocation(key)
	gl.Uniform1f(int32(loc), value)
}

func (s *ShaderProgram) SetIntUniformLocation(key string, value int32) {
	loc := s.GetUniformLocation(key)
	gl.Uniform1i(int32(loc), value)
}

func (s *ShaderProgram) SetFloat32AttribLocation(key string, value *float32) {
	loc := s.GetAttribLocation(key)
	gl.UniformMatrix4fv(int32(loc), 1, false, value)
}

func (s *ShaderProgram) SetMat4AttribLocation(key string, value *mgl32.Mat4) {
	loc := s.GetAttribLocation(key)
	gl.UniformMatrix4fv(int32(loc), 1, false, &value[0])
}

func (s *ShaderProgram) SetVec3AttribLocation(key string, value *float32) {
	loc := s.GetAttribLocation(key)
	gl.Uniform3fv(int32(loc), 1, value)
}

func (s *ShaderProgram) SetFloatAttribLocation(key string, value float32) {
	loc := s.GetAttribLocation(key)
	gl.Uniform1f(int32(loc), value)
}

func (s *ShaderProgram) SetIntAttribLocation(key string, value int32) {
	loc := s.GetAttribLocation(key)
	gl.Uniform1i(int32(loc), value)
}

func LoadShader(vertShader string, fragShader string) *ShaderProgram {
	shaderProgram, err := NewShaderProgram(vertShader, fragShader)
	if err != nil {
		panic(err)
	}
	return shaderProgram
}

type getGlParam func(uint32, uint32, *int32)
type getInfoLog func(uint32, int32, *int32, *uint8)

func checkGlError(glObject uint32, errorParam uint32, getParamFn getGlParam,
	getInfoLogFn getInfoLog, failMsg string) {

	var success int32
	getParamFn(glObject, errorParam, &success)
	if success != 1 {
		var infoLog [512]byte
		getInfoLogFn(glObject, 512, nil, (*uint8)(unsafe.Pointer(&infoLog)))
		log.Fatalln(failMsg, "\n", string(infoLog[:512]))
	}
}

func checkShaderCompileErrors(shader uint32) {
	checkGlError(shader, gl.COMPILE_STATUS, gl.GetShaderiv, gl.GetShaderInfoLog,
		"ERROR::SHADER::COMPILE_FAILURE")
}

func checkProgramLinkErrors(program uint32) {
	checkGlError(program, gl.LINK_STATUS, gl.GetProgramiv, gl.GetProgramInfoLog,
		"ERROR::PROGRAM::LINKING_FAILURE")
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
	cvertexShaderSource, freeVertexShader := gl.Strs(vertexShaderSource + "\x00")
	defer freeVertexShader()
	gl.ShaderSource(vertexShader, 1, cvertexShaderSource, nil)
	gl.CompileShader(vertexShader)
	checkShaderCompileErrors(vertexShader)

	// Compile the fragment shader
	fragmentShader := gl.CreateShader(gl.FRAGMENT_SHADER)
	cfragmentShaderSource, freeFragmentShader := gl.Strs(fragmentShaderSource + "\x00")
	gl.ShaderSource(fragmentShader, 1, cfragmentShaderSource, nil)
	defer freeFragmentShader()
	gl.CompileShader(fragmentShader)
	checkShaderCompileErrors(fragmentShader)

	// Link the shader program
	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)
	checkProgramLinkErrors(program)

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return &ShaderProgram{program: program, UniformLocationCache: map[string]uint32{}, AttribLocationCache: map[string]uint32{}}, nil
}
