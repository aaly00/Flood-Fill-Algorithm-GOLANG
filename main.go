package main

import (
	"floodfill"
)

func main() {
	maze := floodFill.Maze{}
	maze = maze.Constructor(5, 5,2,2)
	floodFill.Solve(maze);
	//  fmt.Print(maze.ROWS);
	floodFill.AddWall(maze,1)
	floodFill.AddWallByLocation(maze,3,3,0)
	maze.Print()
}
