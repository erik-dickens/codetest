package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestShoot(t *testing.T) {
	ships := getShips()

	testCases := []testCase{
		{shots: []shot{{1, "G", false}, {1, "H", true}, {1, "I", false}}, expectedSunk: false},
		{shots: []shot{{1, "H", true}, {2, "H", true}, {3, "H", true}, {4, "H", true}}, expectedSunk: true},
		{shots: []shot{{1, "D", false}, {7, "F", true}, {8, "F", true}}, expectedSunk: true},
		{shots: []shot{{10, "D", true}, {9, "D", false}, {10, "C", false},
			{10, "E", true}, {10, "F", true}, {10, "G", true}, {10, "H", true}},
			expectedSunk: true},
		{shots: []shot{{7, "J", true}, {8, "I", false}, {9, "H", false}, {10, "G", true}}, expectedSunk: false},
		{shots: []shot{{1, "H", true}, {8, "I", false}, {10, "CC", false}, {10, "G", true}}, expectedSunk: false, err: ErrIncorrectLetter},
		{shots: []shot{{1, "G", false}, {1, "H", true}, {11, "G", false}}, expectedSunk: false, err: ErrOutOfGridBoundaries},
		{shots: []shot{{1, "G", false}, {1, "H", true}, {10, "P", false}}, expectedSunk: false, err: ErrOutOfGridBoundaries},
		{shots: []shot{{1, "G", false}, {12, "K", false}, {11, "G", false}}, expectedSunk: false, err: ErrOutOfGridBoundaries},
	}
	for ind, test := range testCases {
		t.Run(fmt.Sprint(ind), func(t *testing.T) {
			grid := NewGrid(ships)
			var latestShootResult ShootResult
			for _, shot := range test.shots {
				var err error
				latestShootResult, err = grid.Shoot(shot.num, shot.letter)
				if !cmp.Equal(shot.expectedHit, latestShootResult.Hit) {
					t.Log(cmp.Diff(shot.expectedHit, latestShootResult.Hit))
					t.Fail()
				}
				if err != nil && test.err != nil {
					if !errors.Is(err, test.err) {
						t.Log("err is incorrect")
						t.Fail()
					}
				}
			}

			if !cmp.Equal(test.expectedSunk, latestShootResult.Sunk) {
				t.Log(cmp.Diff(test.expectedSunk, latestShootResult.Sunk))
				t.Fail()
			}

			grid.ResetShips()
		})
	}
}
