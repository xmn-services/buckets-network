package math

import "fmt"

// String returns the string representation of the vector
func (obj Vec2) String() string {
	return fmt.Sprintf("%f, %f", obj[0], obj[1])
}
