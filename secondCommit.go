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

func findCoordinateWithoutCollison(size [2]int, marsRover rover) ([3]string, error) {

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

func findCoordinateWithCollison(size [2]int, marsRover rover, previousCoordinateX int, previousCoordinateY int) ([3]string, error) {

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
		if x == previousCoordinateX && y == previousCoordinateY {
			return errorReturn, errors.New("Warning collison")
		}

	}

	marsRover.coordinate[0] = strconv.Itoa(x)
	marsRover.coordinate[1] = strconv.Itoa(y)

	return marsRover.coordinate, nil
}

func withCollison() {
	var err error

	size := [2]int{5, 5}

	var rover1 rover
	rover1.coordinate = [3]string{"1", "2", "N"}
	rover1.instruction = "LMLMLMLMM"

	var rover2 rover
	rover2.coordinate = [3]string{"3", "3", "E"}
	rover2.instruction = "MMRMMRMRRM"

	rover1.coordinate, err = findCoordinateWithoutCollison(size, rover1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rover1.coordinate)
	}

	x, err := strconv.Atoi(rover1.coordinate[0])
	y, err := strconv.Atoi(rover1.coordinate[1])
	if err == nil {

		rover2.coordinate, err = findCoordinateWithCollison(size, rover2, x, y)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(rover2.coordinate)
		}

	}
}
func withOutCollison(){
	var err error

	size := [2]int{5, 5}

	var rover1 rover
	rover1.coordinate = [3]string{"1", "2", "N"}
	rover1.instruction = "LMLMLMLMM"

	var rover2 rover
	rover2.coordinate = [3]string{"3", "3", "E"}
	rover2.instruction = "MMRMMRMRRM"

	rover1.coordinate, err = findCoordinateWithoutCollison(size, rover1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rover1.coordinate)
	}

	rover2.coordinate, err = findCoordinateWithoutCollison(size, rover2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rover2.coordinate)
	}

}

func stepByStepMain(){
	var err error

	size := [2]int{5, 5}

	var rover1 rover
	rover1.coordinate = [3]string{"1", "2", "N"}
	rover1.instruction = "LMLMLMLMMR"

	var rover2 rover
	rover2.coordinate = [3]string{"3", "3", "E"}
	rover2.instruction = "MMRMMRMRRM"




	if err != nil {
		return 
	}

	for i := 0; i < len(rover1.instruction) || i < len(rover2.instruction); i++ {

		rover1.coordinate, err = stepByStep(size, rover1, string(rover1.instruction[i]), rover2 )
		rover2.coordinate, err = stepByStep(size, rover2, string(rover1.instruction[i]), rover1 )
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rover1.coordinate)
		fmt.Println(rover2.coordinate)
	}



}

func stepByStep(size [2]int, marsRover rover, order string, otherRover rover) ([3]string ,error){

	x, err := strconv.Atoi(marsRover.coordinate[0])
	y, err := strconv.Atoi(marsRover.coordinate[1])

	otherx, err := strconv.Atoi(otherRover.coordinate[0])
	othery, err := strconv.Atoi(otherRover.coordinate[1])

	errorReturn := [3]string{"0", "0", "0"}

	if err != nil {
		return errorReturn, errors.New("Couldn't get the coordinates")
	}

	
		direct := strings.Index(compass, marsRover.coordinate[2])

		switch os := order; os {
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
		if(x == otherx && y == othery){
			return marsRover.coordinate, errors.New("warning collison")
		}
		
	

	marsRover.coordinate[0] = strconv.Itoa(x)
	marsRover.coordinate[1] = strconv.Itoa(y)

	return marsRover.coordinate, nil
}

func main() {
	
	withOutCollison()
	withCollison()
	stepByStepMain()


}
