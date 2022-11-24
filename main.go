/*
Rubik's cube rotation adhoc version

Limitations:

1) No support for rotating in another direction than inverse clock
*/
package main

import "fmt"

type Color int

const (
	RED Color = iota
	BLUE
	GREEN
	WHITE
	YELLOW
	BLACK
)

type Face int

const (
	TOP Face = iota
	BOTTOM
	LEFT
	RIGHT
	FRONT
	BACK
	NONE
)

type Axis string

const (
	X Axis = "X"
	Y Axis = "Y"
)

type Cell [2]int

type Cube struct {
	cube map[Face][][]Color
}

func transpose(face [][]Color) {
	for i, j := 0, len(face)-1; i < j; i, j = i+1, j-1 {
		face[i], face[j] = face[j], face[i]
	}

	for i := 0; i < len(face); i++ {
		for j := 0; j < i; j++ {
			face[i][j], face[j][i] = face[j][i], face[i][j]
		}
	}
}

// orientationShift given a face which is the next face that should be used to shift the colors
var orientationShift = map[Axis]map[Face]Face{
	X: {
		FRONT: LEFT,
		RIGHT: FRONT,
		BACK:  RIGHT,
		LEFT:  BACK,
	},
	Y: {
		FRONT:  TOP,
		BOTTOM: FRONT,
		BACK:   BOTTOM,
		TOP:    BACK,
	},
}

// Movement a movement is basically a hint on how to move a kube for one movement.
type Movement struct {
	// shift points to the affected cells for a specific operation
	shift [3]Cell
	// transpose points which face should be transposed after the shift operation
	// there are movements that don't require transpose (i.e mid slices)
	transpose Face
	// orientation points out the start face for the operation
	orientation Face
	// axis Y or X
	axis Axis
}

func (m Movement) Apply(c *Cube) {
	face := c.cube[m.orientation]
	current := make([]Color, 3)
	for idx, cell := range m.shift {
		current[idx] = face[cell[0]][cell[1]]
	}
	shifter := orientationShift[m.axis][m.orientation]
	// shift until get the initial orientation
	for shifter != m.orientation {
		temp := make([]Color, 3)
		for idx, cell := range m.shift {
			temp[idx] = c.cube[shifter][cell[0]][cell[1]]
			c.cube[shifter][cell[0]][cell[1]] = current[idx]
		}
		current = temp
		shifter = orientationShift[m.axis][shifter]
	}

	face, ok := c.cube[m.transpose]
	if ok {
		transpose(face)
	}
}

var movements = map[int]Movement{
	0: {
		shift:       [3]Cell{{0, 0}, {1, 0}, {2, 0}},
		transpose:   LEFT,
		orientation: FRONT,
		axis:        Y,
	},
	1: {
		shift:       [3]Cell{{0, 0}, {0, 1}, {0, 2}},
		transpose:   TOP,
		orientation: FRONT,
		axis:        X,
	},
	2: {
		shift:       [3]Cell{{1, 0}, {1, 1}, {1, 2}},
		transpose:   NONE,
		orientation: FRONT,
		axis:        X,
	},
	3: {
		shift:       [3]Cell{{0, 2}, {1, 2}, {2, 2}},
		transpose:   RIGHT,
		orientation: FRONT,
		axis:        Y,
	},
	4: {
		shift:       [3]Cell{{1, 2}, {1, 1}, {1, 0}},
		transpose:   NONE,
		orientation: FRONT,
		axis:        X,
	},
	5: {
		shift:       [3]Cell{{2, 2}, {2, 1}, {2, 0}},
		transpose:   BOTTOM,
		orientation: FRONT,
		axis:        X,
	},
	6: {
		shift:       [3]Cell{{0, 1}, {1, 1}, {2, 1}},
		transpose:   NONE,
		orientation: RIGHT,
		axis:        Y,
	},
	7: {
		shift:       [3]Cell{{0, 2}, {1, 2}, {2, 2}},
		transpose:   BACK,
		orientation: RIGHT,
		axis:        Y,
	},
	8: {
		shift:       [3]Cell{{0, 2}, {1, 2}, {2, 2}},
		transpose:   FRONT,
		orientation: RIGHT,
		axis:        Y,
	},
}

func (c *Cube) Rotate(slice int, op int) {
	moves := movements[slice]
	for i := 0; i < op; i++ {
		moves.Apply(c)
	}
}

func NewSolvedCube() *Cube {
	c := &Cube{
		cube: map[Face][][]Color{},
	}

	for face, color := range map[Face]Color{TOP: RED, BOTTOM: GREEN, FRONT: YELLOW, BACK: WHITE, LEFT: BLUE, RIGHT: BLACK} {
		c.cube[face] = make([][]Color, 3)
		for i := 0; i < 3; i++ {
			c.cube[face][i] = make([]Color, 3)
			for j := 0; j < 3; j++ {
				c.cube[face][i][j] = color
			}
		}
	}
	return c
}

func main() {
	c := NewSolvedCube()
	fmt.Println(c.cube)
	c.Rotate(0, 1)
	fmt.Println(c.cube)
}
