package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// constructing the rover
type rover struct {
	coordinate  [3]string
	instruction string
	direction   string
}

// to control and change direction
var compass string = "NESW"
var err error
var size [2]int
var rover1 rover
var rover2 rover

// for test purpose change variables here
func intilazeRovers() {

	size = [2]int{5, 5}
	rover1.coordinate = [3]string{"1", "2", "N"}
	rover1.instruction = "LMLMLMLMM"
	rover2.coordinate = [3]string{"3", "3", "E"}
	rover2.instruction = "MMRMMRMRRM"
}

func findCoordinateWithoutCollison(size [2]int, marsRover rover, marsRover2 rover) ([3]string, error) {
	//initial error is nil
	err = nil
	//for loop for every instruction bit
	for i := 0; i < len(marsRover.instruction); i++ {
		if err != nil {
			return marsRover.coordinate, err
		}
		//stepByStep function returns new coordinates if there is no error
		marsRover.coordinate, err = stepByStep(size, marsRover, string(marsRover.instruction[i]), marsRover2, false)
	}
	//returning new coordinates
	return marsRover.coordinate, err
}

func findCoordinateWithCollison(size [2]int, marsRover rover, marsRover2 rover) ([3]string, error) {
	//initial error is nil
	err = nil
	//for loop for every instruction bit
	for i := 0; i < len(marsRover.instruction); i++ {
		if err != nil {
			return marsRover.coordinate, err
		}
		//stepByStep function returns new coordinates if there is no error
		marsRover.coordinate, err = stepByStep(size, marsRover, string(marsRover.instruction[i]), marsRover2, true)
	}

	//returning new coordinates
	return marsRover.coordinate, nil
}

func withCollison() {
	//initilazing rovers
	intilazeRovers()

	rover1.coordinate, err = findCoordinateWithoutCollison(size, rover1, rover2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rover1.coordinate)
	}

	rover2.coordinate, err = findCoordinateWithCollison(size, rover2, rover1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rover2.coordinate)
	}

}
func withOutCollison() {
	intilazeRovers()
	rover1.coordinate, err = findCoordinateWithoutCollison(size, rover1, rover2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rover1.coordinate)
	}

	rover2.coordinate, err = findCoordinateWithoutCollison(size, rover2, rover1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rover2.coordinate)
	}

}

func stepByStepMain() error {

	intilazeRovers()
	//rover1 make a move then rover2 make a move for every loop
	for i := 0; i < len(rover1.instruction) && i < len(rover2.instruction); i++ {
		if err != nil {
			return err
		}
		rover1.coordinate, err = stepByStep(size, rover1, string(rover1.instruction[i]), rover2, true)
		rover2.coordinate, err = stepByStep(size, rover2, string(rover1.instruction[i]), rover1, true)
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rover1.coordinate)
		fmt.Println(rover2.coordinate)
	}
	return err

}

func stepByStep(size [2]int, marsRover rover, order string, otherRover rover, collisonOn bool) ([3]string, error) {
	//taking x and y coordinates from rover
	x, err := strconv.Atoi(marsRover.coordinate[0])
	y, err := strconv.Atoi(marsRover.coordinate[1])
	//if collison is needed to be checked taking other rovers coordinates
	otherx, err := strconv.Atoi(otherRover.coordinate[0])
	othery, err := strconv.Atoi(otherRover.coordinate[1])
	//initilazing "0" string slice to return in case of an error
	errorReturn := [3]string{"0", "0", "0"}
	//checking if there is an error for string to int
	if err != nil {
		return errorReturn, errors.New("Couldn't get the coordinates")
	}

	//getting the direction of the rover
	direct := strings.Index(compass, marsRover.coordinate[2])
	//switch case for instruction
	switch os := order; os {
	case "L":
		//if direction is "W" for left turn we cant get this from compass so giving the value manually
		if direct == 3 {
			marsRover.coordinate[2] = "N"
		} else {
			//other cases we can get the next direction from compass
			//for left turn we need to go right on compass
			marsRover.coordinate[2] = string(compass[direct+1])
		}

	case "R":
		//same as L but for right turns we need to go left on compass
		if direct == 0 {
			marsRover.coordinate[2] = "W"
		} else {
			marsRover.coordinate[2] = string(compass[direct-1])
		}
	case "M":
		//for direction "N" "M" means add 1 to y coordinate, for "E" minus 1 to x coordinate and so on
		switch yon := string(marsRover.coordinate[2]); yon {
		case "N":
			//checking the if next move will be in or out of the area
			if y < size[1] {
				y++
			} else {
				return errorReturn, errors.New("there was a miscalculation in instructions")
			}

		case "E":
			if x > 0 {
				x--
			} else {
				return errorReturn, errors.New("there was a miscalculation in instructions")
			}
		case "S":
			if y > 0 {
				y--
			} else {
				return errorReturn, errors.New("there was a miscalculation in instructions")
			}
		case "W":
			if x < size[0] {
				x++
			} else {
				return errorReturn, errors.New("there was a miscalculation in instructions")
			}

		}

	}
	//if collionOn comes true checks the collison
	if collisonOn {
		if x == otherx && y == othery {
			return marsRover.coordinate, errors.New("warning collison")
		}
	}
	//gives the new x and y values to rovers coordinates
	marsRover.coordinate[0] = strconv.Itoa(x)
	marsRover.coordinate[1] = strconv.Itoa(y)

	return marsRover.coordinate, nil
}

func main() {

	withOutCollison()
	withCollison()
	err := stepByStepMain()
	if err != nil {
		fmt.Println(err)
	}

}
