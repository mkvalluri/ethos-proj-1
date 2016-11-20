package main

import (
	"ethos/efmt"
	"math"
)

func main() {
	//Create new Box variable and pass values to it.
	myBox := Box { 2, 2, 10, 10}
	efmt.Println("Valluri_Output: Input values are l_x:",myBox.l_x,", l_y:",myBox.l_y,", h_x:",myBox.h_x,", h_y:",myBox.h_y)
	
	//Read values from myBox and calculate distance.
	dist := math.Sqrt(math.Pow(2, (myBox.h_x - myBox.l_x)) + math.Pow(2, (myBox.h_y - myBox.l_y)))
	
	//Print the distance that has been calculated
	efmt.Println("Valluri_Output: Distance between points:", dist)
}
