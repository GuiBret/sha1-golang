package binaryoperations

import (
	"fmt"
	"strconv"
	"testing"
)

func TestShouldReturnAnErrorSinceCounterValueIsInvalid(t *testing.T) {

	currB, _ := strconv.ParseUint("12345678", 10, 64)
	currC, _ := strconv.ParseUint("90ABCDEF", 10, 64)
	currD, _ := strconv.ParseUint("01234567", 10, 64)

	fmt.Printf("%d", currB)
	_, err := PerformOperationOnValue(currB, currC, currD, 80)

	if err == nil {
		t.Errorf("Expected an error and did not get one")
	}
}

func TestShouldPerformTheFirstOperation(t *testing.T) {
	bAsHex, _ := StringToHex("abcdefgh")
	cAsHex, _ := StringToHex("ijklmnop")
	dAsHex, _ := StringToHex("qrstuvwx")
	bAsUint, _ := strconv.ParseUint(bAsHex, 10, 64)
	cAsUint, _ := strconv.ParseUint(cAsHex, 10, 64)
	dAsUint, _ := strconv.ParseUint(dAsHex, 10, 64)

	got, _ := PerformOperationOnValue(bAsUint, cAsUint, dAsUint, 0)
	var want uint64 = 0x1
	if got != want {
		t.Errorf("Got %d want %d", got, want)
	}

}

func TestShouldStringToHex(t *testing.T) {
	got, _ := StringToHex("abcdefgh")
	want := "6162636465666768"
	if got != want {
		t.Errorf("Got %s want %s", got, want)
	}
}

func TestShouldReturnAnErrorSinceTheStringIsEmpty(t *testing.T) {
	_, err := StringToHex("")

	if err == nil {
		t.Errorf("Expected an error and did not get one")
	}
}
