package errors

import "fmt"

var (
	ErrInvalidAge    = fmt.Errorf("empty age")
	ErrEmptyPosition = fmt.Errorf("position is empty")
)
