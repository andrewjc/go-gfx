package engine

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Vertex struct {
	x, y, z float32
}

func (v Vertex) ToVec3() mgl32.Vec3 {
	return mgl32.Vec3{v.x, v.y, v.z}
}

type TexCoord struct {
	u, v float32
}

func (v TexCoord) ToVec2() mgl32.Vec2 {
	return mgl32.Vec2{v.u, v.v}
}

type Normal struct {
	x, y, z float32
}

func (v Normal) ToVec3() mgl32.Vec3 {
	return mgl32.Vec3{v.x, v.y, v.z}
}

type FaceVertex struct {
	vertexIndex   int
	texCoordIndex int
	normalIndex   int
	smooth        bool
}

type ImportedModel struct {
	Objects         []*ImportedMeshObj
	materialLibrary map[string]material
}

type ImportedMeshObj struct {
	Name        string
	Vertices    []Vertex
	TexCoords   []TexCoord
	Normals     []Normal
	FaceIndices []FaceVertex
}

type material struct {
	name      string
	ambient   color.RGBA
	diffuse   color.RGBA
	specular  color.RGBA
	emissive  mgl32.Vec3
	shininess float32
	texture   uint32
}

func (o *ImportedMeshObj) AddVertex(v Vertex) {
	o.Vertices = append(o.Vertices, v)
}

func (o *ImportedMeshObj) AddTexCoord(t TexCoord) {
	o.TexCoords = append(o.TexCoords, t)
}

func (o *ImportedMeshObj) AddNormal(n Normal) {
	o.Normals = append(o.Normals, n)
}

func (o *ImportedMeshObj) AddFaceVertex(v FaceVertex) {
	o.FaceIndices = append(o.FaceIndices, v)
}

func LdrParseObj(filePath string) (*ImportedModel, error) {

	fileContents, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	folderRootPathSplit := strings.Split(strings.ReplaceAll(filePath, "\\", "/"), "/")
	folderRootPath := folderRootPathSplit[:len(folderRootPathSplit)-1][0]

	model := &ImportedModel{
		Objects: make([]*ImportedMeshObj, 0),
	}
	var currentObject *ImportedMeshObj
	model.Objects = append(model.Objects, currentObject)

	scanner := bufio.NewScanner(bytes.NewReader(fileContents))
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		switch fields[0] {
		case "o":
			if len(fields) < 2 {
				return nil, errors.New("malformed obj file, missing object name")
			}
			currentObject = &ImportedMeshObj{
				Name: fields[1],
			}
			model.Objects = append(model.Objects, currentObject)
		case "v":
			v, err := LdrParseVertex(line)
			if err != nil {
				return nil, err
			}
			currentObject.Vertices = append(currentObject.Vertices, v)
		case "vt":
			t, err := LdrParseTexCoord(line)
			if err != nil {
				return nil, err
			}
			currentObject.TexCoords = append(currentObject.TexCoords, t)
		case "vn":
			n, err := LdrParseNormal(line)
			if err != nil {
				return nil, err
			}
			currentObject.Normals = append(currentObject.Normals, n)
		case "f":
			faceIndices := make([]int, 0)
			for _, field := range fields[1:] {
				f, err := LdrParseFaceVertex(field)
				if err != nil {
					return nil, fmt.Errorf("could not parse face Vertex: %v", err)
				}
				currentObject.AddFaceVertex(f)
				faceIndices = append(faceIndices, len(currentObject.FaceIndices)-1)
			}
		case "s":
			continue
			/*if len(fields) > 1 && fields[1] == "1" {
			  	smooth = true
			  } else {
			  	smooth = false
			  }*/
		/*case "usemtl":
		  materialName, err := LdrParseusemtl(line)
		  if err != nil {
		  	return nil, err
		  }
		  currentObject.MaterialName = materialName

		  // Add face material
		  faceIndices := make([]int, 0)
		  for i := 0; i < len(currentObject.FaceIndices); i++ {
		  	faceIndices = append(faceIndices, i)
		  }
		  currentObject.AddFaceMaterial(materialName, faceIndices)*/
		case "mtllib":
			if len(fields) < 2 {
				return nil, errors.New("malformed obj file, missing material library name")
			}
			materialMap, err := LdrParseMtlLib(filepath.Join(folderRootPath, fields[1]))
			if err != nil {
				return nil, err
			}

			model.materialLibrary = materialMap
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Center the model
	for _, mesh := range model.Objects {
		if mesh == nil {
			continue
		}
		var totalX, totalY, totalZ float32
		for _, vertex := range mesh.Vertices {
			totalX += vertex.x
			totalY += vertex.y
			totalZ += vertex.z
		}
		centerX := totalX / float32(len(mesh.Vertices))
		centerY := totalY / float32(len(mesh.Vertices))
		centerZ := totalZ / float32(len(mesh.Vertices))
		for i := range mesh.Vertices {
			mesh.Vertices[i].x -= centerX
			mesh.Vertices[i].y -= centerY
			mesh.Vertices[i].z -= centerZ
		}
	}
	return model, nil
}

func LdrParseMtlLib(path string) (map[string]material, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	mtls := make(map[string]material)
	var curMtl *material

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "newmtl ") {
			name := strings.TrimSpace(line[7:])
			curMtl = &material{}
			mtls[name] = *curMtl
		} else if curMtl != nil {
			fields := strings.Fields(line)
			if len(fields) > 0 {
				switch fields[0] {
				case "Ka":
					c, err := LdrParseColor(line)
					if err != nil {
						return nil, fmt.Errorf("could not parse ambient color for material %q: %v", curMtl.name, err)
					}
					curMtl.ambient = c
				case "Kd":
					c, err := LdrParseColor(line)
					if err != nil {
						return nil, fmt.Errorf("could not parse diffuse color for material %q: %v", curMtl.name, err)
					}
					curMtl.diffuse = c
				case "Ks":
					c, err := LdrParseColor(line)
					if err != nil {
						return nil, fmt.Errorf("could not parse specular color for material %q: %v", curMtl.name, err)
					}
					curMtl.specular = c
				case "Ns":
					ns, err := strconv.ParseFloat(fields[1], 32)
					if err != nil {
						return nil, fmt.Errorf("could not parse specular exponent for material %q: %v", curMtl.name, err)
					}
					curMtl.shininess = float32(ns)
				case "map_Kd":
					texturePathTemp := strings.Join(fields[1:], " ")

					texturePathTemp = strings.Replace(texturePathTemp, "\\", "/", -1)
					textureFileTemp := strings.Split(texturePathTemp, "/")

					texturePath := filepath.Join(filepath.Dir(path), textureFileTemp[len(textureFileTemp)-1])

					texture, err := loadImage(texturePath)
					if err != nil {
						return nil, fmt.Errorf("could not load texture image for material %q: %v", curMtl.name, err)
					}
					curMtl.texture = texture
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return mtls, nil
}

func LdrParseColor(line string) (color.RGBA, error) {
	fields := strings.Fields(line)[1:]
	if len(fields) != 3 {
		return color.RGBA{}, fmt.Errorf("expected 3 fields in color, found %d", len(fields))
	}
	r, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return color.RGBA{}, fmt.Errorf("could not parse color red component: %v", err)
	}
	g, err := strconv.ParseFloat(fields[1], 64)
	if err != nil {
		return color.RGBA{}, fmt.Errorf("could not parse color green component: %v", err)
	}
	b, err := strconv.ParseFloat(fields[2], 64)
	if err != nil {
		return color.RGBA{}, fmt.Errorf("could not parse color blue component: %v", err)
	}
	return color.RGBA{uint8(r * 255), uint8(g * 255), uint8(b * 255), 255}, nil
}

func loadImage(texturePath string) (uint32, error) {
	file, err := os.Open(texturePath)
	if err != nil {
		return 0, fmt.Errorf("failed to open texture file: %w", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return 0, fmt.Errorf("failed to decode texture file: %w", err)
	}

	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{}, draw.Src)

	var textureID uint32
	gl.GenTextures(1, &textureID)

	gl.BindTexture(gl.TEXTURE_2D, textureID)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix),
	)

	return textureID, nil
}

func LdrParseusemtl(line string) (string, error) {
	fields := strings.Fields(line)[1:]
	if len(fields) != 1 {
		return "", fmt.Errorf("expected 1 field in usemtl, found %d", len(fields))
	}
	return fields[0], nil
}

func LdrParseVertex(line string) (Vertex, error) {
	fields := strings.Fields(line)[1:]
	if len(fields) != 3 {
		return Vertex{}, fmt.Errorf("expected 3 fields in Vertex, found %d", len(fields))
	}
	x, err := strconv.ParseFloat(fields[0], 32)
	if err != nil {
		return Vertex{}, fmt.Errorf("could not parse Vertex x coordinate: %v", err)
	}
	y, err := strconv.ParseFloat(fields[1], 32)
	if err != nil {
		return Vertex{}, fmt.Errorf("could not parse Vertex y coordinate: %v", err)
	}
	z, err := strconv.ParseFloat(fields[2], 32)
	if err != nil {
		return Vertex{}, fmt.Errorf("could not parse Vertex z coordinate: %v", err)
	}
	return Vertex{float32(x), float32(y), float32(z)}, nil
}

func LdrParseTexCoord(line string) (TexCoord, error) {
	fields := strings.Fields(line)[1:]
	if len(fields) != 2 {
		return TexCoord{}, fmt.Errorf("expected 2 fields in texture coordinate, found %d", len(fields))
	}
	u, err := strconv.ParseFloat(fields[0], 32)
	if err != nil {
		return TexCoord{}, fmt.Errorf("could not parse texture coordinate u value: %v", err)
	}
	v, err := strconv.ParseFloat(fields[1], 32)
	if err != nil {
		return TexCoord{}, fmt.Errorf("could not parse texture coordinate v value: %v", err)
	}
	return TexCoord{float32(u), float32(v)}, nil
}

func LdrParseNormal(line string) (Normal, error) {
	fields := strings.Fields(line)[1:]
	if len(fields) != 3 {
		return Normal{}, fmt.Errorf("expected 3 fields in Normal, found %d", len(fields))
	}
	x, err := strconv.ParseFloat(fields[0], 32)
	if err != nil {
		return Normal{}, fmt.Errorf("could not parse Normal x component: %v", err)
	}
	y, err := strconv.ParseFloat(fields[1], 32)
	if err != nil {
		return Normal{}, fmt.Errorf("could not parse Normal y component: %v", err)
	}
	z, err := strconv.ParseFloat(fields[2], 32)
	if err != nil {
		return Normal{}, fmt.Errorf("could not parse Normal z component: %v", err)
	}
	return Normal{float32(x), float32(y), float32(z)}, nil
}

func LdrParseFaceVertex(field string) (FaceVertex, error) {
	faceVertex := FaceVertex{}

	indices := strings.Split(field, "/")
	switch len(indices) {
	case 1:
		// Vertex index
		vertexIndex, err := strconv.Atoi(indices[0])
		if err != nil {
			return faceVertex, err
		}
		faceVertex.vertexIndex = vertexIndex - 1
	case 2:
		// Vertex index / texture coordinate index
		vertexIndex, err := strconv.Atoi(indices[0])
		if err != nil {
			return faceVertex, err
		}
		faceVertex.vertexIndex = vertexIndex - 1

		texCoordIndex, err := strconv.Atoi(indices[1])
		if err != nil {
			return faceVertex, err
		}
		faceVertex.texCoordIndex = texCoordIndex - 1
	case 3:
		// Vertex index / texture coordinate index / Normal index
		vertexIndex, err := strconv.Atoi(indices[0])
		if err != nil {
			return faceVertex, err
		}
		faceVertex.vertexIndex = vertexIndex - 1

		if len(indices[1]) > 0 {
			texCoordIndex, err := strconv.Atoi(indices[1])
			if err != nil {
				return faceVertex, err
			}
			faceVertex.texCoordIndex = texCoordIndex - 1
		}

		normalIndex, err := strconv.Atoi(indices[2])
		if err != nil {
			return faceVertex, err
		}
		faceVertex.normalIndex = normalIndex - 1
	default:
		return faceVertex, fmt.Errorf("invalid number of indices for face Vertex: %s", field)
	}

	return faceVertex, nil
}
