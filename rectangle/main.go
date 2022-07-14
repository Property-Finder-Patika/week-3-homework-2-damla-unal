package main

import (
	"errors"
	"fmt"
)

/*
Create a rectangle abstraction using struct.
Create two functions to calculate area and circumference of given rectangle instance and
set produced values on given rectangle instance.
Create another function to create an instance of rectangle struct and return it.
But that function would be able to return an error in case of passing invalid arguments such as negative length or height.
Choose one of strategies mentioned i the text book and exemplified in error.go of ch05 in example project.
*/

// create a rectangle abstraction
type rectangle struct {
	length int
	width  int
}

//createRectangle creates an instance of rectangle struct and return it
func createRectangle(length int, width int) rectangle {
	return rectangle{
		length: length,
		width:  width,
	}
}

//calculateArea calculates the ares of the given rectangle
func calculateArea(r rectangle) int {
	return r.width * r.length
}

//calculateCircumference calculates the circumference of the given rectangle
func calculateCircumference(r rectangle) int {
	return 2 * (r.width + r.length)
}

//createRectangleWithValidation firstly checks the validations then creates an instance of rectangle struct and return it with error
func createRectangleWithValidation(length int, width int) (rectangle, error) {
	if length == 0 || width == 0 {
		return rectangle{}, errors.New("Zero values: " + fmt.Sprintf("length: %d, width: %d", length, width))
	}
	if length < 0 || width < 0 {
		return rectangle{}, errors.New("Negative values: " + fmt.Sprintf("length: %d, width: %d", length, width))
	}

	return createRectangle(length, width), nil
}

func main() {
	rec, err := createRectangleWithValidation(5, 3)
	if err != nil {
		fmt.Printf("Error occurred: %s\n", err)
	}
	fmt.Printf("This is a rectangle: %+v\n", rec)

	fmt.Printf("Area of this rectange: %d\n", calculateArea(rec))
	fmt.Printf("Circumference of this rectange: %d\n", calculateCircumference(rec))

	// implemented strategy 1 for the error handling
	negativeRec, err := createRectangleWithValidation(-8, 2)
	if err != nil {
		fmt.Printf("Error occurred: %s, this %+v rectangle is not legal\n", err, negativeRec)
	}

	// implemented strategy 1 for the error handling
	zeroRec, err := createRectangleWithValidation(8, 0)
	if err != nil {
		fmt.Printf("Error occurred: %s, this %+v rectangle is not legal\n", err, zeroRec)
	}

}
