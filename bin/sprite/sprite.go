package sprite

import (
	"fmt"
	"image"
	"image/color"
)

type Sprite [][]color.Color

func FromImageSection(image image.Image, startX, startY, spritesize int) Sprite {
	t := Sprite{}
	for x := 0; x < spritesize; x++ {
		t = append(t, []color.Color{})
		for y := 0; y < spritesize; y++ {
			t[x] = append(t[x], image.At(x+startX, y+startY))
		}
	}
	return t
}

func (s Sprite) Identical(b Sprite) bool {
	if len(s) != len(b) {
		return false
	}
	for x, row := range s {
		for y := range row {
			if s[x][y] != b[x][y] {
				return false
			}
		}
	}
	return true
}

func (s Sprite) IdenticalIfRotated(b Sprite) bool {
	copy := b.Copy()
	for i := 0; i < 4; i++ {
		if s.Identical(copy) {
			return true
		}

		copy.Rotate()
	}
	return false
}

func (s Sprite) IdenticalIfFlippedHorizontally(b Sprite) bool {
	copy := b.Copy()
	copy.FlipHorizontally()
	return s.Identical(copy)
}

func (s Sprite) IdenticalIfFlippedVertically(b Sprite) bool {
	copy := b.Copy()
	copy.FlipVertically()
	return s.Identical(copy)
}

// Prevents accidental mutation
// by copying the sprite into a
// new sprite. This as values
// of slices in Go are references.
func (s Sprite) Copy() Sprite {
	var copy [][]color.Color
	for x, row := range s {
		copy = append(copy, []color.Color{})
		for y := range row {
			copy[x] = append(copy[x], s[x][y])
		}
	}
	return copy
}

func (s Sprite) Rotate() {
	n := len(s)

	// Transpose.
	for i := 0; i < n-1; i++ {
		for ii := i; ii < n; ii++ {
			s[i][ii], s[ii][i] = s[ii][i], s[i][ii]
		}
	}

	// Flip.
	for _, row := range s {
		for i := 0; i < n/2; i++ {
			row[i], row[n-1-i] = row[n-1-i], row[i]
		}
	}
}

func (s Sprite) FlipHorizontally() {
	for _, row := range s {
		n := len(row)
		for i := 0; i < n/2; i++ {
			row[i], row[n-1-i] = row[n-1-i], row[i]
		}
	}
}

func (s Sprite) FlipVertically() {
	n := len(s)
	for i := 0; i < n/2; i++ {
		s[i], s[n-1-i] = s[n-1-i], s[i]
	}
}

// Generates an ID for the sprite by
// concatenating all values into a string.
func (s Sprite) ID() string {
	var id string
	for _, row := range s {
		for _, sprite := range row {
			id += fmt.Sprintf("%v_", sprite)
		}
	}
	return id[:len(id)-1]
}
