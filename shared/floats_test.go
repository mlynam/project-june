package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVec3Convert64To32(t *testing.T) {
	v := [3]float64{1.0, 1.0, 1.0}

	converted := Vec3Convert64To32(v)

	assert.Equal(t, [3]float32{1.0, 1.0, 1.0}, converted)
}

func TestVec4Convert64To32(t *testing.T) {
	v := [4]float64{1.0, 1.0, 1.0, 1.0}

	converted := Vec4Convert64to32(v)

	assert.Equal(t, [4]float32{1.0, 1.0, 1.0, 1.0}, converted)
}
