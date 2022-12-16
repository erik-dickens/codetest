package main

import (
	"errors"
)

type Coordinate struct {
	Num    int
	Letter rune
}

type Position struct {
	Start Coordinate
	End   Coordinate
}

var (
	ErrIncorrectLetter     = errors.New("incorrect input in string")
	ErrOutOfGridBoundaries = errors.New("shot out of grid boundaries")
)

var GridSize = Position{Start: Coordinate{Num: 1, Letter: 'A'}, End: Coordinate{Num: 10, Letter: 'J'}}

type Grid struct {
	ships []*ship
	Shots int
}

type ship struct {
	start Coordinate
	end   Coordinate
}

type ShootResult struct {
	Hit  bool
	Sunk bool
}

func NewGrid(ships []Position) *Grid {
	var realShips []*ship
	for _, p := range ships {
		ship := ship{
			start: p.Start,
			end:   p.End,
		}
		realShips = append(realShips, &ship)
	}

	grid := Grid{ships: realShips}
	return &grid
}
func (grid *Grid) Shoot(shotNum int, shotLetter string) (ShootResult, error) {
	//TODO: implement here
	return ShootResult{}, nil
}

func (grid *Grid) ResetShips() {
	grid.Shots = 0
}

type shot struct {
	num         int
	letter      string
	expectedHit bool
}

type testCase struct {
	shots        []shot
	expectedSunk bool
	err          error
}

func getShips() []Position {
	// count  name              size
	//   1    Aircraft Carrier   5
	//   1    Battleship         4
	//   1    Cruiser            3
	//   2    Destroyer          2
	//   2    Submarine          1
	//
	// 		  A B C D E F G H I J
	//		1               @
	//		2 @             @
	//		3         @     @
	//		4               @
	//		5   @ @
	//		6
	//		7           @       @
	//		8           @       @
	//		9                   @
	//	   10       @ @ @ @ @
	//
	var ships []Position
	ships = append(ships, Position{
		Start: Coordinate{2, 'A'},
		End:   Coordinate{2, 'A'},
	})
	ships = append(ships, Position{
		Start: Coordinate{3, 'E'},
		End:   Coordinate{3, 'E'},
	})
	ships = append(ships, Position{
		Start: Coordinate{1, 'H'},
		End:   Coordinate{4, 'H'},
	})
	ships = append(ships, Position{
		Start: Coordinate{5, 'B'},
		End:   Coordinate{5, 'C'},
	})
	ships = append(ships, Position{
		Start: Coordinate{7, 'F'},
		End:   Coordinate{8, 'F'},
	})
	ships = append(ships, Position{
		Start: Coordinate{7, 'I'},
		End:   Coordinate{9, 'I'},
	})
	ships = append(ships, Position{
		Start: Coordinate{10, 'D'},
		End:   Coordinate{10, 'H'},
	})

	return ships
}
