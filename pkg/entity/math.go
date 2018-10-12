package entity

import (
	"errors"
)

// MathOperationType types of math operation
type MathOperationType string

const (
	// Sum SUM
	Sum MathOperationType = "SUM"
	// Sub SUB
	Sub MathOperationType = "SUB"
)

// MathOperationRequest defines a request to math operation
type MathOperationRequest struct {
	Operation  MathOperationType `json:"operation"`
	Parameters []float64         `json:"parameters"`
}

// MathOperationResponse defines a response to math operation
type MathOperationResponse struct {
	Result float64 `json:"result"`
}

var (
	// ErrInvalidOperation Error for invalid operation
	ErrInvalidOperation = errors.New("invalid value 'operation'")

	// ErrInvalidParameters Error for invalid parameters
	ErrInvalidParameters = errors.New("invalid value 'parameters'")
)

// Validate Return error when MathOperationRequest is not valid
func (m *MathOperationRequest) Validate() error {

	if len(m.Operation) == 0 {
		return ErrInvalidOperation
	}

	if len(m.Parameters) == 0 {
		return ErrInvalidParameters
	}

	return nil
}
