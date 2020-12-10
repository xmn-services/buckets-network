package math

import "fmt"

// String returns the string representation of the vector
func (obj *Vec3) String() string {
	return fmt.Sprintf("%f, %f, %f", obj[0], obj[1], obj[2])
}
