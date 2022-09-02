package basic

import (
	"errors"
	"fmt"
)

func errorException() {
	newNumber := 5
	checkNumber, err := getError(newNumber)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(checkNumber)
}

func getError(number int) (string, error) {
	var message string
	if number < 10 {
		err := errors.New("Error: haha")
		return message, err
	} else {
		message = "Ok"
		return message, nil
	}
}
