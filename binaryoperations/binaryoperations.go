package binaryoperations

import (
	"errors"
	"fmt"
)

func PerformOperationOnValue(b uint64, c uint64, d uint64, counter int) (res uint64, err error) {

	if counter < 0 || counter >= 80 {
		return 0x0, errors.New("Invalid counter")

	}

	if counter >= 0 && counter <= 19 {
		fmt.Printf("B : %d", b)
		fmt.Printf("C : %d", c)
		fmt.Printf("Res : %d", (b & c))
		return (b&c | (^b & d)), nil
	}
	return 0x0, nil
}

func StringToHex(str string) (string, error) {
	result := ""

	if len(str) == 0 {
		return "", errors.New("Empty string")
	}

	for _, char := range str {

		result += fmt.Sprintf("%x", char)
	}

	return result, nil
}
