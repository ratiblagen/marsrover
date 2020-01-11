package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type rover struct {
	coordinate  [3]string
	instruction string
	direction   string
}

var compass string = "NESW"

func findCoordinate(size [2]int, marsRover rover) ([3]string, error) {

	x, err := strconv.Atoi(marsRover.coordinate[0])
	y, err := strconv.Atoi(marsRover.coordinate[1])
	
	errorReturn := [3]string{"0", "0", "0"}

	if err != nil {
		return errorReturn, errors.New("Couldn't get the coordinates")
	}

	for i := 0; i < len(marsRover.instruction); i++ {

		direct := strings.Index(compass, marsRover.coordinate[2])

		switch os := string(marsRover.instruction[i]); os {
		case "L":
			if direct == 3 {
				marsRover.coordinate[2] = "N"
			} else {
				marsRover.coordinate[2] = string(compass[direct+1])
			}

		case "R":
			if direct == 0 {
				marsRover.coordinate[2] = "W"
			} else {
				marsRover.coordinate[2] = string(compass[direct-1])
			}
		case "M":
			switch yon := string(marsRover.coordinate[2]); yon {
			case "N":
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
	}

	marsRover.coordinate[0] = strconv.Itoa(x)
	marsRover.coordinate[1] = strconv.Itoa(y)

	return marsRover.coordinate, nil
}

func main() {
	var err error

	size := [2]int{5, 5}

	var rover1 rover
	rover1.coordinate = [3]string{"1", "2", "N"}
	rover1.instruction = "LMLMLMLMM"

	var rover2 rover
	rover2.coordinate = [3]string{"3", "3", "E"}
	rover2.instruction = "MMRMMRMRRM"

	rover1.coordinate, err = findCoordinate(size, rover1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rover1.coordinate)
	}

	rover2.coordinate, err = findCoordinate(size, rover2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rover2.coordinate)
	}
}
