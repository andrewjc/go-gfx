package engine

import "github.com/go-gl/gl/v4.1-core/gl"

type Mesh struct {
    Vertices   []float32
    TexCoords  []float32
    Normals    []float32
    Indices    []uint32
    Depth      float32
    Vao        uint32
    IndexCount int32
    Material   Material

    vertexBuffer   uint32
    texCoordBuffer uint32
    normalBuffer   uint32
    indexBuffer    uint32
}

func (m *Mesh) GetAttribSize(attr uint32) int32 {
    return 16
}

func NewMesh(combinedVertices []CombinedVertex, indices []uint32) *Mesh {
    vertexCount := len(combinedVertices)
    mesh := &Mesh{
        Vertices:  make([]float32, vertexCount*3),
        TexCoords: make([]float32, vertexCount*2),
        Normals:   make([]float32, vertexCount*3),
        Indices:   indices,

        Depth: 0.0,

        Vao:        0,
        IndexCount: int32(len(indices)),

        Material: Material{},
    }

    for i, cv := range combinedVertices {
        mesh.Vertices[i*3] = cv.Position.X()
        mesh.Vertices[i*3+1] = cv.Position.Y()
        mesh.Vertices[i*3+2] = cv.Position.Z()

        mesh.TexCoords[i*2] = cv.TexCoord.X()
        mesh.TexCoords[i*2+1] = cv.TexCoord.Y()

        mesh.Normals[i*3] = cv.Normal.X()
        mesh.Normals[i*3+1] = cv.Normal.Y()
        mesh.Normals[i*3+2] = cv.Normal.Z()
    }

    mesh.SetupGLBuffers()
    return mesh
}

func NewMeshNormalLines(mesh *Mesh, scale float32) *Mesh {
    vertexCount := len(mesh.Vertices) / 3
    normalLineVertices := make([]float32, vertexCount*2*3)

    for i := 0; i < vertexCount; i++ {
        normalLineVertices[i*6] = mesh.Vertices[i*3]
        normalLineVertices[i*6+1] = mesh.Vertices[i*3+1]
        normalLineVertices[i*6+2] = mesh.Vertices[i*3+2]

        normalLineVertices[i*6+3] = mesh.Vertices[i*3] + mesh.Normals[i*3]*scale
        normalLineVertices[i*6+4] = mesh.Vertices[i*3+1] + mesh.Normals[i*3+1]*scale
        normalLineVertices[i*6+5] = mesh.Vertices[i*3+2] + mesh.Normals[i*3+2]*scale
    }

    normalLineMesh := &Mesh{
        Vertices: normalLineVertices,
        Vao:      0,
    }
    normalLineMesh.SetupGLBuffers()

    return normalLineMesh
}

func (mesh *Mesh) SetupGLBuffers() {
    gl.GenVertexArrays(1, &mesh.Vao)
    gl.BindVertexArray(mesh.Vao)

    // Setup vertex buffer
    gl.GenBuffers(1, &mesh.vertexBuffer)
    gl.BindBuffer(gl.ARRAY_BUFFER, mesh.vertexBuffer)
    gl.BufferData(gl.ARRAY_BUFFER, len(mesh.Vertices)*4, gl.Ptr(&mesh.Vertices[0]), gl.STATIC_DRAW)

    // Set vertex attribute pointer for vertex positions
    gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))
    gl.EnableVertexAttribArray(0)

    // Setup texture coordinate buffer
    if mesh.TexCoords != nil {
        gl.GenBuffers(1, &mesh.texCoordBuffer)
        gl.BindBuffer(gl.ARRAY_BUFFER, mesh.texCoordBuffer)
        gl.BufferData(gl.ARRAY_BUFFER, len(mesh.TexCoords)*4, gl.Ptr(&mesh.TexCoords[0]), gl.STATIC_DRAW)
    }

    // Set vertex attribute pointer for texture coordinates
    gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 0, gl.PtrOffset(0))
    gl.EnableVertexAttribArray(1)

    // Setup normal buffer
    if mesh.Normals != nil {
        gl.GenBuffers(1, &mesh.normalBuffer)
        gl.BindBuffer(gl.ARRAY_BUFFER, mesh.normalBuffer)
        gl.BufferData(gl.ARRAY_BUFFER, len(mesh.Normals)*4, gl.Ptr(&mesh.Normals[0]), gl.STATIC_DRAW)
    }

    // Set vertex attribute pointer for normals
    gl.VertexAttribPointer(2, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))
    gl.EnableVertexAttribArray(2)

    // Setup index buffer

    if mesh.Indices != nil {
        gl.GenBuffers(1, &mesh.indexBuffer)
        gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, mesh.indexBuffer)
        gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(mesh.Indices)*4, gl.Ptr(&mesh.Indices[0]), gl.STATIC_DRAW)
    }

    // Unbind buffers and VAO
    gl.BindBuffer(gl.ARRAY_BUFFER, 0)
    gl.BindVertexArray(0)
}
