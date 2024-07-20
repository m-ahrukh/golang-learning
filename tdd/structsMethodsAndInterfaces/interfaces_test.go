package structsmethodsandinterfaces

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	assert.Equal(t, got, want)
}

func TestArea(t *testing.T) {
	got := Area(12.0, 6.0)
	want := 72.0

	assert.Equal(t, got, want)
}
*/

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	assert.Equal(t, want, got)
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{12.0, 6.0}
	got := rectangle.Area()
	want := 72.0

	assert.Equal(t, want, got)
}

func TestAreas(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		got := rectangle.Area()
		want := 72.0

		assert.Equal(t, want, got)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		assert.Equal(t, want, got)
	})
}
