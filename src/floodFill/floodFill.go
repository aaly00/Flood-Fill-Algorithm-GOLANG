package floodFill

import "fmt"

const NORTH = 0
const EAST = 1
const SOUTH = 2
const WEST = 3

var neigbhoringCells [4][2]int = [4][2]int{[2]int{-1, 0}, [2]int{0, 1}, [2]int{1, 0}, [2]int{0, -1}}
var neigbhoringWalls [4][2]byte = [4][2]byte{[2]byte{0, 0}, [2]byte{0, 1}, [2]byte{1, 0}, [2]byte{0, 0}}

const ROW byte = 7
const COLUMN byte = 7

type Maze struct {
	ROWS, COLUMNS, mouseRow, mouseColumn, MouseHeading, targetRow, targetColumn byte

	//	verticalWalls   [ROW][COLUMN + 1]bool
	verticalWalls   [][]bool
	horizontalWalls [][]bool
	//	horizontalWalls [ROW + 1][COLUMN]bool
	values [][]byte
	//	values          [ROW][COLUMN]byte
}

func (m Maze) Constructor(ROW, COLUMN,Tr,Tc byte) Maze{
	m.ROWS = ROW
	m.COLUMNS = COLUMN
	m.MouseHeading = 0
	m.targetRow=Tr;
	m.targetColumn=Tc;
	m.verticalWalls = Create2dSliceBool(ROW+1, COLUMN+2)
	m.horizontalWalls = Create2dSliceBool(ROW+2, COLUMN+1)
	m.values = Create2dSliceByte(ROW+1, COLUMN+1)

//	fmt.Println(m.horizontalWalls[3][4])	
	for i := byte(0); i <ROW; i++ {
		for j := byte(0); j < COLUMN+1; j++ {
			if j == 0 || j == COLUMN {
				m.verticalWalls[i][j] = true
			}
		}
	}

	//initialize horizontalWalls (add exterior walls)
	for i := byte(0); i <ROW+1; i++ {
		for j := byte(0); j < COLUMN; j++ {
			if i == 0 || i == ROW {
				m.horizontalWalls[i][j] = true
			}
		}
	}
	return m;
}
func (m *Maze) SetTargetCell(x, y byte) {
	m.values[x][y] = 0
	m.targetRow = x
	m.targetColumn = y
}
func (m *Maze) TurnTowardBestNeighbor() {
	//Remember to replace this funtion since it is for debugging only
	//you need to implement a functional flow of control

	desiredHeading := int(m.FindBestNeighbor())
	headingDifference := int(m.MouseHeading) - desiredHeading

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

	m.MouseHeading = byte(desiredHeading)
}

func (m *Maze) FindBestNeighbor() byte {
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

			if m.values[neigbhoringCellRow][neigbhoringCellColumn] == valueBestNeighbor && m.MouseHeading != k {
				continue
			}
			valueBestNeighbor = m.values[neigbhoringCellRow][neigbhoringCellColumn]
			desiredHeading = k
		}
	}
	return desiredHeading
}

func Create2dSliceByte(x, y byte) [][]byte {
	slice := make([][]byte, y)
	for i := range slice {
		slice[i] = make([]byte, x)
	}
	return slice
}

func Create2dSliceBool(x, y byte) [][]bool {
	slice := make([][]bool, y)
	for i := range slice {
		slice[i] = make([]bool, x)
	}
	return slice
}
func Create2dSliceInt(x, y byte) [][]int {
	slice := make([][]int, y)
	for i := range slice {
		slice[i] = make([]int, x)
	}
	return slice
}

func (m *Maze) Print() {
	
	for  i := byte(0);(i < 2*m.ROWS+1);i++  {
		
		for j := byte(0); j < 2*m.COLUMNS+1; j++ {
			//Add Horizontal Walls
			if (i%2 == 0 && j%2 == 1) {
				if (m.horizontalWalls[i/2][j/2] == true) {
					fmt.Print(" ---")
				} else {
					fmt.Print("    ")
				}
			}

			//Add Vertical Walls
			if i%2 == 1 && j%2 == 0 {
				if m.verticalWalls[i/2][j/2] == true {
					fmt.Print("|")
				} else {
					fmt.Print(" ")
				}
			}

			//Add Flood Fill Values
			if i%2 == 1 && j%2 == 1 {
				if (i-1)/2 == m.mouseRow && (j-1)/2 == m.mouseColumn {
					if m.MouseHeading == NORTH {
						fmt.Print(" ↑ ")
					} else if m.MouseHeading == EAST {
						fmt.Print(" → ")
					} else if m.MouseHeading == SOUTH {
						fmt.Print(" ↓ ")
					} else if m.MouseHeading == WEST {
						fmt.Print(" ← ")
					}
				} else {
					value := byte(m.values[(i-1)/2][(j-1)/2])
					if value >= 100 {
						fmt.Print(value)
					} else {
						fmt.Print(" ")
						fmt.Print(value)
					}
					if value < 10 {
						fmt.Print(" ")
					}
				}
			}
		}
		fmt.Print("\n")
	}
	
	fmt.Println("\n")
}

func Solve(m Maze){


        for i:=byte(0);i<m.ROWS;i++ {
            for j:=byte(0);j<m.COLUMNS;j++            {
                m.values[i][j]=255;
            }
     }
         m.values[m.targetRow][m.targetColumn] = 0
	 
         solvingFlag := true;
        for solvingFlag        {   solvingFlag= false;
            for i:=byte(0);i<m.ROWS;i++            {
            for j:=byte(0);j<m.COLUMNS;j++                {
                    if(m.values[i][j]<255)                    {
                        for k:=byte(0);k<4;k++                        {
                             neigbhoringCellRow:= i + byte(neigbhoringCells[k][0]);
                             neigbhoringCellColumn:= j + byte(neigbhoringCells[k][1]);
							
                             neigbhoringWallRow:= i + neigbhoringWalls[k][0];
                             neigbhoringWallColumn:= j + neigbhoringWalls[k][1];
                          
                             wallExists:= false;
							if(neigbhoringCellRow==255){
								neigbhoringCellRow=0;
							}
							if(neigbhoringCellColumn==255){
								neigbhoringCellColumn=0;
							}
                            if(k==NORTH||k==SOUTH){
                                wallExists = m.horizontalWalls[neigbhoringWallRow][neigbhoringWallColumn];
                            	
                            }else{ //Must be east or west since they are equal
                                wallExists = m.verticalWalls[neigbhoringWallRow][neigbhoringWallColumn];
                            	
                            }
                                fmt.Println(neigbhoringCellRow,neigbhoringCellColumn);
                            if((m.values[neigbhoringCellRow][neigbhoringCellColumn]==255) && !(wallExists))                            {   
                                m.values[neigbhoringCellRow][neigbhoringCellColumn]=m.values[i][j]+1;
                                solvingFlag=true;
                            }
                        }
                    }
                }
             }
        }
    
    }
 func AddWall(m Maze, cardinialDirection byte){

        switch cardinialDirection{
            case NORTH:
                m.horizontalWalls[m.mouseRow][m.mouseColumn] = true;
                break;
            case EAST:
                m.verticalWalls[m.mouseRow][m.mouseColumn+1] = true;
                break;
            case SOUTH:
                m.horizontalWalls[m.mouseRow+1][m.mouseColumn] = true;
                break;
            case WEST:
                m.verticalWalls[m.mouseRow][m.mouseColumn] = true;
                break;
        }
 }
  func AddWallByLocation(m Maze, row, column, c_Direction byte)    {
        /*
        *Have to remember the how to find the specific location using the array above
        *The function will take a byte arguments
        *THe function has to have a third argument to specify wheater the wall is vertical or horzintal
        */
        switch c_Direction{
            case NORTH:
                m.horizontalWalls[row][column] = true;
                break;
            case EAST:
                m.verticalWalls[row][column+1] = true;
                break;
            case SOUTH:
                m.horizontalWalls[row+1][column] = true;
                break;
            case WEST:
                m.verticalWalls[row][column] = true;
                break;
        }
  	
  }