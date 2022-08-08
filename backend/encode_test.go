package backend

import (
	"fmt"
	"testing"
)

func TestEncodeString(t *testing.T) {
	str := `abc"\\n<😁😁😁世界"`

	encoder := NewEncoder(WithEscapeUnicode(), WithEscapeHTML())
	defer ReleaseEncoder(encoder)

	encoder.EncodeString(str)

	fmt.Println(encoder.String())
	fmt.Println(encoder.Bytes())
}
