package engine

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func LoadObjFile(filename string) (*ComplexMesh, error) {
	// Get parent directory of filename
	parentDir := filepath.Dir(filename)

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create arrays to hold the vertex, normal, and texture coordinate data
	var vertices []Vec3
	var normals []Vec3
	var texCoords []Vec2

	// Create arrays to hold the vertex, normal, and texture coordinate indices
	var vertexIndices []int
	var normalIndices []int
	var texCoordIndices []int

	// Create a map to hold the vertex data (to avoid duplicates)
	vertexMap := make(map[string]int)

	// Create a map to hold the material data
	materialMap := make(map[string]*BasicMaterial)

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line into tokens
		tokens := strings.Fields(line)

		if len(tokens) > 0 {
			switch tokens[0] {
			case "#":
				// Comment
				log.Println("Comment:", line)
			case "mtllib":
				// Material library
				log.Println("Material library:", line)

				materials, err := parseMaterialLibrary(filepath.Join(parentDir, tokens[1]))
				if err != nil {
					return nil, err
				}

				for _, material := range materials {
					materialMap[material.Name] = material
				}

			case "o":
				// Object name
				log.Println("Object name:", line)

			case "g":
				// Group name
				log.Println("Group name:", line)
			case "s":
				// Smoothing group
				log.Println("Smoothing group:", line)
			case "usemtl":
				// Material name
				log.Println("Material name:", line)
				//materialName := tokens[1]
			case "v":
				// Vertex data
				x, _ := strconv.ParseFloat(tokens[1], 32)
				y, _ := strconv.ParseFloat(tokens[2], 32)
				z, _ := strconv.ParseFloat(tokens[3], 32)
				vertex := Vec3{float64(float32(x)), float64(float32(y)), float64(float32(z))}

				// Add the vertex to the vertex array (if it doesn't already exist)
				key := fmt.Sprintf("%.6f,%.6f,%.6f", vertex.X, vertex.Y, vertex.Z)
				index, ok := vertexMap[key]
				if !ok {
					vertices = append(vertices, vertex)
					index = len(vertices) - 1
					vertexMap[key] = index
				}

				// Add the vertex index to the vertex index array
				vertexIndices = append(vertexIndices, index)
			case "vn":
				// Normal data
				x, _ := strconv.ParseFloat(tokens[1], 32)
				y, _ := strconv.ParseFloat(tokens[2], 32)
				z, _ := strconv.ParseFloat(tokens[3], 32)
				normal := Vec3{float64(float32(x)), float64(float32(y)), float64(float32(z))}

				// Add the normal to the normal array
				normals = append(normals, normal)
			case "vt":
				// Texture coordinate data
				u, _ := strconv.ParseFloat(tokens[1], 32)
				v, _ := strconv.ParseFloat(tokens[2], 32)
				texCoord := Vec2{float64(float32(u)), float64(float32(v))}

				// Add the texture coordinate to the texture coordinate array
				texCoords = append(texCoords, texCoord)
			case "f":
				// Face data
				for i := 1; i < len(tokens); i++ {
					// Split the face token into its vertex, texture coordinate, and normal indices
					faceTokens := strings.Split(tokens[i], "/")
					vertexIndex, _ := strconv.Atoi(faceTokens[0])
					texCoordIndex, _ := strconv.Atoi(faceTokens[1])
					normalIndex, _ := strconv.Atoi(faceTokens[2])

					// Add the vertex, texture coordinate, and normal indices to their respective arrays
					vertexIndices = append(vertexIndices, vertexIndex-1)
					texCoordIndices = append(texCoordIndices, texCoordIndex-1)
					normalIndices = append(normalIndices, normalIndex-1)
				}
			default:
				log.Fatal("Unknown token: " + tokens[0])
			}
		}
	}

	// Create arrays to hold the final vertex, normal, and texture coordinate data
	var finalVertices []float32
	var finalNormals []float32
	var finalTexCoords []float32

	// Create a map to hold the final vertex data (to avoid duplicates)
	finalVertexMap := make(map[string]int)

	// Create arrays to hold the final vertex, normal, and texture coordinate indices
	var finalVertexIndices []uint32
	var finalNormalIndices []uint32
	var finalTexCoordIndices []uint32

	// Loop through the vertex indices
	for i := 0; i < len(vertexIndices); i++ {
		// Get the vertex index
		vertexIndex := vertexIndices[i]

		// Get the vertex
		vertex := vertices[vertexIndex]

		// Add the vertex to the final vertex array (if it doesn't already exist)
		key := fmt.Sprintf("%.6f,%.6f,%.6f", vertex.X, vertex.Y, vertex.Z)
		index, ok := finalVertexMap[key]
		if !ok {
			finalVertices = append(finalVertices, float32(vertex.X), float32(vertex.Y), float32(vertex.Z))
			index = len(finalVertices)/3 - 1
			finalVertexMap[key] = index
		}

		// Add the vertex index to the final vertex index array
		finalVertexIndices = append(finalVertexIndices, uint32(index))

		// Get the normal index
		normalIndex := normalIndices[i]

		// Get the normal
		normal := normals[normalIndex]

		// Add the normal to the final normal array
		finalNormals = append(finalNormals, float32(normal.X), float32(normal.Y), float32(normal.Z))

		// Add the normal index to the final normal index array
		finalNormalIndices = append(finalNormalIndices, uint32(index))

		// Get the texture coordinate index
		texCoordIndex := texCoordIndices[i]

		// Get the texture coordinate
		texCoord := texCoords[texCoordIndex]

		// Add the texture coordinate to the final texture coordinate array
		finalTexCoords = append(finalTexCoords, float32(texCoord.X), float32(texCoord.Y))

		// Add the texture coordinate index to the final texture coordinate index array
		finalTexCoordIndices = append(finalTexCoordIndices, uint32(index))

	}

	textureHandle := 0

	mesh := NewComplexMesh(finalVertices, finalVertexIndices, finalTexCoordIndices, uint32(textureHandle))

	// Return the mesh
	return mesh, nil
}

func parseMaterialLibrary(libraryFile string) (map[string]*BasicMaterial, error) {
	parentDir := filepath.Dir(libraryFile)

	file, err := os.Open(libraryFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	materials := make(map[string]*BasicMaterial)
	var currentMaterial BasicMaterial

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) == 0 {
			continue // Skip empty lines
		}

		switch fields[0] {
		case "#":
			// Comment
		case "newmtl":
			if len(fields) != 2 {
				return nil, fmt.Errorf("invalid newmtl line: %v", line)
			}
			newMaterial := NewBasicMaterial()
			newMaterial.Name = fields[1]

			currentMaterial = *newMaterial
			materials[newMaterial.Name] = &currentMaterial
			// Parse other material properties as needed...
		case "Kd":
			if len(fields) != 4 {
				return nil, fmt.Errorf("invalid kd line: %v", line)
			}
			r, _ := strconv.ParseFloat(fields[1], 32)
			g, _ := strconv.ParseFloat(fields[2], 32)
			b, _ := strconv.ParseFloat(fields[3], 32)
			currentMaterial.DiffuseColor = Vec3{float64(float32(r)), float64(float32(g)), float64(float32(b))}
		case "map_Kd":
			if len(fields) != 2 {
				return nil, fmt.Errorf("invalid map_Kd line: %v", line)
			}
			textureFile := fields[1]

			textureFile = strings.ReplaceAll(textureFile, "\\", string(filepath.Separator))
			_, filePath := filepath.Split(textureFile)

			currentMaterial.DiffuseTexturePath = filepath.Join(parentDir, filePath)
		default:
			log.Fatal("Unknown token: " + fields[0])
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return materials, nil
}
