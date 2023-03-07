package engine

import (
	"github.com/go-gl/mathgl/mgl32"
	"io/ioutil"
	"math"
	"os"
)

func LoadFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func CenterCameraOnModel(vertices []mgl32.Vec3, camera *Camera) {
	// Compute the center point of the model
	var center mgl32.Vec3
	for _, v := range vertices {
		center = center.Add(v)
	}
	center = center.Mul(1.0 / float32(len(vertices)))

	// Compute the bounding box of the model
	var min, max mgl32.Vec3 = vertices[0], vertices[0]
	for _, v := range vertices {
		if v.X() < min.X() {
			min[0] = v.X()
		}
		if v.Y() < min.Y() {
			min[1] = v.Y()
		}
		if v.Z() < min.Z() {
			min[2] = v.Z()
		}
		if v.X() > max.X() {
			max[0] = v.X()
		}
		if v.Y() > max.Y() {
			max[1] = v.Y()
		}
		if v.Z() > max.Z() {
			max[2] = v.Z()
		}
	}

	// Compute the distance between the camera and the center of the bounding box
	dx := camera.Position.X() - center.X()
	dy := camera.Position.Y() - center.Y()
	dz := camera.Position.Z() - center.Z()
	distance := float32(math.Sqrt(float64(dx*dx + dy*dy + dz*dz)))

	// Move the camera to look at the center point of the model
	camera.Position = center.Add(mgl32.Vec3{distance, distance, distance})
	camera.Target = center
	camera.Up = mgl32.Vec3{0, 1, 0} // Adjust as needed
}
