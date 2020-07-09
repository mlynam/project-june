package util

import "github.com/go-gl/mathgl/mgl32"

// Vec3Convert64To32 converts a float64 vec3 to a float32 vec3
func Vec3Convert64To32(v [3]float64) [3]float32 {
	return [3]float32{float32(v[0]), float32(v[1]), float32(v[2])}
}

// Vec4Convert64to32 converts a float64 vec4 to a float32 vec4
func Vec4Convert64to32(v [4]float64) [4]float32 {
	return [4]float32{float32(v[0]), float32(v[1]), float32(v[2]), float32(v[3])}
}

// Vec4Convert64ToQuat converts a [4]float64 to a Quaternion
func Vec4Convert64ToQuat(v [4]float64) mgl32.Quat {
	var vec [3]float64
	vec[0] = v[0]
	vec[1] = v[1]
	vec[2] = v[2]
	return mgl32.Quat{
		W: float32(v[3]),
		V: Vec3Convert64To32(vec),
	}
}
