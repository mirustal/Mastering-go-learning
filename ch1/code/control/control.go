package control

import (
	"fmt"
	"os"
	"strconv"
)

func Control() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a command line argument")
		return
	}
	argument := os.Args[1]

	switch argument {
	case "0":
		fmt.Println("zero")
	case "1":
		fmt.Println("one")
	case "2", "3", "4":
		fmt.Println("two, three or four")
		fallthrough
	default:
		fmt.Println("Value:  ", argument)
	}

	value, err := strconv.Atoi(argument)
	if err != nil {
		fmt.Println("Cannot convert to int:", argument)
		return
	}

	switch {
	case value == 0:
		fmt.Println("zero")
	case value < 0:
		fmt.Println("negative int")
	case value > 0:
		fmt.Println("positive int")
	default:
		fmt.Println("This should not happen", value)
	}

}
