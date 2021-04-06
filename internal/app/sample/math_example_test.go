package sample

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlus(t *testing.T) {
	actual, err := plus(2, 3)
	expected := 5

	assert.Equal(t, actual, expected)
	assert.Equal(t, err, nil)
}

func TestMinus(t *testing.T) {
	actual, err := minus(2, 1)
	expected := 1

	assert.Equal(t, expected, actual)
	assert.Equal(t, nil, err)
}

func TestDivide_Success(t *testing.T) {
	actual, err := divide(10, 2)
	expected := 5

	assert.Equal(t, expected, actual)
	assert.Equal(t, nil, err)
}

func TestDivide_by_Zero(t *testing.T) {
	_, err := divide(10, 0)
	expected := "invalid number y cannot be 0"
	assert.Equal(t, expected, err.Error())
}
