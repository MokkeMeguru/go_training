/*
Package sample is the package to calculate two interger numbers with safe gards
*/
package sample

import "errors"

func plus(x int, y int) (int, error) {
	return x + y, nil

}

func minus(x int, y int) (int, error) {
	return x - y, nil
}

func divide(x int, y int) (int, error) {
	if y == 0 {
		return -1, errors.New("invalid number y cannot be 0")
	} else {
		return x / y, nil
	}

}

func multiple(x int, y int) (int, error) {
	return x * y, nil
}
