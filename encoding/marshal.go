package encoding

import (
	"bytes"
	"fmt"
)

// MarshalBinary writes an obj to a gob
func (obj OBJ) MarshalBinary() ([]byte, error) {
	var b bytes.Buffer
	fmt.Fprintln(&b, "mtllib ", obj.filename)
	fmt.Fprintln(&b, "o ", obj.name)

	return b.Bytes(), nil
}
