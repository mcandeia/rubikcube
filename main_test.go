package main

import (
	"testing"
)

func TestRotate(t *testing.T) {
	c := NewSolvedCube()
	currentFrontColor := []Color{c.cube[FRONT][0][0], c.cube[FRONT][1][0], c.cube[FRONT][2][0]}
	currentLeftColor := []Color{c.cube[LEFT][0][0], c.cube[LEFT][1][0], c.cube[LEFT][2][0]}
	c.Rotate(0, 2) //rotate top on Y axis 2 time
	rotatedTopColor := []Color{c.cube[BACK][0][0], c.cube[BACK][1][0], c.cube[BACK][2][0]}
	// front goes to back 2 times on Y axis
	if currentFrontColor[0] != rotatedTopColor[0] || currentFrontColor[1] != rotatedTopColor[1] || currentFrontColor[2] != rotatedTopColor[2] {
		t.Fail()
	}
	// left should be transposed
	rotatedLeftColor := []Color{c.cube[LEFT][2][2], c.cube[LEFT][1][2], c.cube[LEFT][0][2]}
	if currentLeftColor[0] != rotatedLeftColor[0] || currentLeftColor[1] != rotatedLeftColor[1] || currentLeftColor[2] != rotatedLeftColor[2] {
		t.Fail()
	}
}
