package handler

import "testing"

func TestValidateNumber(t *testing.T) {
	//arrange
	number := []byte("123")

	// act
	err := validateNumber(number)

	// assert
	if err == nil {
		t.Error("validation does not work")
	}
}
