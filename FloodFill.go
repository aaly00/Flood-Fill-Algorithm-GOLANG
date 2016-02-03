package floodfill

import (
//"fmt"
)

const NORTH = 0
const EAST = 1
const SOUTH = 2
const WEST = 3

var neigbhoringCells [4][2]int = [4][2]int{[2]int{-1, 0}, [2]int{0, 1}, [2]int{1, 0}, [2]int{0, -1}}
var neigbhoringWalls [4][2]byte = [4][2]byte{[2]byte{0, 0}, [2]byte{0, 1}, [2]byte{1, 0}, [2]byte{0, 0}}

const ROW byte = 7
const COLUMN byte = 7

type maze struct {
	ROWS, COLUMNS, mouseRow, mouseColumn, mouseHeading, targetRow, targetColumn byte

	verticalWalls   [ROW][COLUMN + 1]bool
	horizontalWalls [ROW + 1][COLUMN]bool
	values          [ROW][COLUMN]byte
}

func (m *maze) constructor() *maze {
	m.ROWS = ROW
	m.COLUMNS = COLUMN
	m.mouseHeading = 0

	var i, j byte
	for i = 0; i < m.ROWS; i++ {
		for j = 0; j < m.COLUMNS+1; j++ {
			if j == 0 || j == m.COLUMNS {
				m.verticalWalls[i][j] = true
			}
		}
	}
	i = 0
	j = 0
	//initialize horizontalWalls (add exterior walls)
	for i = 0; i < m.ROWS+1; i++ {
		for j = 0; j < m.COLUMNS; j++ {
			if i == 0 || i == m.ROWS {
				m.horizontalWalls[i][j] = true
			}
		}
	}
	return m
}
func (m *maze) setTargetCell(x, y byte) {
	m.values[x][y] = 0
	m.targetRow = x
	m.targetColumn = y
}
func (m *maze) turnTowardBestNeighbor() {
	//Remember to replace this funtion since it is for debugging only
	//you need to implement a functional flow of control

	desiredHeading := findBestNeighbor()
	headingDifference := m.mouseHeading - desiredHeading

	switch headingDifference {
	//move Right
	case -1, 3:
		{
		}
		break

	//move Left
	case 1, -3:
		{
		}
		break

	//turn 180
	case 2, -2:
		{
		}
		break
	}

	mouseHeading = desiredHeading
}

func (m *maze) findBestNeighbor() byte {
	var valueBestNeighbor byte = 255
	var desiredHeading byte = NORTH
	for k := byte(0); k < 4; k++ {
		var neigbhoringCellRow byte = byte(int(m.mouseRow) + neigbhoringCells[k][0])
		var neigbhoringCellColumn byte = byte(int(m.mouseColumn) + neigbhoringCells[k][1])

		var neigbhoringWallRow byte = m.mouseRow + neigbhoringWalls[k][0]
		var neigbhoringWallColumn byte = m.mouseColumn + neigbhoringWalls[k][1]

		var wallExists bool = false

		if k == NORTH || k == SOUTH {
			wallExists = m.horizontalWalls[neigbhoringWallRow][neigbhoringWallColumn]
		} else { //Must be east or west since they are equal
			wallExists = m.verticalWalls[neigbhoringWallRow][neigbhoringWallColumn]
		}
		if (m.values[neigbhoringCellRow][neigbhoringCellColumn] <= valueBestNeighbor) && !(wallExists) {

			if m.values[neigbhoringCellRow][neigbhoringCellColumn] == valueBestNeighbor && mouseHeading != k {
				continue
			}
			valueBestNeighbor = m.values[neigbhoringCellRow][neigbhoringCellColumn]
			desiredHeading = k
		}
	}
	return desiredHeading
}
